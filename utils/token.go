package utils

import (
	goalibs "git.chinaopen.ai/yottacloud/go-libs/goa-libs"

	"github.com/dgrijalva/jwt-go"

	"time"
)

// jwt token 的最长有效期, default: 7天
const (
	JwtTokenCookieName = "jwt_token"
	JwtMaxLifetime     = 7 * 24 * time.Hour
)

// 生成token和过期时间(秒)
// @param id uuid
// @param tokenStr 生成的token字符串
// @param expireAt 有效截至时间
func GenToken(id string, expireAt time.Time) (tokenStr string, err error) {
	claims := goalibs.NewExtendedClaims(id, []string{"api:read", "api:write"}, expireAt)
	tokenStr, err = claims.Sign(jwt.SigningMethodHS256)
	return
}
