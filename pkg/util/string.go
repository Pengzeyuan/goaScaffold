package util

import (
	"strings"
	"time"
)

var (
	idCard_Coefficient []int32 = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	idCard_code        []byte  = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

func GetString(ptr *string) string {
	if ptr == nil {
		return ""
	}

	return *ptr
}

func CurrentTimestamp() int64 {
	return int64(time.Now().UnixNano() / int64(time.Millisecond))
}

func FindString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// 校验一个身份证是否是合法的身份证
func CheckIdCard(idCardNo string) bool {
	if len(idCardNo) != 18 {
		return false
	}

	idByte := []byte(strings.ToUpper(idCardNo))

	sum := int32(0)
	for i := 0; i < 17; i++ {
		sum += int32(byte(idByte[i])-byte('0')) * idCard_Coefficient[i]
	}
	return idCard_code[sum%11] == idByte[17]
}
