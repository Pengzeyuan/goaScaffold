package boot

import (
	"boot/config"
	"boot/dao"
	importfile "boot/gen/import_file"
	"boot/gen/log"
	mdlwr "boot/middleware"
	"boot/model"

	"boot/serializer"
	"boot/service"
	"boot/utils"

	//"boot/utils"
	"context"
	"errors"
	"fmt"
	"git.chinaopen.ai/yottacloud/go-libs/redis"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"path"
	"strings"
)

// import_file service example implementation.
// The example methods log the requests and return zero values.
type importFilesrvc struct {
	logger *log.Logger
}

// NewImportFile returns the import_file service implementation.
func NewImportFile(logger *log.Logger) importfile.Service {
	return &importFilesrvc{logger}
}

// 服务层
var FileUtilSVC = func(tx *gorm.DB, logger *zap.Logger, ctx context.Context) service.FileUtilService {
	return service.NewFileUtilSVCImpl(tx, logger, ctx, dao.NewFileUtilDaoImpl)
}

var itemsDoworksSVC = func(ctx context.Context, db *gorm.DB, db2 *gorm.DB, logger *zap.Logger) service.ItemsDoworkSVC {
	return service.NewItemsDoWorkSVCImpl(ctx, db, db2, *logger)
}

var simulationSVC = func(ctx context.Context, db *gorm.DB, logger *zap.Logger) service.SimulationSVC {
	return service.NewSimulationSVCImpl(ctx, db, logger)
}

func FileImportDecoderFunc(mr *multipart.Reader, p **importfile.ImportExcelFilePayload) error {
	// Add multipart request decoder logic here
	var fileName string

	fileMaxSize := config.C.UploadFile.Size
	maxMemory := fileMaxSize + int64(1<<20)
	form, err := mr.ReadForm(maxMemory)
	if err != nil {
		zap.L().Error("read form failed", zap.Error(err))
		return MakeBadRequest(err)
	}

	fileHandles, ok := form.File["file"]
	if !ok {
		zap.L().Error("user upload not a file")
		return MakeBadRequest(errors.New("没有上传任何文件"))
	}

	fileHandle := fileHandles[0]

	if fileHandle.Size > fileMaxSize {
		err := newFileSizeErr(fileMaxSize)
		zap.L().Error("file size is greater than file max size", zap.Error(err))
		return MakeBadRequest(err)
	}

	// 构造新文件名
	fileName = fileHandle.Filename
	newFileName := buildNewFileName(fileName)
	file, err := fileHandle.Open()
	if err != nil {
		zap.L().Error("open fileHandle failed", zap.Error(err))
		return MakeBadRequest(err)
	}

	// 读取文件
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		zap.L().Error("read file failed", zap.Error(err))
		if err := file.Close(); err != nil {
			zap.L().Error("close file failed", zap.Error(err))
			return MakeBadRequest(err)
		}
		return MakeBadRequest(err)
	}
	// 文件关闭
	if err := file.Close(); err != nil {
		zap.L().Error("close file failed", zap.Error(err))
		return MakeBadRequest(err)
	}

	*p = &importfile.ImportExcelFilePayload{
		File:     fileData,
		Filename: newFileName,
	}
	return nil
}

