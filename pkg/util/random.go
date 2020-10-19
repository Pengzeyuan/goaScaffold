package util

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
	"time"
)

// NOTE: go test 返回随机数不随机
func RandIntnV2(n int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(n))
	return result.Int64()
}

func RandIntn(n int) int {
	mathRand.Seed(time.Now().UnixNano())
	return mathRand.Intn(n)
}
