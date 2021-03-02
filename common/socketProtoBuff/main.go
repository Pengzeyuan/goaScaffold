package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {

	//编码
	data, err := Marshal()
	if err != nil {
		fmt.Println("Marshal() error: ", err)
	}
	fmt.Println("Marshal:\n", data)

	//解码
	Unmarshal(data)
}

func Marshal() ([]byte, error) {

	var status UserStatus
	status = UserStatus_ONLINE

	userInfo := &UserInfo{
		Id:     proto.Int32(10),
		Name:   proto.String("XCL"),
		Status: &status,
	}

	return proto.Marshal(userInfo)
}

func Unmarshal(data []byte) {
	userInfo := &UserInfo{}

	err := proto.Unmarshal(data, userInfo)
	if err != nil {
		fmt.Println("Unmarshal() error: ", err)
	}

	fmt.Println("Unmarshal()\n userInfo:", userInfo)
}

/*
运行结果:
Marshal:
 [8 10 18 3 88 67 76 24 1]
Unmarshal()
 userInfo: id:10 name:"XCL" status:ONLINE
*/
