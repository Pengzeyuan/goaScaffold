package model

//大厅管理系统
type HallManagementSystemResponse struct {
	Data []struct {
		CardNum string `json:"cardnum"`
		Name    string `json:"name"`
		OuName  string `json:"ouname"`
	} `json:"userlist"`
}

// 法人用户列表
type LegalPersonUser struct {

	// 用户Id
	ID int `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY;INDEX ;comment:'用户Id'"` //
	// 名字
	Name string `gorm:"type:varchar(64);not null;default:'';comment:'名字'"json:"worker"` // 名字
	// 公司
	Companies []CompanyProfile `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
}

// 公司简洁
type CompanyProfile struct {
	// 行业Id
	Industry int `gorm:"TYPE:INT(11);DEFAULT:0"`
	// 公司名字
	Name string `gorm:"TYPE:VARCHAR(128);DEFAULT:'';"`
	// 法人名字
	UserID string `gorm:"TYPE:int(11);NOT NULL;INDEX"`
}
