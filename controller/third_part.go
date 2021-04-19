package boot

import (
	"boot/dao"
	log "boot/gen/log"
	thirdpart "boot/gen/third_part"
	helperthird "boot/helper/thidPart"
	"boot/model"
	"boot/service"
	"context"
	"encoding/json"

	"go.uber.org/zap"
)

//thirdPart service example implementation.
//The example methods log the requests and return zero values.
type thirdPartsrvc struct {
	logger *log.Logger
}

func (s *thirdPartsrvc) ReceiveThirdPartyPushData(ctx context.Context, payload *thirdpart.ReceiveThirdPartyPushDataPayload) (res *thirdpart.ReceiveThirdPartyPushDataResult, err error) {
	res = &thirdpart.ReceiveThirdPartyPushDataResult{}
	logger := L(ctx, s.logger)
	logger.Info("thirdPartsrvc.ReceiveThirdPartyPushData")

	svc := service.NewHallActualTimeSVCImpl(ctx, dao.DpDB, logger)

	marshal, err := json.Marshal(payload.Data)
	queryModel := model.CommonQueryModel{
		PullData: marshal,
		Method:   payload.MethodName,
		Count:    payload.Count,
	}

	actualTimeResp, err := svc.ReceiveThirdPartyPushData(queryModel)
	if err != nil {
		logger.Error("接收第三方推送数据失败", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "接收第三方推送数据失败")
	}
	res.Result = actualTimeResp.Msg

	return res, nil

}

func (s *thirdPartsrvc) GormRelatedSearch(ctx context.Context) (res *thirdpart.GormRelatedSearchResult, err error) {
	panic("implement me")
}

// NewThirdPart returns the thirdPart service implementation.
func NewThirdPart(logger *log.Logger) thirdpart.Service {
	return &thirdPartsrvc{logger}
}

// 接收大厅管理的数据
func (s *thirdPartsrvc) GetActualTimeData(ctx context.Context) (res *thirdpart.GetActualTimeDataResult, err error) {
	res = &thirdpart.GetActualTimeDataResult{}
	resps := []*thirdpart.HallManagementResp{}
	logger := L(ctx, s.logger)

	logger.Info("thirdPart.GetActualTimeData")
	//获取大厅管理数据
	cli := helperthird.NewClient(logger, dao.DpDB)

	//解析HallManagementResponse
	infos, err := cli.CurrentHallManagementInfo()

	for i := 0; i < len(infos); i++ {
		hallManagementInform := &thirdpart.HallManagementResp{
			CardNum: infos[i].CardNum,
			Name:    infos[i].Name,
			OuName:  infos[i].CardNum,
		}
		resps = append(resps, hallManagementInform)
	}

	res.Data = resps

	return res, err
}

// gorm关联查询
//func (s *thirdPartsrvc) GormRelatedSearch(ctx context.Context) (res *thirdpart.GormRelatedSearchResult, err error) {
//	res = &thirdpart.GormRelatedSearchResult{}
//	s.logger.Info("thirdPart.GormRelatedSearch")
//	logger := L(ctx, s.logger)
//	svc := service.NewThirdPartSVCImpl(ctx, dao.DpDB, logger)
//	queryModel := model.CommonQueryModel{}
//	search, err := svc.GormRelationSearch(queryModel)
//
//	// 序列化
//	legalUsers := []*thirdpart.LegalPersonUserResp{}
//	for i := 0; i < len(search); i++ {
//		ltemp := &thirdpart.LegalPersonUserResp{}
//		ltemp.Name = search[i].Name
//		ltemp.ID = search[i].ID
//
//		for j := 0; j < len(search[i].Companies); j++ {
//			ctemp := &thirdpart.CompanyProfileResp{}
//			ctemp.Name = search[i].Companies[j].Name
//			ctemp.Industry = search[i].Companies[j].Industry
//			ctemp.UserID = search[i].Companies[j].UserID
//			ltemp.Companies = append(ltemp.Companies, ctemp)
//		}
//		legalUsers = append(legalUsers, ltemp)
//	}
//	res.Data = legalUsers
//	return res, nil
//}
