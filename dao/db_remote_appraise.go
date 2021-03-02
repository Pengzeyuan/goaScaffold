package dao

//
//import (
//	"fmt"
//	libsgorm "git.chinaopen.ai/yottacloud/go-libs/gorm"
//	"github.com/jinzhu/gorm"
//	"go.uber.org/zap"
//	systemconfiguration "gzzwdp/gen/system_configuration"
//	"gzzwdp/model"
//	"strings"
//)
//
//type RemoteAppraiseDao interface {
//
//	// 拉取第三方的评价系统统计总数到本地数据库
//	SaveAppraiseSummery(value *model.AppraiseSummeryCount) error
//	// 拉取第三方的评价系统部门评价到本地数据库
//	SaveAppraiseDepartment(value *model.DepartmentAvgScore) error
//	// 拉取第三方的评价系统用户评价到本地数据库
//	SaveAppraiseStaff(value *model.UserAvgScore) error
//	// 批量拉取第三方的评价系统用户评价到本地数据库
//	BatchSaveAppraiseStaff(values []*model.UserAvgScore) error
//	// 获取大厅管理系统情况
//	SaveHallManagement(value *model.HallManagementInfo) error
//	//  筛选着存储好差评部门表的数据
//	SaveGoodBadDepartmentTable(value model.GoodBadDepartment) error
//	//  清空 好差评部门表的数据 这是临时表
//	TRUNCATEGoodBadDepartmentTable() error
//	// 获取系统配置大厅部门关联列表
//	GetHallDepartmentAssociationList() ([]*model.HallManagementInfo, error)
//	// 获取系统配置好差评部门关联列表
//	GetGoodBadDepartmentAssociationList() ([]*model.GoodBadDepartment, error)
//
//	// 设置大厅部门列表 和 好差评部门列表的关联
//	SetHallDepartment2GoodBadDepartment([]*systemconfiguration.HallDepartment2GoodBadDepartment) error
//
//	// 得到职员信息列表
//	GetStaffInfo() ([]*model.UserAvgScore, error)
//	// 得到部门信息列表
//	GetDepartmentInfo() ([]*model.DepartmentAvgScore, error)
//
//	// 得到大厅总评分列表
//	GetSummeryAppraiseInfo(queryModel model.CommonQueryModel) (*model.AppraiseSummeryCount, error)
//
//	// 得到办事人员评分列表
//	GetStaffAppraiseInfo(queryModel model.CommonQueryModel) ([]*model.UserAvgScore, error)
//
//	// 得到办事人员评分列表
//	GetStaffAppraiseEndInfo(queryModel model.CommonQueryModel) ([]*model.UserAvgScore, *libsgorm.Page, error)
//	//
//	//// 得到办理部门评分列表
//	//GetDepartmentAppraiseInfo()(error)
//
//}
//type RemoteAppraiseDaoImpl struct {
//	db     *gorm.DB
//	logger *zap.Logger
//}
//type NewRemoteAppraiseDaoFunc = func(*gorm.DB, *zap.Logger) RemoteAppraiseDao
//
//func NewRemoteAppraiseDaoImpl(db *gorm.DB, logger *zap.Logger) RemoteAppraiseDao {
//	return RemoteAppraiseDaoImpl{db: db, logger: logger}
//}
//
//// 得到办事人员评分列表
//func (r RemoteAppraiseDaoImpl) GetStaffAppraiseInfo(queryModel model.CommonQueryModel) ([]*model.UserAvgScore, error) {
//
//	res := []*model.UserAvgScore{}
//
//	querySQL := `
//		SELECT t_a.avgdf , t_a.user_id,t_a.user_name,t_a.belong_depart,t_a.year_month_only, (t_a.avgdf-t_b.avgdf) as  avg_change
//		from
//		(
//		select avgdf , user_id,user_name,belong_depart,year_month_only
//		FROM user_avg_score
//		Where year_month_only = ?
//		) t_a
//		JOIN
//		(
//		SELECT avgdf , user_id,user_name,belong_depart,year_month_only
//		FROM user_avg_score
//		where year_month_only = ?
//		) t_b
//		ON t_a.user_id=t_b.user_id
//		ORDER BY t_a.avgdf  %s
//    `
//	querySQL = fmt.Sprintf(querySQL, queryModel.ArgsData)
//	err := r.db.Raw(querySQL, queryModel.StartDate, queryModel.EndDate).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("GetStaffAppraiseInfo failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//func NewRemoteScopeConstructor() RemoteScopeConstructor {
//	return RemoteScopeConstructor{
//		scopes: []libsgorm.Scope{},
//	}
//}
//
//// scopes相关
//type RemoteScopeConstructor struct {
//	scopes []libsgorm.Scope
//}
//
//func (s RemoteScopeConstructor) AddOrderByCreateAtDesc() RemoteScopeConstructor {
//	query := func(db *gorm.DB) *gorm.DB {
//		return db.Order("create_at DESC")
//	}
//	s.scopes = append(s.scopes, query)
//	return s
//}
//func (s RemoteScopeConstructor) AddCityType(cityType int) RemoteScopeConstructor {
//	query := func(db *gorm.DB) *gorm.DB {
//		return db.Where("type = ?", cityType)
//	}
//	s.scopes = append(s.scopes, query)
//	return s
//}
//
//// 获取scopes
//func (s RemoteScopeConstructor) Scopes() []libsgorm.Scope {
//	return s.scopes
//}
//
//// 得到办事人员评分列表
//func (r RemoteAppraiseDaoImpl) GetStaffAppraiseEndInfo(queryModel model.CommonQueryModel, scopes RemoteScopeConstructor) ([]*model.UserAvgScore, *libsgorm.Page, error) {
//
//	var res []model.UserAvgScore
//	db := r.db.Scopes(scopes.Scopes()...)
//	query := FilterDeleted(db)
//	page, err := libsgorm.Pagination(query, limit, cursor, &res)
//	if err != nil {
//		r.logger.Error("list cities for pagination failed", zap.Error(err))
//		return res, nil, err
//	}
//	return res, page, nil
//
//	res := []*model.UserAvgScore{}
//
//	querySQL := `
//		SELECT t_a.avgdf , t_a.user_id,t_a.user_name,t_a.belong_depart,t_a.year_month_only, (t_a.avgdf-t_b.avgdf) as  avg_change
//		FROM
//		(
//		SELECT avgdf , user_id, user_name, belong_depart,year_month_only
//		FROM user_avg_score
//		WHERE year_month_only = ?
//		) t_a
//		JOIN
//		(
//		SELECT avgdf , user_id, user_name, belong_depart,year_month_only
//		FROM user_avg_score
//		WHERE year_month_only = ?
//		) t_b
//		ON t_a.user_id=t_b.user_id
//		ORDER BY  %s   %s
//    `
//
//	if queryModel.ArgsSecondData != "" {
//		querySQL = fmt.Sprintf(querySQL, queryModel.ArgsData, queryModel.ArgsSecondData)
//	} else {
//		querySQL = fmt.Sprintf(querySQL, queryModel.ArgsData)
//	}
//
//	scan := r.db.Raw(querySQL, queryModel.StartDate, queryModel.EndDate).Scan(&res)
//	query := FilterDeleted(scan)
//	page, err := libsgorm.Pagination(query, queryModel.Limit, queryModel.Courser, &res)
//	if err != nil {
//		r.logger.Error("GetStaffAppraiseInfo failed", zap.Error(err))
//		return nil, nil, err
//	}
//
//	return res, page, nil
//}
//
//// 得到大厅总评分列表
//func (r RemoteAppraiseDaoImpl) GetSummeryAppraiseInfo(queryModel model.CommonQueryModel) (*model.AppraiseSummeryCount, error) {
//	res := &model.AppraiseSummeryCount{}
//
//	querySQL := `
//	SELECT zhpf,   pjrycount, fcmyl, myl, jbmyl, bmyl , fcbmyl,year_month_only
//    FROM appraise_summery_count
//    WHERE year_month_only = ?
//	`
//
//	err := r.db.Raw(querySQL, queryModel.StartDate).Scan(res).Error
//
//	if err != nil {
//		r.logger.Error("get do_count_analysis failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//// 获取系统配置好差评部门关联列表
//func (r RemoteAppraiseDaoImpl) GetGoodBadDepartmentAssociationList() ([]*model.GoodBadDepartment, error) {
//	res := []*model.GoodBadDepartment{}
//
//	querySQL := `
//		SELECT o_name , o_code
//		FROM good_bad_department
//		Where o_code  like "%520100%"
//			OR  o_name like "%贵阳%"
//		GROUP BY o_code
//    `
//
//	err := r.db.Raw(querySQL).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("GetGoodBadDepartmentAssociationList failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//// 获取系统配置大厅部门关联列表
//func (r RemoteAppraiseDaoImpl) GetHallDepartmentAssociationList() ([]*model.HallManagementInfo, error) {
//	res := []*model.HallManagementInfo{}
//
//	querySQL := `
//	SELECT ou_code,ou_name,department_id,department_name
//	FROM hall_management_info
//    GROUP BY ou_code
//    `
//
//	err := r.db.Raw(querySQL).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("GetHallDepartmentList failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//// 拉取第三方的评价系统统计总数到本地数据库
//func (r RemoteAppraiseDaoImpl) SaveAppraiseSummery(value *model.AppraiseSummeryCount) error {
//
//	if err := r.db.Exec(
//		"INSERT INTO appraise_summery_count (`created_at`,`updated_at`,`deleted_at`,"+
//			"`zhpf`, `pjrycount`, `fcmyl`,`myl`, `jbmyl`, `bmyl`, `fcbmyl`, `is_search_year_or_month`, `evaluate_time`,`year_month_only`,`region_code`)"+
//			"VALUES (?,?,?,?,?,  ?,?,?,?,?,  ?,?,?,?) "+
//			"ON DUPLICATE KEY UPDATE "+
//			"updated_at = ?, deleted_at = ?, zhpf = ?, pjrycount = ?, fcmyl = ?, myl = ?, jbmyl = ?, bmyl = ?, "+
//			"fcbmyl = ?, is_search_year_or_month= ? ,evaluate_time = ?, year_month_only=?,region_code= ?",
//
//		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.Zhpf, value.Pjrycount, value.Fcbmyl,
//		value.Myl, value.Jbmyl, value.Bmyl, value.Fcbmyl, value.IsSearchYearOrMonth, value.EvaluateTime,
//		value.YearMonthOnly, value.RegionCode,
//
//		value.Model.UpdatedAt, value.Model.DeletedAt, value.Zhpf, value.Pjrycount, value.Fcbmyl,
//		value.Myl, value.Jbmyl, value.Bmyl, value.Fcbmyl, value.IsSearchYearOrMonth, value.EvaluateTime,
//		value.YearMonthOnly, value.RegionCode,
//	).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 拉取第三方的评价系统部门评价到本地数据库
//func (r RemoteAppraiseDaoImpl) SaveAppraiseDepartment(value *model.DepartmentAvgScore) error {
//
//	if err := r.db.Exec(
//		"INSERT INTO department_avg_score (`created_at`,`updated_at`,`deleted_at`,`avgdf`, `department_id`,  "+
//			"  `is_search_year_or_month`, `evaluate_time`,`year_month_only`,`region_code`,`department_name`)"+
//			"VALUES (?,?,?,?,?,  ?,?,?,?,?) "+
//			"ON DUPLICATE KEY UPDATE "+
//			"updated_at = ?, deleted_at = ?, avgdf = ?, evaluate_time = ?,region_code = ?,department_name = ?",
//
//		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.Avgdf, value.DepartmentId,
//		value.IsSearchYearOrMonth, value.EvaluateTime, value.YearMonthOnly, value.RegionCode, value.DepartmentName,
//
//		value.Model.UpdatedAt, value.Model.DeletedAt, value.Avgdf, value.EvaluateTime, value.RegionCode, value.DepartmentName,
//	).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 拉取第三方的评价系统部门评价到本地数据库
//func (r RemoteAppraiseDaoImpl) SaveAppraiseStaff(value *model.UserAvgScore) error {
//
//	if err := r.db.Exec(
//		"INSERT INTO user_avg_score (`created_at`,`updated_at`,`deleted_at`,`avgdf`, `user_id`,  "+
//			" `region_code`, `is_search_year_or_month`, `evaluate_time`,`year_month_only`,`user_name`,`belong_depart`)"+
//			"VALUES (?,?,?,?,?,  ?,?,?,?,?  ,?) "+
//			"ON DUPLICATE KEY UPDATE "+
//			"updated_at = ?, deleted_at = ?, avgdf = ?, region_code = ?, evaluate_time = ?,user_name=?,belong_depart=?",
//
//		value.Model.CreatedAt, value.Model.UpdatedAt, value.Model.DeletedAt, value.Avgdf, value.UserId, value.RegionCode,
//		value.IsSearchYearOrMonth, value.EvaluateTime, value.YearMonthOnly, value.UserName, value.BelongDepart,
//
//		value.Model.UpdatedAt, value.Model.DeletedAt, value.Avgdf, value.RegionCode, value.EvaluateTime, value.UserName, value.BelongDepart,
//	).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 批量拉取第三方的评价系统部门评价到本地数据库
//func (r RemoteAppraiseDaoImpl) BatchSaveAppraiseStaff(values []*model.UserAvgScore) error {
//
//	//  批量保存到中间库
//	sqlStr := "INSERT INTO user_avg_score( created_at,updated_at, deleted_at, avgdf,user_id,region_code,is_search_year_or_month," +
//		"evaluate_time,year_month_only,user_name,belong_depart) VALUES "
//	vals := []interface{}{}
//	const rowSQL = "(?,?,?,?,?,  ?,?,?,?,? ,?)"
//	var inserts []string
//	for _, elem := range values {
//		inserts = append(inserts, rowSQL)
//		vals = append(vals, &elem.CreatedAt, &elem.UpdatedAt, &elem.DeletedAt, &elem.Avgdf, &elem.UserId, &elem.RegionCode,
//			&elem.IsSearchYearOrMonth, &elem.EvaluateTime, &elem.YearMonthOnly, &elem.UserName, &elem.BelongDepart)
//	}
//	sqlStr = sqlStr + strings.Join(inserts, ",") + "ON DUPLICATE KEY UPDATE " +
//		"updated_at =VALUES(updated_at), avgdf = VALUES(avgdf), region_code =VALUES(region_code), " +
//		"evaluate_time = VALUES(evaluate_time),user_name= VALUES(user_name),belong_depart= VALUES(belong_depart)"
//	if err := r.db.Exec(sqlStr, vals...).Error; err != nil {
//		return err
//	}
//
//	return nil
//
//}
//
////  存储大厅窗口设立情况信息接口数据
//func (r RemoteAppraiseDaoImpl) SaveHallManagement(value *model.HallManagementInfo) error {
//	if err := r.db.Exec(
//		"INSERT INTO hall_management_info( staff_name,ou_name,card_num,ou_code) "+
//			"VALUES (?, ?, ?, ?) "+
//			"ON DUPLICATE KEY UPDATE   staff_name = ?, ou_name = ? , card_num = ?,ou_code = ? ",
//
//		value.StaffName, value.OuName, value.CardNum, value.OuCode,
//
//		value.StaffName, value.OuName, value.CardNum, value.OuCode,
//	).Error; err != nil {
//		r.logger.Error(" Save HallManagement is failed", zap.Error(err))
//		return err
//	}
//	return nil
//}
//
////  清空 好差评部门表的数据 这是临时表
//func (r RemoteAppraiseDaoImpl) TRUNCATEGoodBadDepartmentTable() error {
//	if err := r.db.Exec("TRUNCATE TABLE good_bad_department").Error; err != nil {
//		r.logger.Error(" TRUNCATE TABLE good_bad_department is failed", zap.Error(err))
//		return err
//	}
//	return nil
//}
//
////  筛选着存储好差评部门表的数据
//func (r RemoteAppraiseDaoImpl) SaveGoodBadDepartmentTable(value model.GoodBadDepartment) error {
//
//	if err := r.db.Exec(
//		"INSERT INTO good_bad_department (staff_id,staff_name,staff_avg,o_code,o_name) "+
//			"VALUES (?,?,?,?,?) "+
//			"ON DUPLICATE KEY UPDATE   staff_id = ?, staff_name = ?,staff_avg= ?, o_code=? ,o_name=?",
//
//		value.StaffId, value.StaffName, value.StaffAvg, value.OCode, value.OName,
//
//		value.StaffId, value.StaffName, value.StaffAvg, value.OCode, value.OName,
//	).Error; err != nil {
//		r.logger.Error(" SaveStaffAssociationTable is failed", zap.Error(err))
//		return err
//	}
//	return nil
//
//}
//func (r RemoteAppraiseDaoImpl) GetGoodBadDepartmentId(queryModel model.CommonQueryModel) ([]*model.GoodBadDepartment, error) {
//	res := []*model.GoodBadDepartment{}
//	querySQL := `
//		SELECT o_code
//		FROM good_bad_department
//		WHERE o_name in (?)
//		GROUP BY o_code
//		ORDER BY FIND_IN_SET(o_name,'%s')
//   `
//	querySQL = fmt.Sprintf(querySQL, queryModel.ArgsData)
//
//	err := r.db.Raw(querySQL, queryModel.Scope).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("get do_count_analysis failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//// 设置大厅部门列表 和 好差评部门列表的关联
//func (r RemoteAppraiseDaoImpl) SetHallDepartment2GoodBadDepartment(departments []*systemconfiguration.HallDepartment2GoodBadDepartment) error {
//
//	querySQL := `
//		UPDATE hall_management_info hall
//		SET hall.department_name =  ? ,hall.department_id= IFNULL(   ?	,'')
//		WHERE   hall.ou_code = ? AND  hall.ou_name = ?
//	`
//
//	for i := 0; i < len(departments); i++ {
//		err := r.db.Exec(querySQL, departments[i].GoodBadDepartment, departments[i].GoodBadDepartmentID,
//			departments[i].HallDepartmentID, departments[i].HallDepartment).Error
//		if err != nil {
//			r.logger.Error("SetHallDepartment2GoodBadDepartment failed", zap.Error(err))
//			return err
//		}
//	}
//
//	return nil
//
//}
//
//// 得到职员信息列表
//func (r RemoteAppraiseDaoImpl) GetStaffInfo() ([]*model.UserAvgScore, error) {
//	res := []*model.UserAvgScore{}
//
//	querySQL := `
//		SELECT good.staff_avg as avgdf, good.staff_id as user_id,  hall.staff_name as user_name,hall.ou_name as belong_depart
//		FROM good_bad_department good
//		JOIN hall_management_info hall
//		ON good.o_code=hall.department_id
//		GROUP BY hall.card_num
//   `
//
//	err := r.db.Raw(querySQL).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("GetStaffInfo failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
//
//// 得到部门信息列表
//func (r RemoteAppraiseDaoImpl) GetDepartmentInfo() ([]*model.DepartmentAvgScore, error) {
//	res := []*model.DepartmentAvgScore{}
//
//	querySQL := `
//		SELECT hall.department_id as department_id, hall.department_name as department_name
//		FROM hall_management_info hall
//		WHERE IFNULL(hall.department_id,'kong' )!='kong'
// 				AND hall.department_id !=' '
//		GROUP BY hall.department_id
//   `
//
//	err := r.db.Raw(querySQL).Scan(&res).Error
//
//	if err != nil {
//		r.logger.Error("GetDepartmentInfo failed", zap.Error(err))
//		return nil, err
//	}
//	return res, nil
//}
