package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("starter", func() {
	Title("微服务")

	Server("starter", func() {
		Description("微服务")
		Services("starter")

		Host("localhost", func() {
			Description("default host")
			URI("http://localhost:8000/starter")
			URI("grpc://localhost:8080/starter")
		})
	})
})

// APIKeyAuth defines a security scheme that uses API keys.
var APIKeyAuth = APIKeySecurity("api_key", func() {
	Description("API Key 认证")
})

// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description(`JWT 认证`)
	Scope("role:user", "用户")
})
