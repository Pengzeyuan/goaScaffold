package model

import (
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"
const ctLayout_nosec = "2006-01-02 15:04"
const ctLayout_date = "2006-01-02"

func (this *DateTime) UnmarshalJSON(b []byte) (err error) {

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sv := string(b)
	if len(sv) == 10 {
		sv += " 00:00:00"
	} else if len(sv) == 16 {
		sv += ":00"
	}
	this.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	if err != nil {
		if this.Time, err = time.ParseInLocation(ctLayout_nosec, string(b), loc); err != nil {
			this.Time, err = time.ParseInLocation(ctLayout_date, string(b), loc)
		}
	}

	return
}

func (this *DateTime) MarshalJSON() ([]byte, error) {
	//invalid character '-' after top-level value
	//开始没有认真看，以为“-”号不合法，就换了一个，结果错误一样：
	//
	//invalid character '/' after top-level value
	//看来，根本不是分割符的问题，仔细分析错误，发现“top-level”字样，我这返回的就是一个字符串，怎么可能top-level呢！想到这儿突然醒悟，是不是返回字符串应该自己加引号呢，急忙修改代码一试，果然！~_~
	//rs := []byte(this.Time.Format(ctLayout))
	rs := []byte(fmt.Sprintf(`"%s"`, this.Time.Format(ctLayout)))
	return rs, nil
}

var nilTime = (time.Time{}).UnixNano()

func (this *DateTime) IsSet() bool {
	return this.UnixNano() != nilTime
}

//然后，把结构中声明为time.Time的都修改为自定义的类型DateTime，试了一下，发现已经可以正确解析网页发来的时间，但是在输出时，
//总是不对，好像并没有调用自定义的Marshal方法。编写测试方法发现，原来json.Marshal方法调用DateTime.Marshal时出错了！
