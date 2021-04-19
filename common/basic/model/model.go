package model

type Cat struct {
	//gorm.Model
	Id       int    `gorm:"type:int;not null;default:0;primary_key;column:'id'; comment:'id号'"`
	CatName  string `gorm:"type:VARCHAR(255);DEFAULT:'';column:'cat_name';  comment:'猫名称'"`
	CatPrice string `gorm:"type:VARCHAR(255);DEFAULT:'0';column:'cat_price'; comment:'猫价格'"`
	CatType  int    `gorm:"type:int(11);DEFAULT:0; column:'cat_type';  comment:'猫类型'" `
}

// Employee ...
type Employee struct {
	Id       int    `gorm:"type:int;not null; primary_key;column:id; comment:'id号'"`
	UserName string `gorm:"type:VARCHAR(255);DEFAULT:'';column:user_name;  comment:'名称'"`
	Age      int    `gorm:"type:int(11);DEFAULT:0; column:age;  comment:''" `
	Addr     string `gorm:"type:VARCHAR(255);DEFAULT:'0';column:addr; comment:'地址'"`
}

func (Employee) TableName() string {
	return "employee"
}
