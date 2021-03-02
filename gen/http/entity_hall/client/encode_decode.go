// Code generated by goa v3.2.4, DO NOT EDIT.
//
// entity_hall HTTP client encoders and decoders
//
// Command:
// $ goa gen boot/design

package client

import (
	entityhall "boot/gen/entity_hall"
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildWaitLineOverviewRequest instantiates a HTTP request object with method
// and path set to call the "entity_hall" service "WaitLineOverview" endpoint
func (c *Client) BuildWaitLineOverviewRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: WaitLineOverviewEntityHallPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("entity_hall", "WaitLineOverview", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeWaitLineOverviewRequest returns an encoder for requests sent to the
// entity_hall WaitLineOverview server.
func EncodeWaitLineOverviewRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*entityhall.WaitLineOverviewPayload)
		if !ok {
			return goahttp.ErrInvalidType("entity_hall", "WaitLineOverview", "*entityhall.WaitLineOverviewPayload", v)
		}
		body := NewWaitLineOverviewRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("entity_hall", "WaitLineOverview", err)
		}
		return nil
	}
}

// DecodeWaitLineOverviewResponse returns a decoder for responses returned by
// the entity_hall WaitLineOverview endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeWaitLineOverviewResponse may return the following errors:
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "internal_server_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeWaitLineOverviewResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body WaitLineOverviewResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity_hall", "WaitLineOverview", err)
			}
			err = ValidateWaitLineOverviewResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("entity_hall", "WaitLineOverview", err)
			}
			res := NewWaitLineOverviewResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body WaitLineOverviewBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity_hall", "WaitLineOverview", err)
			}
			err = ValidateWaitLineOverviewBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("entity_hall", "WaitLineOverview", err)
			}
			return nil, NewWaitLineOverviewBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body WaitLineOverviewInternalServerErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("entity_hall", "WaitLineOverview", err)
			}
			err = ValidateWaitLineOverviewInternalServerErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("entity_hall", "WaitLineOverview", err)
			}
			return nil, NewWaitLineOverviewInternalServerError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("entity_hall", "WaitLineOverview", resp.StatusCode, string(body))
		}
	}
}

// unmarshalWaitLineOverviewRespResponseBodyToEntityhallWaitLineOverviewResp
// builds a value of type *entityhall.WaitLineOverviewResp from a value of type
// *WaitLineOverviewRespResponseBody.
func unmarshalWaitLineOverviewRespResponseBodyToEntityhallWaitLineOverviewResp(v *WaitLineOverviewRespResponseBody) *entityhall.WaitLineOverviewResp {
	res := &entityhall.WaitLineOverviewResp{
		TodayDQ:      *v.TodayDQ,
		CumulativeDQ: *v.CumulativeDQ,
	}

	return res
}
