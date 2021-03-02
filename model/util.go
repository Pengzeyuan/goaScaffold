package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Strings []string

func (s *Strings) Scan(src interface{}) error {
	switch typ := src.(type) {
	default:
		return fmt.Errorf("%s not supported", typ)
	case []byte:
		return json.Unmarshal(src.([]byte), s)
	}
}

func (s Strings) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type BaseModel struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type BaseUUIDModel struct {
	ID        string     `gorm:"primaryKey;type:varchar(36);not null;"`
	CreatedAt time.Time  `msgpack:"-"`
	UpdatedAt time.Time  `msgpack:"-"`
	DeletedAt *time.Time `gorm:"index" msgpack:"-"`
}

func NewBaseUUIDModel() BaseUUIDModel {
	return BaseUUIDModel{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func IsDuplicateError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}

	return false
}

type CanalInfo struct {
	CallNoID    string      `json:"callNoID"`    // 叫号ID
	InfoDetails interface{} `json:"infoDetails"` // 监听消息详情
}
