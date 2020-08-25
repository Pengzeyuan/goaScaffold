// Code generated by goa v3.2.3, DO NOT EDIT.
//
// User client
//
// Command:
// $ goa gen starter/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "User" service client.
type Client struct {
	LoginByUsernameEndpoint goa.Endpoint
	LoginBySmsCodeEndpoint  goa.Endpoint
	UpdatePasswordEndpoint  goa.Endpoint
	GetCaptchaImageEndpoint goa.Endpoint
	SendSmsCodeEndpoint     goa.Endpoint
}

// NewClient initializes a "User" service client given the endpoints.
func NewClient(loginByUsername, loginBySmsCode, updatePassword, getCaptchaImage, sendSmsCode goa.Endpoint) *Client {
	return &Client{
		LoginByUsernameEndpoint: loginByUsername,
		LoginBySmsCodeEndpoint:  loginBySmsCode,
		UpdatePasswordEndpoint:  updatePassword,
		GetCaptchaImageEndpoint: getCaptchaImage,
		SendSmsCodeEndpoint:     sendSmsCode,
	}
}

// LoginByUsername calls the "LoginByUsername" endpoint of the "User" service.
func (c *Client) LoginByUsername(ctx context.Context, p *LoginByUsernamePayload) (res *LoginByUsernameResult, err error) {
	var ires interface{}
	ires, err = c.LoginByUsernameEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginByUsernameResult), nil
}

// LoginBySmsCode calls the "LoginBySmsCode" endpoint of the "User" service.
func (c *Client) LoginBySmsCode(ctx context.Context, p *LoginBySmsCodePayload) (res *LoginBySmsCodeResult, err error) {
	var ires interface{}
	ires, err = c.LoginBySmsCodeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginBySmsCodeResult), nil
}

// UpdatePassword calls the "UpdatePassword" endpoint of the "User" service.
func (c *Client) UpdatePassword(ctx context.Context, p *UpdatePasswordPayload) (res *UpdatePasswordResult, err error) {
	var ires interface{}
	ires, err = c.UpdatePasswordEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdatePasswordResult), nil
}

// GetCaptchaImage calls the "GetCaptchaImage" endpoint of the "User" service.
func (c *Client) GetCaptchaImage(ctx context.Context) (res *GetCaptchaImageResult, err error) {
	var ires interface{}
	ires, err = c.GetCaptchaImageEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*GetCaptchaImageResult), nil
}

// SendSmsCode calls the "SendSmsCode" endpoint of the "User" service.
func (c *Client) SendSmsCode(ctx context.Context, p *SendSmsCodePayload) (res *SendSmsCodeResult, err error) {
	var ires interface{}
	ires, err = c.SendSmsCodeEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SendSmsCodeResult), nil
}
