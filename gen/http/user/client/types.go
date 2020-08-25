// Code generated by goa v3.2.3, DO NOT EDIT.
//
// User HTTP client types
//
// Command:
// $ goa gen starter/design

package client

import (
	user "starter/gen/user"

	goa "goa.design/goa/v3/pkg"
)

// LoginByUsernameRequestBody is the type of the "User" service
// "LoginByUsername" endpoint HTTP request body.
type LoginByUsernameRequestBody struct {
	// 用户名
	Username string `form:"username" json:"username" xml:"username"`
	// 密码
	Password string `form:"password" json:"password" xml:"password"`
	// 图形验证码
	HumanCode string `form:"humanCode" json:"humanCode" xml:"humanCode"`
	// 图形验证码ID
	CaptchaID string `form:"captchaId" json:"captchaId" xml:"captchaId"`
}

// LoginBySmsCodeRequestBody is the type of the "User" service "LoginBySmsCode"
// endpoint HTTP request body.
type LoginBySmsCodeRequestBody struct {
	// 手机号
	Mobile string `form:"mobile" json:"mobile" xml:"mobile"`
	// 短信验证码
	VerifyCode string `form:"verify_code" json:"verify_code" xml:"verify_code"`
}

// UpdatePasswordRequestBody is the type of the "User" service "UpdatePassword"
// endpoint HTTP request body.
type UpdatePasswordRequestBody struct {
	OldPassword string `form:"old_password" json:"old_password" xml:"old_password"`
	NewPassword string `form:"new_password" json:"new_password" xml:"new_password"`
}

// SendSmsCodeRequestBody is the type of the "User" service "SendSmsCode"
// endpoint HTTP request body.
type SendSmsCodeRequestBody struct {
	// 手机号
	Mobile string `form:"mobile" json:"mobile" xml:"mobile"`
	// 图形验证码
	HumanCode string `form:"humanCode" json:"humanCode" xml:"humanCode"`
	// 图形验证码ID
	CaptchaID string `form:"captchaId" json:"captchaId" xml:"captchaId"`
}

// LoginByUsernameResponseBody is the type of the "User" service
// "LoginByUsername" endpoint HTTP response body.
type LoginByUsernameResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string              `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *SessionResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// LoginBySmsCodeResponseBody is the type of the "User" service
// "LoginBySmsCode" endpoint HTTP response body.
type LoginBySmsCodeResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string              `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *SessionResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// UpdatePasswordResponseBody is the type of the "User" service
// "UpdatePassword" endpoint HTTP response body.
type UpdatePasswordResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string                    `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *SuccessResultResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// GetCaptchaImageResponseBody is the type of the "User" service
// "GetCaptchaImage" endpoint HTTP response body.
type GetCaptchaImageResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string              `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *CaptchaResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// SendSmsCodeResponseBody is the type of the "User" service "SendSmsCode"
// endpoint HTTP response body.
type SendSmsCodeResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string                    `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *SuccessResultResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// LoginByUsernameInternalServerErrorResponseBody is the type of the "User"
// service "LoginByUsername" endpoint HTTP response body for the
// "internal_server_error" error.
type LoginByUsernameInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LoginByUsernameBadRequestResponseBody is the type of the "User" service
// "LoginByUsername" endpoint HTTP response body for the "bad_request" error.
type LoginByUsernameBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LoginBySmsCodeInternalServerErrorResponseBody is the type of the "User"
// service "LoginBySmsCode" endpoint HTTP response body for the
// "internal_server_error" error.
type LoginBySmsCodeInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// LoginBySmsCodeBadRequestResponseBody is the type of the "User" service
// "LoginBySmsCode" endpoint HTTP response body for the "bad_request" error.
type LoginBySmsCodeBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdatePasswordInternalServerErrorResponseBody is the type of the "User"
// service "UpdatePassword" endpoint HTTP response body for the
// "internal_server_error" error.
type UpdatePasswordInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdatePasswordBadRequestResponseBody is the type of the "User" service
// "UpdatePassword" endpoint HTTP response body for the "bad_request" error.
type UpdatePasswordBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetCaptchaImageInternalServerErrorResponseBody is the type of the "User"
// service "GetCaptchaImage" endpoint HTTP response body for the
// "internal_server_error" error.
type GetCaptchaImageInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetCaptchaImageBadRequestResponseBody is the type of the "User" service
// "GetCaptchaImage" endpoint HTTP response body for the "bad_request" error.
type GetCaptchaImageBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SendSmsCodeInternalServerErrorResponseBody is the type of the "User" service
// "SendSmsCode" endpoint HTTP response body for the "internal_server_error"
// error.
type SendSmsCodeInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SendSmsCodeBadRequestResponseBody is the type of the "User" service
// "SendSmsCode" endpoint HTTP response body for the "bad_request" error.
type SendSmsCodeBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SessionResponseBody is used to define fields on response body types.
type SessionResponseBody struct {
	User        *UserResponseBody        `form:"user,omitempty" json:"user,omitempty" xml:"user,omitempty"`
	Credentials *CredentialsResponseBody `form:"credentials,omitempty" json:"credentials,omitempty" xml:"credentials,omitempty"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	// ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 用户名
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// 昵称
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" xml:"nickname,omitempty"`
	// 手机号
	Mobile *string `form:"mobile,omitempty" json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 是否可用
	IsActive *bool `form:"isActive,omitempty" json:"isActive,omitempty" xml:"isActive,omitempty"`
}

// CredentialsResponseBody is used to define fields on response body types.
type CredentialsResponseBody struct {
	// JWT token
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	// 有效时长（秒）：生成之后x秒内有效
	ExpiresIn *int `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
}

