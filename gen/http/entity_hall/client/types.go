// Code generated by goa v3.2.4, DO NOT EDIT.
//
// entity_hall HTTP client types
//
// Command:
// $ goa gen boot/design

package client

import (
	entityhall "boot/gen/entity_hall"

	goa "goa.design/goa/v3/pkg"
)

// WaitLineOverviewRequestBody is the type of the "entity_hall" service
// "WaitLineOverview" endpoint HTTP request body.
type WaitLineOverviewRequestBody struct {
	// 行政区划代码
	RegionCode string `form:"regionCode" json:"regionCode" xml:"regionCode"`
	// 起始时间
	StartDate *string `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
	// 结束时间
	EndDate *string `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
}

// WaitLineOverviewResponseBody is the type of the "entity_hall" service
// "WaitLineOverview" endpoint HTTP response body.
type WaitLineOverviewResponseBody struct {
	// 错误码
	Errcode *int `form:"errcode,omitempty" json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误消息
	Errmsg *string                           `form:"errmsg,omitempty" json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	Data   *WaitLineOverviewRespResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
}

// WaitLineOverviewBadRequestResponseBody is the type of the "entity_hall"
// service "WaitLineOverview" endpoint HTTP response body for the "bad_request"
// error.
type WaitLineOverviewBadRequestResponseBody struct {
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

// WaitLineOverviewInternalServerErrorResponseBody is the type of the
// "entity_hall" service "WaitLineOverview" endpoint HTTP response body for the
// "internal_server_error" error.
type WaitLineOverviewInternalServerErrorResponseBody struct {
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

// WaitLineOverviewRespResponseBody is used to define fields on response body
// types.
type WaitLineOverviewRespResponseBody struct {
	// 累计排号数
	TodayDQ *int32 `form:"todayDQ,omitempty" json:"todayDQ,omitempty" xml:"todayDQ,omitempty"`
	// 累计办件量
	CumulativeDQ *int32 `form:"cumulativeDQ,omitempty" json:"cumulativeDQ,omitempty" xml:"cumulativeDQ,omitempty"`
}

// NewWaitLineOverviewRequestBody builds the HTTP request body from the payload
// of the "WaitLineOverview" endpoint of the "entity_hall" service.
func NewWaitLineOverviewRequestBody(p *entityhall.WaitLineOverviewPayload) *WaitLineOverviewRequestBody {
	body := &WaitLineOverviewRequestBody{
		RegionCode: p.RegionCode,
		StartDate:  p.StartDate,
		EndDate:    p.EndDate,
	}
	return body
}

// NewWaitLineOverviewResultOK builds a "entity_hall" service
// "WaitLineOverview" endpoint result from a HTTP "OK" response.
func NewWaitLineOverviewResultOK(body *WaitLineOverviewResponseBody) *entityhall.WaitLineOverviewResult {
	v := &entityhall.WaitLineOverviewResult{
		Errcode: *body.Errcode,
		Errmsg:  *body.Errmsg,
	}
	v.Data = unmarshalWaitLineOverviewRespResponseBodyToEntityhallWaitLineOverviewResp(body.Data)

	return v
}

// NewWaitLineOverviewBadRequest builds a entity_hall service WaitLineOverview
// endpoint bad_request error.
func NewWaitLineOverviewBadRequest(body *WaitLineOverviewBadRequestResponseBody) *goa.ServiceError {
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

// NewWaitLineOverviewInternalServerError builds a entity_hall service
// WaitLineOverview endpoint internal_server_error error.
func NewWaitLineOverviewInternalServerError(body *WaitLineOverviewInternalServerErrorResponseBody) *goa.ServiceError {
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

// ValidateWaitLineOverviewResponseBody runs the validations defined on
// WaitLineOverviewResponseBody
func ValidateWaitLineOverviewResponseBody(body *WaitLineOverviewResponseBody) (err error) {
	if body.Errcode == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errcode", "body"))
	}
	if body.Errmsg == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("errmsg", "body"))
	}
	if body.Data == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("data", "body"))
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
		if err2 := ValidateWaitLineOverviewRespResponseBody(body.Data); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateWaitLineOverviewBadRequestResponseBody runs the validations defined
// on WaitLineOverview_bad_request_Response_Body
func ValidateWaitLineOverviewBadRequestResponseBody(body *WaitLineOverviewBadRequestResponseBody) (err error) {
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

// ValidateWaitLineOverviewInternalServerErrorResponseBody runs the validations
// defined on WaitLineOverview_internal_server_error_Response_Body
func ValidateWaitLineOverviewInternalServerErrorResponseBody(body *WaitLineOverviewInternalServerErrorResponseBody) (err error) {
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

// ValidateWaitLineOverviewRespResponseBody runs the validations defined on
// WaitLineOverviewRespResponseBody
func ValidateWaitLineOverviewRespResponseBody(body *WaitLineOverviewRespResponseBody) (err error) {
	if body.TodayDQ == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("todayDQ", "body"))
	}
	if body.CumulativeDQ == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cumulativeDQ", "body"))
	}
	return
}
