package test

import (
	"boot/dao"
	"boot/model"

	log "boot/gen/log"
	testpkg "boot/utils/testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"testing"
)

type DaoTestSuite struct {
	suite.Suite
	db         *gorm.DB
	zapLogger  *zap.Logger
	logger     *log.Logger
	TableNames []string
}

// suite中 tests 开始之前，先创建数据库表
func (tt *DaoTestSuite) SetupSuite() {

	tt.db = testpkg.DbCnnForTest(tt.T())
	tt.zapLogger = zap.L().With(zap.String("gzzwdp", "test"))
	tt.logger = log.New("test_dao", false)
}

func TestDaoTestSuite(t *testing.T) {
	suite.Run(t, new(DaoTestSuite))
}

////run test debug test
//func (tt *DaoTestSuite) Test_TaskBasicDao_StatDeptAndItems() {
//	taskBasicDao := dao.NewTaskBasicDaoImpl(dao.ItemsDB, tt.logger)
//	scopes := dao.NewTaskBasicScopeConstructor()
//	scopes = scopes.AddIsItem(1)
//	// scopes = scopes.AddRegionCode("520100")
//	res, err := taskBasicDao.StatDeptAndItems(scopes)
//	tt.NoError(err)
//	fmt.Printf("测试结果%+v", res)
//}

//测试 30天
//func (tt *DaoTestSuite) Test_30() {
//	taskGeneralDao := dao.NewProAcceptDaoImpl(dao.DoWorkDB, tt.zapLogger)
//
//	queryModel := model.CommonQueryModel{RegionCode: "520100", EndDate: "2019-03-27"}
//	res, err := taskGeneralDao.StatChannelTrend30Day(queryModel)
//	tt.NoError(err)
//
//	fmt.Printf("测试结果%+v", res[0])
//	for i, _ := range res {
//		fmt.Printf("测试结果%+v", res[i])
//	}
//
//}

// 测试 公共服务
func (tt *DaoTestSuite) Test_TaskGeneralDao_StatPubCountByTotalCount() {
	dao := dao.NewHallActualTimeDaoImpl(dao.DpDB, tt.zapLogger)

	info := model.WindowInfo{WindowNum: 20, WindowName: "洛丹伦"}
	err := dao.SaveWindowInfoData(&info)

	tt.NoError(err)

	//fmt.Printf("测试结果%+v \n", res)

}

// 测试 公共服务
func (tt *DaoTestSuite) Test_TaskGeneralDao_StatPubCountByTotalCount2() {
	windowInfo := model.WindowInfo{WindowNum: 20, WindowName: "洛丹伦"}
	dao.DpDB.Select(windowInfo).Where(queryDBInfo.QueryCriteria).Order(queryDBInfo.SortCriteria).Offset(queryDBInfo.Offset).Limit(queryDBInfo.Count).Find(&find)

	info := model.WindowInfo{WindowNum: 20, WindowName: "洛丹伦"}
	err := dao.SaveWindowInfoData(&info)

	tt.NoError(err)

	//fmt.Printf("测试结果%+v \n", res)

}
