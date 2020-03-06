package wechat

import (
	"encoding/json"

	"gopkg.in/resty.v1"
)

const (
	baseURL = "https://api.weixin.qq.com"
)

type CheckRealNameInfoPayload struct {
	OpenID   string `json:"openid"`
	RealName string `json:"real_name"`
	CredId   string `json:"cred_id"`
	CredType string `json:"cred_type"`
	Code     string `json:"code"`
}

type CheckRealNameInfoResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	// V_OP_NA:用户暂未实名认证；V_OP_NM_MA:用户与姓名匹配；V_OP_NM_UM:用户与姓名不匹配。	有多个结果时用分号”;”连接；
	VerifyOpenid string `json:"verify_openid"`
	// 当verify_openid 为V_OP_NM_MA 时返回:V_NM_ID_MA:姓名与证件号匹配；V_NM_ID_UM:姓名与证件号不匹配。
	VerifyRealName string `json:"verify_real_name"`
}

// 小程序实名信息校验接口
func Checkrealnameinfo(accessToken string, p CheckRealNameInfoPayload) (*CheckRealNameInfoResult, error) {
	url := baseURL + "/intp/realname/checkrealnameinfo"
	p.CredType = "1"

	r, err := resty.New().R().SetPathParams(map[string]string{"access_token": accessToken}).
		SetBody(p).Post(url)
	if err != nil {
		return nil, err
	}

	result := CheckRealNameInfoResult{}
	if err := json.Unmarshal(r.Body(), &result); err != nil {
		return nil, err
	}

	return &result, nil
}
