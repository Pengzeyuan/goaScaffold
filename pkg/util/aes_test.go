package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesEncrypt(t *testing.T) {

	testIdcard := "510823198808093322"
	key := []byte("0123456789abcdef")

	crypted, err := AesEncrypt([]byte(testIdcard), key)
	assert.Equal(t, nil, err)

	raw, err := AesDecrypt(crypted, key)
	assert.Equal(t, nil, err)

	assert.Equal(t, testIdcard, string(raw))
}