// SuccessResultResponseBody is used to define fields on response body types.
type SuccessResultResponseBody struct {
	// success
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// CaptchaResponseBody is used to define fields on response body types.
type CaptchaResponseBody struct {
	// 图片base64
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// 验证码ID
	CaptchaID *string `form:"captchaId,omitempty" json:"captchaId,omitempty" xml:"captchaId,omitempty"`
}

// NewLoginByUsernameRequestBody builds the HTTP request body from the payload
// of the "LoginByUsername" endpoint of the "User" service.
func NewLoginByUsernameRequestBody(p *user.LoginByUsernamePayload) *LoginByUsernameRequestBody {
	body := &LoginByUsernameRequestBody{
		Username:  p.Username,
		Password:  p.Password,
		HumanCode: p.HumanCode,
		CaptchaID: p.CaptchaID,
	}
	return body
}

// NewLoginBySmsCodeRequestBody builds the HTTP request body from the payload
// of the "LoginBySmsCode" endpoint of the "User" service.
func NewLoginBySmsCodeRequestBody(p *user.LoginBySmsCodePayload) *LoginBySmsCodeRequestBody {
	body := &LoginBySmsCodeRequestBody{
		Mobile:     p.Mobile,
		VerifyCode: p.VerifyCode,
	}
	return body
}

// NewUpdatePasswordRequestBody builds the HTTP request body from the payload
// of the "UpdatePassword" endpoint of the "User" service.
func NewUpdatePasswordRequestBody(p *user.UpdatePasswordPayload) *UpdatePasswordRequestBody {
	body := &UpdatePasswordRequestBody{
		OldPassword: p.OldPassword,
		NewPassword: p.NewPassword,
	}
	return body
}

// NewSendSmsCodeRequestBody builds the HTTP request body from the payload of
// the "SendSmsCode" endpoint of the "User" service.
func NewSendSmsCodeRequestBody(p *user.SendSmsCodePayload) *SendSmsCodeRequestBody {
	body := &SendSmsCodeRequestBody{
		Mobile:    p.Mobile,
		HumanCode: p.HumanCode,
		CaptchaID: p.CaptchaID,
	}
	return body
}

// NewLoginByUsernameResultOK builds a "User" service "LoginByUsername"
// endpoint result from a HTTP "OK" response.
func NewLoginByUsernameResultOK(body *LoginByUsernameResponseBody) *user.LoginByUsernameResult {
	v := &user.LoginByUsernameResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	if body.Data != nil {
		v.Data = unmarshalSessionResponseBodyToUserSession(body.Data)
	}

	return v
}

// NewLoginByUsernameInternalServerError builds a User service LoginByUsername
// endpoint internal_server_error error.
func NewLoginByUsernameInternalServerError(body *LoginByUsernameInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewLoginByUsernameBadRequest builds a User service LoginByUsername endpoint
// bad_request error.
func NewLoginByUsernameBadRequest(body *LoginByUsernameBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewLoginBySmsCodeResultOK builds a "User" service "LoginBySmsCode" endpoint
// result from a HTTP "OK" response.
func NewLoginBySmsCodeResultOK(body *LoginBySmsCodeResponseBody) *user.LoginBySmsCodeResult {
	v := &user.LoginBySmsCodeResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	if body.Data != nil {
		v.Data = unmarshalSessionResponseBodyToUserSession(body.Data)
	}

	return v
}

// NewLoginBySmsCodeInternalServerError builds a User service LoginBySmsCode
// endpoint internal_server_error error.
func NewLoginBySmsCodeInternalServerError(body *LoginBySmsCodeInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewLoginBySmsCodeBadRequest builds a User service LoginBySmsCode endpoint
// bad_request error.
func NewLoginBySmsCodeBadRequest(body *LoginBySmsCodeBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdatePasswordResultOK builds a "User" service "UpdatePassword" endpoint
// result from a HTTP "OK" response.
func NewUpdatePasswordResultOK(body *UpdatePasswordResponseBody) *user.UpdatePasswordResult {
	v := &user.UpdatePasswordResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	if body.Data != nil {
		v.Data = unmarshalSuccessResultResponseBodyToUserSuccessResult(body.Data)
	}

	return v
}

// NewUpdatePasswordInternalServerError builds a User service UpdatePassword
// endpoint internal_server_error error.
func NewUpdatePasswordInternalServerError(body *UpdatePasswordInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdatePasswordBadRequest builds a User service UpdatePassword endpoint
// bad_request error.
func NewUpdatePasswordBadRequest(body *UpdatePasswordBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetCaptchaImageResultOK builds a "User" service "GetCaptchaImage"
// endpoint result from a HTTP "OK" response.
func NewGetCaptchaImageResultOK(body *GetCaptchaImageResponseBody) *user.GetCaptchaImageResult {
	v := &user.GetCaptchaImageResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	if body.Data != nil {
		v.Data = unmarshalCaptchaResponseBodyToUserCaptcha(body.Data)
	}

	return v
}

// NewGetCaptchaImageInternalServerError builds a User service GetCaptchaImage
// endpoint internal_server_error error.
func NewGetCaptchaImageInternalServerError(body *GetCaptchaImageInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetCaptchaImageBadRequest builds a User service GetCaptchaImage endpoint
// bad_request error.
func NewGetCaptchaImageBadRequest(body *GetCaptchaImageBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSendSmsCodeResultOK builds a "User" service "SendSmsCode" endpoint result
// from a HTTP "OK" response.
func NewSendSmsCodeResultOK(body *SendSmsCodeResponseBody) *user.SendSmsCodeResult {
	v := &user.SendSmsCodeResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	if body.Data != nil {
		v.Data = unmarshalSuccessResultResponseBodyToUserSuccessResult(body.Data)
	}

	return v
}

// NewSendSmsCodeInternalServerError builds a User service SendSmsCode endpoint
// internal_server_error error.
func NewSendSmsCodeInternalServerError(body *SendSmsCodeInternalServerErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSendSmsCodeBadRequest builds a User service SendSmsCode endpoint
// bad_request error.
func NewSendSmsCodeBadRequest(body *SendSmsCodeBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateLoginByUsernameResponseBody runs the validations defined on
// LoginByUsernameResponseBody
func ValidateLoginByUsernameResponseBody(body *LoginByUsernameResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Errcode != nil {
		if *body.Errcode < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 0, true))
		}
	}
	if body.Errcode != nil {
		if *body.Errcode > 999999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 999999, false))
		}
	}
	if body.Data != nil {
		if err2 := ValidateSessionResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateLoginBySmsCodeResponseBody runs the validations defined on
// LoginBySmsCodeResponseBody
func ValidateLoginBySmsCodeResponseBody(body *LoginBySmsCodeResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Errcode != nil {
		if *body.Errcode < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 0, true))
		}
	}
	if body.Errcode != nil {
		if *body.Errcode > 999999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 999999, false))
		}
	}
	if body.Data != nil {
		if err2 := ValidateSessionResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdatePasswordResponseBody runs the validations defined on
// UpdatePasswordResponseBody
func ValidateUpdatePasswordResponseBody(body *UpdatePasswordResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Errcode != nil {
		if *body.Errcode < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 0, true))
		}
	}
	if body.Errcode != nil {
		if *body.Errcode > 999999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 999999, false))
		}
	}
	if body.Data != nil {
		if err2 := ValidateSuccessResultResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetCaptchaImageResponseBody runs the validations defined on
// GetCaptchaImageResponseBody
func ValidateGetCaptchaImageResponseBody(body *GetCaptchaImageResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Errcode != nil {
		if *body.Errcode < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 0, true))
		}
	}
	if body.Errcode != nil {
		if *body.Errcode > 999999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 999999, false))
		}
	}
	if body.Data != nil {
		if err2 := ValidateCaptchaResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateSendSmsCodeResponseBody runs the validations defined on
// SendSmsCodeResponseBody
func ValidateSendSmsCodeResponseBody(body *SendSmsCodeResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Errcode != nil {
		if *body.Errcode < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 0, true))
		}
	}
	if body.Errcode != nil {
		if *body.Errcode > 999999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.errcode", *body.Errcode, 999999, false))
		}
	}
	if body.Data != nil {
		if err2 := ValidateSuccessResultResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateLoginByUsernameInternalServerErrorResponseBody runs the validations
// defined on LoginByUsername_internal_server_error_Response_Body
func ValidateLoginByUsernameInternalServerErrorResponseBody(body *LoginByUsernameInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLoginByUsernameBadRequestResponseBody runs the validations defined
// on LoginByUsername_bad_request_Response_Body
func ValidateLoginByUsernameBadRequestResponseBody(body *LoginByUsernameBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLoginBySmsCodeInternalServerErrorResponseBody runs the validations
// defined on LoginBySmsCode_internal_server_error_Response_Body
func ValidateLoginBySmsCodeInternalServerErrorResponseBody(body *LoginBySmsCodeInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateLoginBySmsCodeBadRequestResponseBody runs the validations defined on
// LoginBySmsCode_bad_request_Response_Body
func ValidateLoginBySmsCodeBadRequestResponseBody(body *LoginBySmsCodeBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdatePasswordInternalServerErrorResponseBody runs the validations
// defined on UpdatePassword_internal_server_error_Response_Body
func ValidateUpdatePasswordInternalServerErrorResponseBody(body *UpdatePasswordInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdatePasswordBadRequestResponseBody runs the validations defined on
// UpdatePassword_bad_request_Response_Body
func ValidateUpdatePasswordBadRequestResponseBody(body *UpdatePasswordBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetCaptchaImageInternalServerErrorResponseBody runs the validations
// defined on GetCaptchaImage_internal_server_error_Response_Body
func ValidateGetCaptchaImageInternalServerErrorResponseBody(body *GetCaptchaImageInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetCaptchaImageBadRequestResponseBody runs the validations defined
// on GetCaptchaImage_bad_request_Response_Body
func ValidateGetCaptchaImageBadRequestResponseBody(body *GetCaptchaImageBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSendSmsCodeInternalServerErrorResponseBody runs the validations
// defined on SendSmsCode_internal_server_error_Response_Body
func ValidateSendSmsCodeInternalServerErrorResponseBody(body *SendSmsCodeInternalServerErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSendSmsCodeBadRequestResponseBody runs the validations defined on
// SendSmsCode_bad_request_Response_Body
func ValidateSendSmsCodeBadRequestResponseBody(body *SendSmsCodeBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSessionResponseBody runs the validations defined on
// SessionResponseBody
func ValidateSessionResponseBody(body *SessionResponseBody) (err error) {
	if body.User == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user", "body"))
	}
	if body.Credentials == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("credentials", "body"))
	}
	if body.User != nil {
		if err2 := ValidateUserResponseBody(body.User); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Credentials != nil {
		if err2 := ValidateCredentialsResponseBody(body.Credentials); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUserResponseBody runs the validations defined on UserResponseBody
func ValidateUserResponseBody(body *UserResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Nickname == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("nickname", "body"))
	}
	if body.Mobile == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("mobile", "body"))
	}
	if body.IsActive == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("isActive", "body"))
	}
	return
}

// ValidateCredentialsResponseBody runs the validations defined on
// CredentialsResponseBody
func ValidateCredentialsResponseBody(body *CredentialsResponseBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	if body.ExpiresIn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expires_in", "body"))
	}
	return
}

// ValidateSuccessResultResponseBody runs the validations defined on
// SuccessResultResponseBody
func ValidateSuccessResultResponseBody(body *SuccessResultResponseBody) (err error) {
	if body.OK == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ok", "body"))
	}
	return
}

// ValidateCaptchaResponseBody runs the validations defined on
// CaptchaResponseBody
func ValidateCaptchaResponseBody(body *CaptchaResponseBody) (err error) {
	if body.Image == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image", "body"))
	}
	if body.CaptchaID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("captchaId", "body"))
	}
	return
}
