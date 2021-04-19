package serializer

import (
	importfile "boot/gen/import_file"
	"boot/model"
)

func ReportGetFour2GetFour(info *model.FileCount) importfile.FourDo {
	t := importfile.FourDo{
		Year:  info.Year,
		Count: info.Count,
	}
	return t
}

func ModelGetFour2GetFour(fourDoInfos []*model.FileCount) []*importfile.FourDo {
	var res []*importfile.FourDo
	for i := 0; i < len(fourDoInfos); i++ {
		fourDoInfo := ReportGetFour2GetFour(fourDoInfos[i])
		res = append(res, &fourDoInfo)
	}
	return res
}

func ItemSplitRate2ReformOfAdministrativeResp(itemSplitRate *model.ItemSplitRate) importfile.ReformOfAdministrativeResp {
	res := importfile.ReformOfAdministrativeResp{
		SplitCount:     itemSplitRate.SplitCount,
		PastSplitCount: itemSplitRate.PastYearSplit,
		SplitRate:      itemSplitRate.SplitRate,
	}
	return res
}
