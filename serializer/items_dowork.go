package serializer

import (
	importfile "boot/gen/import_file"
	"boot/model"
)

func LimitSceneNumELOneByAllStat2OneByAllStat(limitSceneNumELOneByAllStat *model.LimitSceneNumELOneByAllStat) importfile.CrowdRunsLittleResp {
	numbers := []*importfile.MatterNumber{}
	runOne := importfile.MatterNumber{
		BeforeAscension: limitSceneNumELOneByAllStat.PastRunOneCount,
		AfterAscension:  limitSceneNumELOneByAllStat.RunOneCount,
	}
	allItems := importfile.MatterNumber{
		BeforeAscension: limitSceneNumELOneByAllStat.PastTotalItemCount,
		AfterAscension:  limitSceneNumELOneByAllStat.TotalItemCount,
	}
	numbers = append(numbers, &runOne, &allItems)
	res := importfile.CrowdRunsLittleResp{
		MattersAccountedProportion: limitSceneNumELOneByAllStat.RunProportion,
		MattersAccounted:           numbers,
	}
	return res
}
