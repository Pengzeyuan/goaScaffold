package config

type HallManagementDataConf struct {
	// 请求网址
	Url string `yaml:"url"`
	// pssId
	PassId string `yaml:"pass_id"`
	// passToken
	PassToken string `yaml:"pass_token"`
	// 超时时间
	TimeOut    string `yaml:"time_out"`
	TimeOutInt int    `yaml:"time_out_int" default:"5"`
}
