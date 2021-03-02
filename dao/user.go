package dao

import (
	"boot/model"
	"boot/pkg/cache"
	libsgorm "git.chinaopen.ai/yottacloud/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

type UserDao interface {
	// 创建超级管理员用户
	CreateUser(userName, password string) (model.AopUser, error)
	// 重设密码
	//SetAdmPassword(user *model.AopUser, rawPassword string) (model.AopUser, error)
	// Get 获取用户
	Get(constructor UserScopeConstructor) (model.AopUser, error)
	// 更新
	Update(userName string, updateFields map[string]interface{}, constructor UserScopeConstructor) (model.AopUser, error)
}

func NewUserDaoImpl(db *gorm.DB, logger *zap.Logger) UserDao {
	return UserDaoImpl{db: db, logger: logger}
}

type NewUserDaoFunc = func(*gorm.DB, *zap.Logger) UserDao

type UserDaoImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func (u UserDaoImpl) Get(constructor UserScopeConstructor) (model.AopUser, error) {
	db := u.db.Scopes(constructor.Scopes()...)
	var account model.AopUser
	if err := db.Find(&account).Error; err != nil {
		u.logger.Error("get account failed", zap.Error(err))
		return account, err
	}
	return account, nil
}

func (u UserDaoImpl) CreateUser(userName, password string) (model.AopUser, error) {
	user := model.AopUser{
		BaseModel: model.BaseModel{},
		UserName:  userName,
	}
	var err error
	user.Password, err = user.CreatePassword(password)
	if err != nil {
		u.logger.Error("create account failed", zap.Error(err))
		return model.AopUser{}, err
	}
	if err := u.db.Create(&user).Error; err != nil {
		u.logger.Error("create user failed", zap.Error(err))
		return model.AopUser{}, err
	}
	return user, nil
}

func (u UserDaoImpl) SetAdmPassword(user *model.AopUser, rawPassword string) (model.AopUser, error) {
	hashedPassword, err := user.CreatePassword(rawPassword)
	if err != nil {
		return *user, err
	}
	var account model.AopUser
	if errs := u.db.Model(&user).Where("id=?", user.ID).Update("password", hashedPassword).Scan(&account).Error; errs != nil {
		return account, errs
	}
	return account, nil
}

func (u UserDaoImpl) Update(userName string, updateFields map[string]interface{},
	constructor UserScopeConstructor) (model.AopUser, error) {
	var account model.AopUser
	updateFields["updated_at"] = time.Now()
	db := u.db.Scopes(constructor.Scopes()...)
	if err := db.Model(&account).Where("user_name = ? ", userName).
		Updates(updateFields).Scan(&account).Error; err != nil {
		u.logger.Error("update account failed", zap.Error(err))
		return account, err
	}
	return account, nil
}

type UserScopeConstructor struct {
	scopes []libsgorm.Scope
}

func NewUserScopeConstructor() UserScopeConstructor {
	return UserScopeConstructor{
		scopes: []libsgorm.Scope{},
	}
}

// 获取scopes
func (u UserScopeConstructor) Scopes() []libsgorm.Scope {
	return u.scopes
}

func (u UserScopeConstructor) AddUserName(userName string) UserScopeConstructor {
	query := func(db *gorm.DB) *gorm.DB {
		return db.Where("user_name = ?", userName)
	}
	u.scopes = append(u.scopes, query)
	return u
}

func NewUserDao(db *gorm.DB, cache *cache.CacheService, logger *zap.Logger) *userDao {
	return &userDao{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

type userDao struct {
	db     *gorm.DB
	logger *zap.Logger
	cache  *cache.CacheService
}
