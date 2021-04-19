package model

//到办事现场次数LimitSceneNum”小于等于1的事项，除以全部事项数目 *100%
type LimitSceneNumELOneByAllStat struct {
	RunProportion      float32 //现场一次办理的占总事项的比率 / 或者提升的比率
	RunOneCount        int32   //现场一次办理事项数
	TotalItemCount     int32   // 总事项数
	PastRunOneCount    int32   //去年现场一次办理事项数
	PastTotalItemCount int32   // 去年总事项数
}
