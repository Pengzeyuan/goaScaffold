package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("thirdPart", func() {
	Description("拉取第三方接口数据")

	Error("bad_request", ErrorResult)
	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})
	HTTP(func() {
		Path("/third_part")
		Response("bad_request", StatusBadRequest)
		Response("internal_server_error", StatusInternalServerError)
	})

	Method("GetActualTimeData", func() {
		Description("接收大厅管理的数据")
		Meta("swagger:summary", "接收大厅管理的数据")

		Payload(func() {
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
			Attribute("data", ArrayOf(HallManagementResp))
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			GET("/get_hall_management_data")
			Response(StatusOK)
		})
	})

	Method("ReceiveThirdPartyPushData", func() {
		Description("接收第三方推送数据--大厅排队办事实时图基础数据--排号、叫号、办结、评价四阶段")
		Meta("swagger:summary", "接收第三方推送数据")
		Payload(func() {
			Attribute("methodName", Int, "推送的具体方法", func() {
				Example(1)
			})
			Attribute("count", Int, "数据数量", func() {
				Example(21)
			})
			Attribute("data", Any, "第三方推送数据", func() {
				Example("")
			})

			Required("methodName", "count", "data")
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
			Attribute("result", String, "success", func() {
				Example("getDataSuccess")
			})
			Required("errcode", "errmsg", "result")
		})

		HTTP(func() {
			POST("/receive_third_party_push_data")
			Response(StatusOK)
		})
	})

	Method("GormRelatedSearch", func() {
		Description("gorm关联查询")
		Meta("swagger:summary", "gorm关联查询")

		Payload(func() {
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
			Attribute("data", ArrayOf(LegalPersonUserResp))
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			GET("/gorm_related_search")
			Response(StatusOK)
		})
	})

})
