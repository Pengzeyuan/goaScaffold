package serializer

import (
	actualtime "boot/gen/actual_time"
	"boot/model"
)

func CanalData2CanalData(d *model.CanalData) *actualtime.CanalDataResp {
	m := &actualtime.CanalDataResp{
		DataType: int32(d.DataType),
		DataInfo: d.InfoDetails,
	}
	return m
}
