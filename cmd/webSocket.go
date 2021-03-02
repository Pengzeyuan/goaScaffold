package main

import (
	"boot/config"
	"boot/model"
	"boot/pkg/sync/errgroup"
	"boot/pkg/websocket"
	"boot/utils"
	"context"
	"encoding/json"
	"fmt"
	"git.chinaopen.ai/yottacloud/go-libs/panichandler"
	"github.com/golang/protobuf/proto"
	ws "github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/withlin/canal-go/client"
	protocol "github.com/withlin/canal-go/protocol"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	srv *http.Server
}

func NewServer(addr string, router *http.ServeMux) *Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return &Server{srv: srv}
}

func WebsocketCmd() *cobra.Command {
	runWebsocketCmd := &cobra.Command{
		Use:   "ws",
		Short: "websocket监听服务",
		Long:  `websocket监听服务`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := zap.L().With(zap.String("import", "gzzwdp"))
			if err := utils.ConnectNats(); err != nil {
				logger.Error("connect nats error", zap.Error(err))
				return err
			}
			RunWsServer(config.C.Server)
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}
	runWebsocketCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return runWebsocketCmd
}

//  开始websocket 服务
func RunWsServer(c config.ServerConfig) {
	// 配置 socket 设置
	router := http.NewServeMux()
	logger := zap.L().With(zap.String("import", "gzzwdp"))
	switch c.Host {
	case "", "localhost", "127.0.0.1":
		c.Host = "0.0.0.0"
	default:
	}

	if c.WebsocketPort == 0 {
		c.WebsocketPort = 8083
	}

	addr := fmt.Sprintf("%s:%d", c.Host, c.WebsocketPort)
	//go util.SetOff(logger)
	//  一个新websocket 管理者
	wsConnManager := websocket.NewManager()
	// 路由的地址
	router.HandleFunc("/gyzwdp-websocks/ws/oneno/notice", func(w http.ResponseWriter, r *http.Request) {
		var (
			// ws 是websocket工具 的网络升级者 5秒的超时
			upgrader = &ws.Upgrader{
				ReadBufferSize:   1024,
				WriteBufferSize:  1024,
				HandshakeTimeout: 5 * time.Second,
				// 解决跨域问题
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}
			// 代表一个 websocket 连接
			wsConn *ws.Conn
			err    error
			// 自己封装的 Connection
			conn *websocket.Connection
		)
		//logger.Info("websocket 监听路径" + fmt.Sprintf(r.RequestURI)) // 打印监听路径，
		logger.Info("websocket 新增一个连接")
		// 升级 读写  成一个 websocket Conn
		if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil { // 升级为websocket
			logger.Error("连接 websocket 失败", zap.Error(err))
			return
		}
		// 在封装一道
		conn = websocket.NewConnection(wsConn) // 初始化conn
		userId := r.FormValue("userId")        // 获取前端传递的id
		//  chan 的订阅
		_ = wsConnManager.Subscribe(userId, conn)

		// 启动读协程，防止异常导致所有服务挂掉
		go func() {
			defer panichandler.ZapHandler(logger).Handle()
			conn.ReadLoop()
		}()
		// 启动写协程，防止异常导致所有服务挂掉
		go func() {
			defer panichandler.ZapHandler(logger).Handle()
			go conn.WriteLoop()
		}()
	})

	newSvr := NewServer(addr, router)
	done := make(chan struct{}, 1)
	group := errgroup.WithContext(context.Background())
	// 监听键入
	group.Go(func(ctx context.Context) error {

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
		for {
			select {
			case <-quit:
				return ctx.Err()
			case <-ctx.Done():
				return nil
			}
		}
	})

	group.Go(func(ctx context.Context) error {
		for {
			//  nats  订阅   了   然后发给  chan
			_, err := utils.NatsCli.QueueSubscribe("channal", "queue", func(m *nats.Msg) { // 主题按需求修改
				natsMessage := &model.PushInfoEvent{} // 数据类型按需求修改
				if err := json.Unmarshal(m.Data, &natsMessage); err != nil {
					logger.Error("反序列化失败出错", zap.Error(err))
					return
				}
				// chan  的发布
				wsConnManager.Publish(m.Data, logger)
			})
			if err != nil {
				logger.Error("订阅 websocket 通知失败", zap.Error(err))
				continue
			}
			_ = utils.NatsCli.Flush()

			select {
			case <-time.After(time.Millisecond * 500):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	group.Go(func(ctx context.Context) error {
		go func() {
			<-ctx.Done()
			ctx2, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := newSvr.Shutdown(ctx2); err != nil {
				logger.Error("出现了一个错误")
			}
			done <- struct{}{}
		}()
		return newSvr.Start()
	})

	if err := group.Wait(); err != nil {
		logger.Info("websocket 服务启动失败")
	}
	<-done
	logger.Info("执行完毕")
}

func (s *Server) Start() error {
	logger := zap.L().With(zap.String("import", "gzzwdp"))
	logger.Info("开始监听端口......")
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger := zap.L().With(zap.String("import", "gzzwdp"))
	logger.Info("关闭端口......")
	return s.srv.Shutdown(ctx)
}

func CanalCmd() *cobra.Command {
	runCanalCmd := &cobra.Command{
		Use:   "canal",
		Short: "canal接收消息发送给websocket",
		Long:  `canal接收消息发送给websocket`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := zap.L().With(zap.String("import", "gzzwdp"))
			if err := utils.ConnectNats(); err != nil {
				logger.Error("connect nats error", zap.Error(err))
				return err
			}
			RunCanalServer(config.C.Canal)
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			return nil
		},
	}
	runCanalCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	return runCanalCmd
}

func RunCanalServer(c config.CanalConf) {
	router := http.NewServeMux()
	logger := zap.L().With(zap.String("import", "gzzwdp"))
	switch c.Host {
	case "0.0.0.0", "":
		c.Host = "127.0.0.1"
	default:
	}
	if c.Port == 0 {
		c.Port = 8086
	}
	connector := client.NewSimpleCanalConnector(c.Host, c.Port, "", "", c.Destination, c.SoTimeOut, c.IdleTimeOut)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	logger.Info("canal 连接成功......")
	err = connector.Subscribe("gyzw_dp\\.*")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	logger.Info("canal 开始监听......")

	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	//  new 一个  http
	newSvr := NewServer(addr, router)
	done := make(chan struct{}, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	//  错误组
	group := errgroup.WithContext(context.Background())
	// 等着 ctx 来Done
	group.Go(func(ctx context.Context) error {
		for {
			select {
			case <-quit:
				return ctx.Err()
			case <-ctx.Done():
				return nil
			}
		}
	})

	group.Go(func(ctx context.Context) error {
		//  一协程
		go func() {
			<-ctx.Done()
			ctx2, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := newSvr.Shutdown(ctx2); err != nil {
				logger.Error("出现了一个错误")
			}
			done <- struct{}{}
		}()
		//  循环来  拿消息
		for {
			message, err := connector.Get(100, nil, nil)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			batchId := message.Id
			if batchId == -1 || len(message.Entries) <= 0 {
				time.Sleep(300 * time.Millisecond)
				continue
			}
			// 发消息
			printEntry(message.Entries)
		}
	})
	//  等带着 所有go完  取消函数
	if err := group.Wait(); err != nil {
		logger.Info("canal 服务启动失败")
	}
	<-done
	logger.Info("执行完毕")
}

func printEntry(entrys []protocol.Entry) {
	//logger := zap.L().With(zap.String("import", "gzzwdp"))
	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		eventType := rowChange.GetEventType()
		//header := entry.GetHeader()
		if eventType == protocol.EventType_INSERT {
			//logger.Info(fmt.Sprintf("===========日志信息==========: tableName:[%s,%s], eventType: %s", header.GetSchemaName(), header.GetTableName(), header.GetEventType()))
			for _, rowData := range rowChange.GetRowDatas() {
				printColumn(rowData.GetAfterColumns())
			}
		}
	}
}

func printColumn(columns []*protocol.Column) {
	colMap := make(map[string]string)
	for _, col := range columns {
		colMap[col.GetName()] = col.GetValue()
		fmt.Println(fmt.Sprintf("%s : %s", col.GetName(), col.GetValue()))
	}
	utils.PublishCanalMsg(colMap)
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
