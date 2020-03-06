package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaskedName(t *testing.T) {
	assert.Equal(t, "张*", GetMaskedName("张三"))
	assert.Equal(t, "王*娃", GetMaskedName("王二娃"))
	assert.Equal(t, "五**尚", GetMaskedName("五台山和尚"))
}

func TestMaskIdNumber(t *testing.T) {
	assert.Equal(t, "140************396", MaskIdNumber("140427200201014396"))
	assert.Equal(t, "140***********", MaskIdNumber("1404272002010"))
	assert.Equal(t, "***", MaskIdNumber("9999"))
}
