// Code generated by goa v3.2.4, DO NOT EDIT.
//
// User endpoints
//
// Command:
// $ goa gen boot/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "User" service endpoints.
type Endpoints struct {
	LoginByUsername goa.Endpoint
	LoginBySmsCode  goa.Endpoint
	UpdatePassword  goa.Endpoint
	GetCaptchaImage goa.Endpoint
	SendSmsCode     goa.Endpoint
}

// NewEndpoints wraps the methods of the "User" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		LoginByUsername: NewLoginByUsernameEndpoint(s),
		LoginBySmsCode:  NewLoginBySmsCodeEndpoint(s),
		UpdatePassword:  NewUpdatePasswordEndpoint(s, a.JWTAuth),
		GetCaptchaImage: NewGetCaptchaImageEndpoint(s),
		SendSmsCode:     NewSendSmsCodeEndpoint(s),
	}
}

// Use applies the given middleware to all the "User" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.LoginByUsername = m(e.LoginByUsername)
	e.LoginBySmsCode = m(e.LoginBySmsCode)
	e.UpdatePassword = m(e.UpdatePassword)
	e.GetCaptchaImage = m(e.GetCaptchaImage)
	e.SendSmsCode = m(e.SendSmsCode)
}

// NewLoginByUsernameEndpoint returns an endpoint function that calls the
// method "LoginByUsername" of service "User".
func NewLoginByUsernameEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginByUsernamePayload)
		return s.LoginByUsername(ctx, p)
	}
}

// NewLoginBySmsCodeEndpoint returns an endpoint function that calls the method
// "LoginBySmsCode" of service "User".
func NewLoginBySmsCodeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginBySmsCodePayload)
		return s.LoginBySmsCode(ctx, p)
	}
}

// NewUpdatePasswordEndpoint returns an endpoint function that calls the method
// "UpdatePassword" of service "User".
func NewUpdatePasswordEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePasswordPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"role:user", "role:admin"},
			RequiredScopes: []string{"role:user"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.UpdatePassword(ctx, p)
	}
}

// NewGetCaptchaImageEndpoint returns an endpoint function that calls the
// method "GetCaptchaImage" of service "User".
func NewGetCaptchaImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetCaptchaImage(ctx)
	}
}

// NewSendSmsCodeEndpoint returns an endpoint function that calls the method
// "SendSmsCode" of service "User".
func NewSendSmsCodeEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SendSmsCodePayload)
		return s.SendSmsCode(ctx, p)
	}
}
