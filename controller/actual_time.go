package boot

import (
	"boot/dao"
	actualtime "boot/gen/actual_time"
	log "boot/gen/log"
	"boot/model"
	"boot/serializer"
	"boot/service"
	"boot/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	protocol "github.com/withlin/canal-go/protocol"
	"go.uber.org/zap"
	"os"
)

// ActualTime service example implementation.
// The example methods log the requests and return zero values.
type actualTimesrvc struct {
	logger *log.Logger
}

// NewActualTime returns the ActualTime service implementation.
func NewActualTime(logger *log.Logger) actualtime.Service {
	return &actualTimesrvc{logger}
}

// 接收第三方推送数据--大厅排队办事实时图基础数据
func (s *actualTimesrvc) ReceiveThirdPartyPushData(ctx context.Context, payload *actualtime.ReceiveThirdPartyPushDataPayload) (res *actualtime.ReceiveThirdPartyPushDataResult, err error) {
	res = &actualtime.ReceiveThirdPartyPushDataResult{}
	logger := L(ctx, s.logger)
	s.logger.Info("actualTime.ReceiveThirdPartyPushData")
	svc := service.NewHallActualTimeSVCImpl(ctx, dao.DpDB, logger)

	utils.PublishRealInfoEvalMsg(payload.MethodName, payload.Data) // 发布消息
	marshal, err := json.Marshal(payload.Data)
	queryModel := model.CommonQueryModel{
		PullData: marshal,
		Method:   payload.MethodName,
		Count:    payload.Count,
	}
	actualTimeResp, err := svc.ReceiveThirdPartyPushData(queryModel)
	if err != nil {
		logger.Error("接收第三方推送数据失败", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "接收第三方推送数据失败")
	}
	res.Result = actualTimeResp.Msg
	return res, nil
}

// 接收数据库监听得到的数据
func (s *actualTimesrvc) GetActualTimeData(ctx context.Context) (res *actualtime.GetActualTimeDataResult, err error) {
	res = &actualtime.GetActualTimeDataResult{}
	logger := L(ctx, s.logger)
	s.logger.Info("actualTime.GetActualTimeData")
	connector := dao.Connector

	message, err := connector.Get(100, nil, nil)
	if err != nil {
		logger.Error("获取错误......")
		return res, err
	}
	batchId := message.Id
	if batchId != -1 && len(message.Entries) > 0 {
		dataMap := printEntry(message.Entries)
		a := &model.CanalData{
			DataType:    1,
			InfoDetails: dataMap,
		}
		res.Data = serializer.CanalData2CanalData(a)
	}

	return res, nil
}
func printEntry(entrys []protocol.Entry) []map[string]string {
	var arrMap []map[string]string
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
				colMap := make(map[string]string, 0)
				for _, col := range rowData.GetAfterColumns() {
					colMap[col.GetName()] = col.GetValue()
					//fmt.Println(fmt.Sprintf("%s : %s", col.GetName(), col.GetValue()))
				}
				arrMap = append(arrMap, colMap)
			}
		}
	}
	return arrMap
}
func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
