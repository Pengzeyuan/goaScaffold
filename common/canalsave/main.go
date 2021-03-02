package main

import (
	"boot/common"
	"boot/model"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"reflect"
	"strings"
	"time"
)

func main() {
	//  连接nats
	err := common.ConnectNats()
	if err != nil {
		log.Println(err)

	}
	for {
		_, _ = common.NatsCli.QueueSubscribe("one_no_notice", "queue", func(m *nats.Msg) {
			//fmt.Printf("m的值 %s \n", m.Data)
			//s := string(m.Data)
			//fmt.Printf("s的字符串 %s \n", s)

		})
		// 这里是  收  数据库的某行变化得到结构体
		_, _ = common.NatsCli.QueueSubscribe("columnChange", "queue", func(m *nats.Msg) {
			fmt.Println(m.Data)
			fmt.Printf("map的值 %s %s \n", m.Data, reflect.TypeOf(m.Data))

			//转换字符串 为json string
			sprintf := string(m.Data)
			//str := strings.Replace(sprintf, "\"", "\\\"", -1)
			str := strings.Replace(sprintf, "\n", "", -1)
			//str = "\"" + str + "\""
			fmt.Println("jsonString:", str)

			//msg:="{\"cat_name\":\"咖1啡1大猫\",\"cat_price\":\"3.99\",\"cat_type\":\"3\",\"description\":\"又大又肥老人大人小孩都喜欢\",\"id\":\"3e171112112111161111c3\"}"
			//fmt.Println("有啥不一样:", str,"\\n",msg)
			//TransformfuncName(str)
			TransformAnimalsName(str)
			//var ptestStruct = *(**model.Cat)(unsafe.Pointer(&m.Data))
			//fmt.Printf("ptestStruct不安全指针: %s \n", ptestStruct)
			//cat := model.Cat{}
			//cat.ID=m.Data.
			//var ptestStruct = *(**model.Cat)(unsafe.Pointer(&m.Data))
			//fmt.Println("ptestStruct.CatName is : ", ptestStruct.CatName,"ptestStruct.CatPrice is :",ptestStruct.CatPrice)

			//marshal, err := json.Marshal( m.Data)
			//if err != nil {
			//	fmt.Printf("%s", "Marshal错误")
			//}

			// byteda:=[]byte{}
			//json.Unmarshal(byteda,columnMap )
			//cat, _ = json.Marshal(byteda)
			//for v:=range columnMap{
			//	columnMap
			//}

		})

		time.Sleep(30000 * time.Millisecond)
		//fmt.Printf("%s", "接受 服务")
	}

}

func Natsave1(m *nats.Msg) {
	fmt.Println(string(m.Data))
	//转换字符串 为json string
	sprintf := string(m.Data)
	str := strings.Replace(sprintf, "\"", "\\\"", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = "\"" + str + "\""
	fmt.Println("jsonString:", str)
	TransformfuncName(str)
}

func TransformfuncName(msg string) {
	var cat model.NatsCat
	if err := json.Unmarshal([]byte(msg), &cat); err == nil {
		fmt.Println(cat)
		fmt.Println(cat.CatName)
	} else {
		fmt.Println(err)
	}
}
func TransformAnimalsName(msg string) {
	var animal model.Animals
	if err := json.Unmarshal([]byte(msg), &animal); err == nil {
		fmt.Println(animal)
		fmt.Println(animal.Name)
	} else {
		fmt.Println(err)
	}
}
