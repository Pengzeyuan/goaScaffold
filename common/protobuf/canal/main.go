package main

import (
	"boot/common"
	"encoding/json"
	"fmt"
	"github.com/siddontang/go-log/log"
)

type MyEventHandler struct {
	DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *RowsEvent) error {
	log.Infof("%s %v\n", e.Action, e.Rows)
	// Action 是insert  e.Rows 是加入的列数据    Table。Columns 中有表列信息

	//colMaps := make([]map[string]string,0)
	colMap := make(map[string]interface{})
	for i := 0; i < len(e.Table.Columns); i++ {

		//newString := e.Rows[0][i].(string)
		colMap[e.Table.Columns[i].Name] = e.Rows[0][i]
	}
	//for _, oneRow :=range e.Rows{
	//	for _, col := range oneRow {
	//		colMap[col.GetName()] = col.GetValue()
	//		//fmt.Println(fmt.Sprintf("%s : %s", col.GetName(), col.GetValue()))
	//	}
	//}

	mapB, _ := json.Marshal(colMap)
	fmt.Printf("map的值：%s \n", mapB)

	//  直接调用接口

	// 通过nat发布列信息
	if err := common.NatsCli.Publish("columnChange", mapB); err != nil {
		fmt.Println("-------> 发送change日志失败")
	}
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	//  连接nats
	err := common.ConnectNats()
	if err != nil {
		log.Println(err)

	}

	cfg := NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "canal"
	cfg.Flavor = "mariadb"
	cfg.Password = "Canal@123456"
	cfg.ServerID = 100
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "ginessential"
	cfg.Dump.Tables = []string{"animals"}

	c, _ := NewCanal(cfg)

	eventHandler := MyEventHandler{
		DummyEventHandler{},
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&eventHandler)

	// Start canal
	c.Run()

}
