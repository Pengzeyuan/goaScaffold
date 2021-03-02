package utils

import (
	"boot/config"
	"boot/model"

	"encoding/json"

	"go.uber.org/zap"
)

func PublishCanalMsg(data interface{}) {
	logger := zap.L().With(zap.String("import", "gzzwdp"))
	infoModel := model.CanalInfo{
		CallNoID:    "id",
		InfoDetails: data,
	}
	c, _ := json.Marshal(infoModel)
	// 通过nat发布窗口评价信息
	if err := NatsCli.Publish(config.C.Nats.CanalSubject, c); err != nil {
		logger.Error("nats 发布canal消息失败", zap.Error(err))
		return
	}
	logger.Info("发布canal消息成功")
}
