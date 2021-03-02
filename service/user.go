package service

import (
	"boot/dao"
	"boot/model"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

type UserSVC interface {
	// 根据昵称获取用户
	GetAccountByUserName(userName string) (model.AopUser, error)
	// 更新登陆时间
	UpdateLoginTime(userName string) error
}

type UserSVCImpl struct {
	db             *gorm.DB
	logger         *zap.Logger
	newUserDaoFunc dao.NewUserDaoFunc
}

func (u UserSVCImpl) GetAccountByUserName(userName string) (model.AopUser, error) {
	userDao := u.newUserDaoFunc(u.db, u.logger)
	scopes := dao.NewUserScopeConstructor()
	scopes = scopes.AddUserName(userName)
	account, err := userDao.Get(scopes)
	if err != nil {
		u.logger.Error("get account by mobile or nickname failed", zap.Error(err))
		return model.AopUser{}, err
	}
	return account, nil
}

func (u UserSVCImpl) UpdateLoginTime(userName string) error {
	userDao := u.newUserDaoFunc(u.db, u.logger)
	scopes := dao.NewUserScopeConstructor()
	updateFields := make(map[string]interface{})
	updateFields["login_at"] = time.Now()
	_, err := userDao.Update(userName, updateFields, scopes)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return err
		}
		u.logger.Error("update failed", zap.Error(err))
		return err
	}
	return nil
}

func NewUserSVCImpl(db *gorm.DB, logger *zap.Logger, daoFunc dao.NewUserDaoFunc) UserSVC {
	return UserSVCImpl{
		db:             db,
		logger:         logger,
		newUserDaoFunc: daoFunc,
	}
}
