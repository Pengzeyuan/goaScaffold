package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseUUIDModel

	Username string  `gorm:"type:varchar(36);not null;unique;"`
	Password string  `gorm:"type:varchar(128);"`
	Mobile   *string `gorm:"type:varchar(15);unique;"`
	Nickname string  `gorm:"type:varchar(30);not null;defult:''"`
}

// CreatePassword 加密密码
func (u *User) CreatePassword(raw string) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the DefaultCost (10)

	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

// CheckPassword 检查密码
func (u *User) CheckPassword(plainPwd, hashedPwd string) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}
