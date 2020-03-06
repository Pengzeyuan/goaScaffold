package util

// 手机号码脱敏
func GetMaskedMobile(mobile string) string {
	if len(mobile) <= 10 {
		return mobile
	}

	return mobile[:3] + "****" + mobile[len(mobile)-4:]
}

// 姓名脱敏
func GetMaskedName(raw string) string {
	name := []rune(raw)
	if len(name) <= 1 {
		return raw + "*"
	} else if len(name) <= 2 {
		return string(name[0]) + "*"
	} else if len(name) <= 3 {
		return string(name[0]) + "*" + string(name[len(name)-1])
	}

	return string(name[0]) + "**" + string(name[len(name)-1])
}

func MaskIdNumber(id string) string {
	if len(id) <= 5 {
		return "***"
	}
	if len(id) <= 15 {
		return id[:3] + "***********"
	}
	// 510322198808081122
	return id[:3] + "************" + id[len(id)-3:]
}