// excel数据批量导入，导入数据采用json格式存储
func (s *importFilesrvc) ImportExcelFile(ctx context.Context, p *importfile.ImportExcelFilePayload) (res *importfile.SuccessResult, err error) {
	res = &importfile.SuccessResult{}

	logger := L(ctx, s.logger)
	logger.Info("importFile.ImportExcelFile")

	tx := dao.DpDB.Begin()
	if tx.Error != nil {
		_ = tx.Rollback()
		logger.Error("begin tx failed", zap.Error(tx.Error))
		return nil, MakeInternalServerError(ctx, "内部服务器错误")
	}

	// 服务层
	svc := FileUtilSVC(tx, logger, ctx)

	// 全表加入
	if 1 != p.Type || 2 != p.Type || 3 != p.Type || 4 != p.Type {
		for i := 1; i <= 4; i++ {
			err = svc.Import(p.Year, i, p.Area, p.File)
			if err != nil {
				_ = tx.Rollback()
				logger.Error("import failed", zap.Error(err))
				return nil, MakeInternalServerError(ctx, "导入失败")
			}
		}
	} else {
		err = svc.Import(p.Year, p.Type, p.Area, p.File)
		if err != nil {
			_ = tx.Rollback()
			logger.Error("import failed", zap.Error(err))
			return nil, MakeInternalServerError(ctx, "导入失败")
		}
	}

	if err := tx.Commit().Error; err != nil {
		_ = tx.Rollback()
		logger.Error("commit tx failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "内部服务器错误")
	}
	res.OK = true

	return res, nil
}

// 获取插入excel数据统计信息
func (s *importFilesrvc) GetImportExcelFileInfo(ctx context.Context, p *importfile.GetImportExcelFileInfoPayload) (res *importfile.FourDoCountResp, err error) {
	res = &importfile.FourDoCountResp{}

	logger := L(ctx, s.logger)
	logger.Info("ImportExcelFile.GetImportExcelFileInfo")
	var regionCode string
	if p.Area == nil {
		regionCode = ""
	} else {
		regionCode = *p.Area
	}

	// 2.从业务层查询
	db := dao.DpDB
	svc := FileUtilSVC(db, logger, ctx)
	startYear := p.EndYear - 1
	modelFour, err := svc.GetFourCount(startYear, p.EndYear, regionCode)
	if err != nil {
		logger.Error("import failed", zap.Error(err))
		return nil, MakeInternalServerError(ctx, "获取四个办信息失败")
	}

	res.ImmediateInfoCount = serializer.ModelGetFour2GetFour(modelFour.ImmediateInfoCount)
	res.OnlineInfoCount = serializer.ModelGetFour2GetFour(modelFour.OnlineInfoCount)
	res.NearbyInfoCount = serializer.ModelGetFour2GetFour(modelFour.NearbyInfoCount)
	res.OnceInfoCount = serializer.ModelGetFour2GetFour(modelFour.OnceInfoCount)

	return res, nil
}

// 行政审批制度改革事项详情
func (s *importFilesrvc) ReformOfAdministrative(ctx context.Context, p *importfile.ReformOfAdministrativePayload) (res *importfile.ReformOfAdministrativeResult, err error) {
	res = &importfile.ReformOfAdministrativeResult{}

	logger := L(ctx, s.logger)
	logger.Info("importFile.ReformOfAdministrative")

	//构建查询体
	queryModel := model.CommonQueryModel{
		RegionCode: p.RegionCode,
	}
	if p.EndDate != nil {
		queryModel.EndDate = *p.EndDate
	}
	//构建Redis标记查询体
	//redisQueryModel := model.RedisQueryModel{ModelName: "reform", MethodName: "ReformOfAdministrative"}

	//// 1.从redis获取
	//storage := util.NewRedisStorage(redis.Client)
	//err = GetRedisData(&queryModel, redisQueryModel, &storage, &res.Data)
	//if err == nil {
	//	return res, nil
	//}

	//2.从业务层查询拿到对应服务和拿到Items 和  dowork数据源
	itemsDoWorksSVC := itemsDoworksSVC(ctx, dao.ItemsDB, dao.DoWorkDB, logger)

	// 事项大项拆分数量
	split, errSplit := itemsDoWorksSVC.StatReformItemSplit(queryModel)

	if errSplit != nil {
		logger.Error("stat reform item split result failed", zap.Error(errSplit))
		return nil, MakeInternalServerError(ctx, "获取事项大项拆分数量失败")
	}

	// 序列化为页面返回数据
	reformOfAdministrativeResp := serializer.ItemSplitRate2ReformOfAdministrativeResp(&split)

	//部门办件情况统计返回结果
	res.Data = &reformOfAdministrativeResp
	////3. 重新保存redis数据
	//SaveRedisData(&queryModel, redisQueryModel, (*utils.RedisStorage)(&storage), &res.Data)
	return res, nil
}

// 群众少跑腿
func (s *importFilesrvc) CrowdRunsLittle(ctx context.Context, p *importfile.CrowdRunsLittlePayload) (res *importfile.CrowdRunsLittleResult, err error) {

	res = &importfile.CrowdRunsLittleResult{}
	logger := L(ctx, s.logger)
	logger.Info("ImportExcelFile.CrowdRunsLittle")

	key, _ := ctx.Value(mdlwr.RequestPathKey).(string)
	svc := simulationSVC(ctx, dao.DpDB, logger)
	isMock, err := svc.GetByUnmarshal(key, &res.Data)

	if isMock == false {
		//构建查询体
		queryModel := model.CommonQueryModel{
			RegionCode: p.RegionCode,
		}
		if p.EndDate != nil {
			queryModel.EndDate = *p.EndDate
		}
		//构建Redis标记查询体
		redisQueryModel := model.RedisQueryModel{ModelName: "ImportExcelFile", MethodName: "CrowdRunsLittle"}
		// 1.从redis获取
		storage := utils.NewRedisStorage(redis.Client)
		err = GetRedisData(&queryModel, redisQueryModel, &storage, &res.Data)
		if err == nil {
			return res, nil
		}

		//2.从业务层查询拿到对应服务和拿到Items 和  dowork数据源
		itemsDoWorksSVC := itemsDoworksSVC(ctx, dao.ItemsDB, dao.DoWorkDB, logger)

		//群众少跑腿
		rate, errRate := itemsDoWorksSVC.StatRunOneRate(queryModel)

		if errRate != nil {
			logger.Error("stat reform item split result failed", zap.Error(errRate))
			return nil, MakeInternalServerError(ctx, "获取事项大项拆分数量失败")
		}
		// 序列化为页面返回数据
		limitSceneNumELOneByAllStat := serializer.LimitSceneNumELOneByAllStat2OneByAllStat(&rate)

		//部门办件情况统计返回结果
		res.Data = &limitSceneNumELOneByAllStat

		//3. 重新保存redis数据
		SaveRedisData(&queryModel, redisQueryModel, &storage, &res.Data)
	}

	return res, nil
}

// buildNewFileName 构建新文件名
// 新文件名由uuid+旧文件后缀组成
// 如果旧文件名是复合扩展名，则只取最后两个扩展名，作为新文件名的扩展名
// @param oldFileName 旧文件名
func buildNewFileName(oldFileName string) string {
	var fileExtension string
	fileExtension = path.Ext(oldFileName)
	oldFileName = strings.TrimSuffix(oldFileName, fileExtension)
	fileExtension = path.Ext(oldFileName) + fileExtension
	newFileName := fmt.Sprintf("%s%s", uuid.New(), fileExtension)
	return newFileName
}

// newFileSizeErr 构建FileSizeErr
// @param fileMaxSize 允许的最大文件大小 单位：byte
func newFileSizeErr(fileMaxSize int64) error {
	var errStr string
	switch {
	case fileMaxSize >= 1024*1024:
		errStr = fmt.Sprintf("文件大小不能超过 %d MB", fileMaxSize>>20)
	case fileMaxSize >= 1024:
		errStr = fmt.Sprintf("文件大小不能超过 %d KB", fileMaxSize>>10)
	default:
		errStr = fmt.Sprintf("文件大小不能超过 %d B", fileMaxSize)
	}
	return FileSizeErr{
		Message: errStr,
	}
}

// FileSizeErr 文件size 错误
type FileSizeErr struct {
	Message string
}

func (f FileSizeErr) Error() string {
	return f.Message
}
