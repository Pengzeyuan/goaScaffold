package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("import_file", func() {
	Description("插入Excel文件")
	Meta("swagger:tag:插入Excel文件")

	Error("bad_request")
	Error("internal_server_error", ErrorResult, func() {
		Fault()
	})

	HTTP(func() {
		Path("/import_file")
		Response("bad_request", StatusBadRequest)
		Response("internal_server_error", StatusInternalServerError)
	})

	Method("ImportExcelFile", func() {
		Description("excel数据批量导入，导入数据采用json格式存储")
		Meta("swagger:summary", "批量导入")

		Payload(func() {
			Attribute("file", Bytes, "上传文件内容", func() {
				MaxLength(50 * 1024 * 1024)
			})
			Attribute("filename", String, "文件名", func() {
				MaxLength(128)
			})
			Attribute("area", String, "区域", func() {
				Example("abc")
				MinLength(1)
			})
			Attribute("year", Int, "年份", func() {
				Example(2020)
			})
			Attribute("type", Int, "类型", func() {
				Example(1)
			})
			Required("file", "filename", "area", "year", "type")
		})

		Result(SuccessResult)

		HTTP(func() {
			POST("/import_excel_file")
			Header("area:X-Area")
			Header("year:X-Year")
			Header("type:X-Type")
			MultipartRequest()
		})
	})

	Method("GetImportExcelFileInfo", func() {
		Description("获取插入excel数据统计信息")
		Meta("swagger:summary", "获取插入excel数据统计信息")

		Payload(func() {
			Attribute("endYear", Int, "结束年份", func() {
				Example(2020)
			})
			Attribute("area", String, "区域代码", func() {
				Example("520103")
			})
			Required("endYear")
		})

		Result(FourDoCountResp)

		HTTP(func() {
			GET("/get_import_excel_file_info")
			Params(func() {
				Param("endYear")
				Param("area")
			})
			Response(StatusOK)
		})
	})

	Method("ReformOfAdministrative", func() {
		Description("行政审批制度改革事项详情")
		Meta("swagger:summary", "行政审批制度改革事项详情")
		Payload(func() {
			Attribute("regionCode", String, "行政区划代码", func() {
				Example("520100")
			})
			Attribute("startDate", String, "起始时间", func() {
				Example("2019-07-27")
			})
			Attribute("endDate", String, "结束时间", func() {
				Example("2020-07-27")
			})
			Required("regionCode")
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
			Attribute("data", ReformOfAdministrativeResp)
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			POST("/reform_of_administrative")
			Response(StatusOK)
		})
	})

	Method("CrowdRunsLittle", func() {
		Description("群众少跑腿")
		Meta("swagger:summary", "群众少跑腿")
		Payload(func() {
			Attribute("regionCode", String, "行政区划代码", func() {
				Example("520100")
			})
			Attribute("startDate", String, "起始时间", func() {
				Example("2019-07-27")
			})
			Attribute("endDate", String, "结束时间", func() {
				Example("2020-07-27")
			})
			Required("regionCode")
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
			Attribute("data", CrowdRunsLittleResp)
			Required("errcode", "errmsg", "data")
		})

		HTTP(func() {
			POST("/crowd_runs_little")
			Response(StatusOK)
		})
	})
})
