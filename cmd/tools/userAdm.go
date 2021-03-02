package tools

import (
	"boot/dao"
	"go.uber.org/zap"
)

func NewAdmTool() AdmTool {
	return AdmTool{}
}

type AdmTool struct {
	logger *zap.Logger
}

func (t *AdmTool) CreateAdm(userName, password string) error {
	logger := zap.L()
	tx := dao.DpDB.Begin() // 开始事务
	if tx.Error != nil {
		_ = tx.Rollback()
		logger.Error("begin tx failed", zap.Error(tx.Error))
		return tx.Error
	}
	u := dao.NewUserDaoImpl(tx, logger)
	_, errs := u.CreateUser(userName, password)
	if errs != nil {
		_ = tx.Rollback()
		t.logger.Error("create user failed", zap.Error(errs))
		return errs
	}
	if err := tx.Commit().Error; err != nil {
		_ = tx.Rollback()
		t.logger.Error("commit tx failed", zap.Error(err))
		return err
	}
	return nil
}
