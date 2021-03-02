package thidPart

import (
	"boot/config"
	"boot/model"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
	"strings"
	"time"
)

type Client struct {
	logger *zap.Logger
	//thirdPartDB dao.ThirdPartDao
	httpClient *resty.Client
}

func NewClient(logger *zap.Logger, db *gorm.DB) *Client {
	timeout := time.Duration(config.C.HallManagementDataConf.TimeOutInt) * time.Second
	//thirdPartDB := dao.NewThirdPartDaoImpl(db, logger)
	httpClient := resty.New().SetTimeout(timeout)
	return &Client{
		logger:     logger,
		httpClient: httpClient,
		//thirdPartDB: thirdPartDB,
	}
}

//得到标记头
func GetSigedHeaders(passId string, paaSToken string) map[string]string {
	timestamp := time.Now().Unix()
	nonce := uuid.New().String()
	headers := map[string]string{
		"Accept":          "application/json",
		"x-tif-nonce":     nonce,
		"x-tif-paasid":    passId,
		"x-tif-timestamp": fmt.Sprintf("%d", timestamp),
	}
	signData := fmt.Sprintf("%d%s%s%d", timestamp, paaSToken, nonce, timestamp)
	headers["x-tif-signature"] = strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(signData))))
	return headers
}
func PostBodyRemoteUrl(url string, header map[string]string, msgData interface{}) ([]byte, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(msgData).
		SetHeaders(header).
		Post(url)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

func GetRemoteHallManagement(c Client) (*model.HallManagementSystemResponse, error) {
	var result model.HallManagementSystemResponse
	paramap := make(map[string]map[string]string)
	paramap["params"] = make(map[string]string)

	paramap["params"]["oucode"] = ""
	paramap["params"]["gonghao"] = ""

	header := GetSigedHeaders(config.C.HallManagementDataConf.PassId, config.C.HallManagementDataConf.PassToken)
	urlRsp, err := PostBodyRemoteUrl(config.C.HallManagementDataConf.Url, header, paramap)

	if err != nil {
		c.logger.Error("post body remote url is failed", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(urlRsp, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

//解析HallManagementResponse
func (c Client) CurrentHallManagementInfo() ([]model.HallManagementInfo, error) {

	// 拉取数据
	hallManagementResponse, err := GetRemoteHallManagement(c)
	if err != nil {
		c.logger.Error(" get remote lobby window info is failed", zap.Error(err))
	}
	infos := []model.HallManagementInfo{}
	for i := 0; i < len(hallManagementResponse.Data); i++ {
		hallManagementInform := model.HallManagementInfo{

			CardNum: hallManagementResponse.Data[i].CardNum,
			Name:    hallManagementResponse.Data[i].Name,
			OuName:  hallManagementResponse.Data[i].CardNum,
		}
		hallManagementInform.CreatedAt = time.Now()
		hallManagementInform.UpdatedAt = time.Now()
		//存储
		//err := c.thirdPartDB.SaveHallManagement(&hallManagementInform)
		infos = append(infos, hallManagementInform)
		if err != nil {
			c.logger.Error(" save hallManagementInform is failed", zap.Error(err))
		}
	}

	return infos, err
}
