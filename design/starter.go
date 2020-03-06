package design

import . "goa.design/goa/v3/dsl"

var _ = Service("starter", func() {
	Description("微服务")

	Error("internal_server_error", ErrorResult)
	Error("bad_request", ErrorResult)

	HTTP(func() {
		Path("/starter")
		Response("internal_server_error", StatusInternalServerError)
		Response("bad_request", StatusBadRequest)
	})

	Method("LoginByUsername", func() {
		Description("使用账号密码登录")
		Meta("swagger:summary", "使用账号密码登录")
		Payload(func() {
			Attribute("username", String, "用户名", func() {
				Example("user")
				MinLength(1)
				MaxLength(128)
			})
			Attribute("password", String, "密码", func() {
				Example("password")
				MinLength(1)
				MaxLength(128)
			})
			Required("username", "password")
		})

		Result(func() {
			Attribute("errcode", Int, "错误码", func() {
				Minimum(0)
				Maximum(999999)
				Example(0)
			})
			Attribute("errmsg", String, "错误消息", func() {
				Example("")
			})
			Attribute("data", Session)
			Required("errcode", "errmsg")
		})

		HTTP(func() {
			POST("/login_by_username")
			Response(StatusOK)
		})
	})

})
