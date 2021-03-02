package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func StructToMapViaReflect() {
	m := make(map[string]interface{})
	t := time.Now()
	person := Persion{
		Id:       98439,
		Name:     "zhaondifnei",
		Address:  "大沙地",
		Email:    "dashdisnin@126.com",
		School:   "广州第十五中学",
		City:     "zhongguoguanzhou",
		Company:  "sndifneinsifnienisn",
		Age:      23,
		Sex:      "F",
		Proviece: "jianxi",
		Com:      "广州兰博基尼",
		PostTo:   "蓝鲸XXXXXXXX",
		Buys:     "shensinfienisnfieni",
		Hos:      "zhonsndifneisnidnfie",
	}
	elem := reflect.ValueOf(&person).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	fmt.Println(m)
	fmt.Printf("duration:%d", time.Now().Sub(t))
}


type CommonObj struct {
	Name           string    `persistence:"name"`
	Age            int       `persistence:"age"`
	LastUpdateTime time.Time `persistence:"lastUpdateTime"`
	score          float64   `persistence:"-"`
}

func StructConvertMapByTag(obj interface{}, tagName string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(tagName)
		fmt.Println(tagName)
		if tagName != "" && tagName != "-" {
			data[tagName] = v.Field(i).Interface()
		}
	}
	return data
}

func Struct2MapPassReflect() {
	obj := CommonObj{
		Name:           "aa",
		Age:            12,
		LastUpdateTime: time.Now(),
		score:          1.2,
	}

	m := StructConvertMapByTag(obj, "persistence")
	fmt.Println(m)

}
//下面是使用json进行转换的实例，可以看到，json在抓换的时候也是忽略没有暴露的字段的（小写字母开头的变量）
type CommonObj2 struct {
	Name           string    `persistence:"name" json:"name"`
	Age            int       `persistence:"age" json:"age"`
	LastUpdateTime time.Time `persistence:"lastUpdateTime" json:"last_update_time"`
	score          float64   `persistence:"-" json:"score"`
}

func StructConvertMapByTag2(obj interface{}, tagName string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get(tagName)
		fmt.Println(tagName)
		if tagName != "" && tagName != "-" {
			data[tagName] = v.Field(i).Interface()
		}
	}
	return data
}

func strut2MapPassJson() {
	obj := CommonObj2{
		Name:           "aa",
		Age:            12,
		LastUpdateTime: time.Now(),
		score:          1.2,
	}
	m := StructConvertMapByTag2(obj, "persistence")
	fmt.Println(m)
	bytes, _ := json.Marshal(obj)
	fmt.Println(string(bytes))
	var obj2 CommonObj2
	json.Unmarshal(bytes, &obj2)
	fmt.Println(obj2)
