package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("simulation", func() {
	Description("模拟数据")
	Error("bad_request", ErrorResult)
	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})
	HTTP(func() {
		Path("/simulation")
		Response("bad_request", StatusBadRequest)
		Response("internal_server_error", StatusInternalServerError)
	})

	Method("SetData", func() {
		Description("设置数据")
		Meta("swagger:summary", "设置数据")
		Security(JWTAuth, func() {
			Scope("api:write")
			Scope("api:read")
		})

		Payload(func() {
			Token("jwtToken", String, func() {
				Description("JWT used for authentication")
				Example("eyJhbGciOiJIUz...")
			})
			Attribute("key", String, "对应模块", func() {
				Example("对应模块")
			})
			Attribute("val", String, "对应值", func() {
				Example("对应值")
			})
			Attribute("isShowMock", Boolean, "是否显示配置项", func() {
				Example(true)
			})
			Attribute("orderBy", Int, "排序规则", func() {
				Default(1)
				Description("1、按评价分正序;2、按评价分倒序")
				Example(1)

			})
			Attribute("orderTimeScope", Int, "排序时间范围", func() {
				Default(1)
				Description("1、按当年度评价;2、按上个月评价")
				Example(1)
			})
			Required("key", "val", "isShowMock")
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
			Attribute("result", SuccessResult)
			Required("errcode", "errmsg", "result")
		})

		HTTP(func() {
			POST("/set_data")
			Response(StatusOK)
		})
	})

	Method("GetData", func() {
		Description("获取模拟数据")
		Meta("swagger:summary", "获取模拟数据")
		Security(JWTAuth, func() {
			Scope("api:write")
			Scope("api:read")
		})

		Payload(func() {
			Token("jwtToken", String, func() {
				Description("JWT used for authentication")
				Example("eyJhbGciOiJIUz...")
			})
			Attribute("key", String, "对应模块", func() {
				Example("对应模块")
			})
			Required("key")
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
			Attribute("data", GetDataResp)
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			POST("/get_data")
			Response(StatusOK)
		})
	})
})
