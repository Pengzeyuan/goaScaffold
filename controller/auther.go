package controller

import (
	"context"
	"errors"
	"fmt"
	"starter/config"
	"starter/middleware"
	"starter/service"
	"time"

	"git.chinaopen.ai/sc-ncov/tif"
	"github.com/dgrijalva/jwt-go"
	"goa.design/goa/v3/security"

	"starter/gen/log"
)

var (
	ErrorUnauthorized = errors.New("请登录后再试")

	KeyUserID = "userId"
)

type Auther struct {
	logger *log.Logger
	cache  *service.CacheService
	tif    *tif.Client
}

// APIKeyAuth implements the authorization logic for service "rework" for the
// "api_key" security scheme.
func (a *Auther) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {
	tifUid, _ := ctx.Value(middleware.RequestTifUidKey).(string)

	// 保存用户ID
	ctx = context.WithValue(ctx, KeyUserID, tifUid)

	return ctx, nil
}

// JWTAuth implements the authorization logic for service "secured_service" for
// the "jwt" security scheme.
func (a *Auther) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// logger := L(ctx, a.logger)

	claims := make(jwt.MapClaims)

	// authorize request
	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	if _, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.C.Jwt.Secret), nil
	}); err != nil {
		return ctx, MakeUnauthorizedError(ctx, "invalid token")
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, MakeUnauthorizedError(ctx, "invalid scopes in token")
	}

	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, MakeUnauthorizedError(ctx, "无操作权限")
	}

	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}

	if err := a.validateScopes(scheme.RequiredScopes, scopesInToken); err != nil {
		return ctx, MakeForbiddenError(ctx, err.Error())
	}

	currentUserID, ok := claims["jti"].(string)
	if !ok {
		return ctx, MakeUnauthorizedError(ctx, "无操作权限")
	}

	// 保存用户ID
	ctx = context.WithValue(ctx, KeyUserID, currentUserID)

	return ctx, nil
}

// create JWT token
func (a *Auther) createJwtToken(userID string, userType int, scopes []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti":    userID,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Second * time.Duration(config.C.Jwt.ExpireIn)).Unix(),
		"type":   userType,
		"scopes": scopes,
	})

	// note that if "SignedString" returns an error then it is returned as
	// an internal error to the client
	return token.SignedString([]byte(config.C.Jwt.Secret))
}

func (a *Auther) validateScopes(expected, actual []string) error {
	for _, r := range expected {
		found := false
		for _, s := range actual {
			if s == r {
				found = true
				break
			}
		}
		if found {
			return nil
		}
	}
	return fmt.Errorf("您没有权限进行此操作")
}

// 获取当前用户ID
func (a *Auther) GetCurrentUserID(ctx context.Context) (string, error) {
	if ctx != nil {
		userID, _ := ctx.Value(KeyUserID).(string)
		return userID, nil
	}
	return "", errors.New("ctx is nil")
}
