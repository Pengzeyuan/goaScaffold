package utils

import (
	"boot/config"
	"boot/model"

	"encoding/json"

	"github.com/nats-io/nats.go"

	"go.uber.org/zap"
)

var (
	NatsCli *nats.Conn
)

func ConnectNats() error {
	if NatsCli != nil {
		return nil
	}
	cli, err := nats.Connect(config.C.Nats.Urls)
	if err != nil {
		return err
	}
	NatsCli = cli
	return nil
}

var (
	callUserID = "userId"
)

// 发布消息详情
func PublishRealInfoEvalMsg(infoType int, data interface{}) {

	logger := zap.L().With(zap.String("import", "gzzwdp"))
	infoModel := model.PushInfoEvent{
		CallNoID:    callUserID,
		InfoType:    infoType,
		InfoDetails: data,
	}
	b, _ := json.Marshal(infoModel)
	// 通过nat发布窗口评价信息
	if err := NatsCli.Publish("channal", b); err != nil {
		logger.Error("nats 发布推送消息失败", zap.Error(err))
		return
	}
	logger.Info("发布推送消息成功")
}
