package model

type NatsData struct {
	CallNoID  string // 叫号ID
	DataBytes []byte
}

type PushInfoEvent struct {
	CallNoID    string      `json:"callNoID"`    // 叫号ID
	InfoType    int         `json:"infoType"`    // 类型 1-窗口相关 2-取号相关 3-叫号相关
	InfoDetails interface{} `json:"infoDetails"` // 窗口详情
}
