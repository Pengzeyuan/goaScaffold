package main

import (
	"boot/model"
	"encoding/json"
	"fmt"
	"strings"
)

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

type student struct {
	people
	ID int `json:"sid"`
}

func jsonStrTObjTest() {
	msg := "{\"name\":\"zhangsan\", \"age\":18, \"id\":122463, \"sid\":122464}"
	var someOne student
	if err := json.Unmarshal([]byte(msg), &someOne); err == nil {
		fmt.Println(someOne)
		fmt.Println(someOne.people)
	} else {
		fmt.Println(err)
	}
	msg = "{\"cat_name\":\"咖1啡1大猫\",\"cat_price\":\"3.99\",\"cat_type\":\"3\",\"description\":\"又大又肥老人大人小孩都喜欢\",\"id\":\"3e171112112111161111c3\"}"

	TransformfuncName(msg)
	fmt.Println("MSG2")
	msg2 := `{"cat_name":"咖1啡1大猫","cat_price":"3.99","cat_type":"3","description":"又大又肥老人大人小孩都喜欢","id":"3e17111213131231135c3"}`
	TransformfuncName(msg2)
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

func str2json() {
	cmd := "[{'read': 2.0, 'write': 1.2}, {'read_mb': 4.0, 'write': 3.2}]"
	str := strings.Replace(string(cmd), "'", "\"", -1)
	str = strings.Replace(str, "\n", "", -1)

	var dat []map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		//fmt.Println(dat["status"])
	} else {
		fmt.Println(err)
	}
}

func Json2Map() {
	print := `[{"read": 2.0 ,"write": 1.2}, {"read_mb": 4.0, "write": 3.2}]`

	str := strings.Replace(string(print), "'", "\"", -1)
	str = strings.Replace(str, "\n", "", -1)

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
		fmt.Println(dat["status"])
	} else {
		fmt.Println(err)
	}
}

type Iot struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Context json.RawMessage `json:"context"` // RawMessage here! (not a string)
}
type Iot2 struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Context string `json:"context"` // RawMessage here! (not a string)
}

func RawMessage() {
	in := []byte(`{"id":1,"name":"test","context":{"key1":"value1","key2":2}}`)

	var iot Iot
	err := json.Unmarshal(in, &iot)
	if err != nil {
		panic(err)
	}

	// Context is []byte, so you can keep it as string in DB
	fmt.Println("ctx:", string(iot.Context))

	// Marshal back to json (as original)
	out, _ := json.Marshal(&iot)
	fmt.Println(string(out))
}
func RawMessage2() {
	in := []byte(`{"id":1,"name":"test","context":{"key1":"value1","key2":2}}`)

	var iot Iot2
	err := json.Unmarshal(in, &iot)
	if err != nil {
		panic(err)
	}

	// Context is []byte, so you can keep it as string in DB
	fmt.Println("ctx:", string(iot.Context))

	// Marshal back to json (as original)
	out, _ := json.Marshal(&iot)
	fmt.Println(string(out))
}

func json2Map2() {
	type Response struct {
		RequestID     string                   `json:"RequestId"`
		SendStatusSet []map[string]interface{} `json:"SendStatusSet"`
	}
	type r struct {
		Response Response `json:"Response"`
	}
	txt := `{
        "Response": {
            "SendStatusSet": [{
                    "SerialNo": "5000:1045710669157053657849499619",
                    "PhoneNumber": "+8618511122233",
                    "Fee": 1,
                    "SessionContext": "test",
                    "Code": "Ok",
                    "Message": "send success",
                    "IsoCode": "CN"
                },
                {
                    "SerialNo": "5000:104571066915705365784949619",
                    "PhoneNumber": "+8618511122266",
                    "Fee": 1,
                    "SessionContext": "test",
                    "Code": "Ok",
                    "Message": "send success",
                    "IsoCode": "CN"
                }
            ],
            "RequestId": "a0aabda6-cf91-4f3e-a81f-9198114a2279"
        }
    }`
	// fmt.Println(txt)
	p := &r{}
	err := json.Unmarshal([]byte(txt), p)
	fmt.Println(err)
	fmt.Println(*p)

	type simple struct {
		Response map[string]interface{}
	}
	zzz := new(simple)
	err = json.Unmarshal([]byte(txt), zzz)
	fmt.Println(err)
	fmt.Println("--------------")
	fmt.Println(*zzz)

	simpleJSON := `{"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}`
	type k struct {
		Name string
	}
	kk := &k{}
	err = json.Unmarshal([]byte(simpleJSON), kk)
	fmt.Println(err)
	fmt.Println(*kk)
}
