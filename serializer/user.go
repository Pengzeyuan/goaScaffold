package serializer

import (
	"boot/gen/user"
	"boot/model"
)

func ModelAccount2AuthAopUser(account model.AopUser) *user.User {
	loginAopUser := &user.User{
		Username:  account.UserName,
		LoginTime: account.LoginAt.String(),
	}

	return loginAopUser
}
