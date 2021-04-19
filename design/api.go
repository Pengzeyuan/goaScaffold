package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zaplogger"
)

var _ = API("boot", func() {
	Title("goa微服务")
	HTTP(func() {
		Path("/api")
	})

	Server("boot", func() {
		Description("goa微服务")
		Services("User")
		Services("entity_hall")
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
	Description("使用 JWT 认证, 需要认证的接口添加 ```Header```: ```Authorization: Bearer {jwtToken}```")
	Scope("role:user", "用户")
	Scope("role:admin", "管理员")
	Scope("api:read", "只读权限")
	Scope("api:write", "读写权限")
	Scope("api:admin", "管理权限")
})
