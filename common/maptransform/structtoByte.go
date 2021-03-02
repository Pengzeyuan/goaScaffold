package main

import (
	"boot/model"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"unsafe"
)

type TestStructTobytes struct {
	Data    int64
	StrData string
}

func (t TestStructTobytes) String() string {
	return fmt.Sprintf("{%d,%s}", t.Data, t.StrData)
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func main() {
	Struct2MapPassReflect()
	//json2Map2()
	//RawMessage()
	//RawMessage2()
	//Json2Map()
	//str2json()
	//jsonStrTObjTest()
	//TestStructToMap()
	//Bytetransformstruct()
	//BytetransformMap()
	//StructToMapViaJson()
	// testStruct := TestStructTobytes{100,"hello"}
	//Len := unsafe.Sizeof(testStruct)
	//testBytes := &SliceMock{
	//	addr: uintptr(unsafe.Pointer(&testStruct)),
	//	cap:  int(Len),
	//	len:  int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//fmt.Println("[]byte is : ", data)
	//
	//tobytes := TestStructTobytes{}
	//m := make(map[string]interface{})
	//json.Unmarshal(data,&tobytes)
	//json.Unmarshal(data,&m)
	//fmt.Printf("tobytes结果：%s %+v \n", tobytes, tobytes)
	//fmt.Printf("反射byte转结构结果：%s %+v \n", m, m)
	//marshal, _ := json.Marshal(testStruct)
	//json.Unmarshal(marshal, &m)
	//fmt.Println()
	//fmt.Println(m)
	//fmt.Printf("Marshal结果：%s %+v", m, m)
}

func Bytetransformstruct() {

	var testStruct = &TestStructTobytes{100, "ship"}
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", data)
	var ptestStruct = *(**TestStructTobytes)(unsafe.Pointer(&data))
	fmt.Println("ptestStruct.data is : ", ptestStruct.Data, "ptestStruct.StrData is :", ptestStruct.StrData)
}

// map 转  byte
func BytetransformMap() {

	mTestStruct := make(map[string]interface{})
	mTestStruct["Data"] = 100
	mTestStruct["StrData"] = "ship"

	var testStruct = TestStructTobytes{}
	err := mapstructure.Decode(mTestStruct, &testStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(testStruct)

	Len := unsafe.Sizeof(testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(&testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", data)
	var ptestStruct = *(**TestStructTobytes)(unsafe.Pointer(&data))
	fmt.Println("ptestStruct.data is : ", ptestStruct.Data, "ptestStruct.StrData is :", ptestStruct.StrData)
}

//struct转map例子
func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
func TestStructToMap() {
	student := model.Cat{CatName: "xian暹罗", CatPrice: "12", Description: "picheap"}
	data := StructToMapDemo(student)
	fmt.Println(data)
}
