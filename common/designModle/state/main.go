package main

//------------------------------------------------------------
//我的代码没有`else`系列
//状态模式
//@auhtor TIGERB<https://github.com/TIGERB>
//------------------------------------------------------------

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Context 上下文
type Context struct {
	Tel        string // 手机号
	Text       string // 短信内容
	TemplateID string // 短信模板ID
}

// SmsServiceInterface 短信服务接口
type SmsServiceInterface interface {
	Send(ctx *Context) error
}

// ServiceProviderAliyun 阿里云
type ServiceProviderAliyun struct {
}

// Send Send
func (s *ServiceProviderAliyun) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【阿里云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// ServiceProviderTencent 腾讯云
type ServiceProviderTencent struct {
}

// Send Send
func (s *ServiceProviderTencent) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【腾讯云】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// ServiceProviderYunpian 云片
type ServiceProviderYunpian struct {
}

// Send Send
func (s *ServiceProviderYunpian) Send(ctx *Context) error {
	fmt.Println(runFuncName(), "【云片】短信发送成功，手机号:"+ctx.Tel)
	return nil
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// ProviderType 短信服务提供商类型
type ProviderType string

const (
	// ProviderTypeAliyun 阿里云
	ProviderTypeAliyun ProviderType = "aliyun"
	// ProviderTypeTencent 腾讯云
	ProviderTypeTencent ProviderType = "tencent"
	// ProviderTypeYunpian 云片
	ProviderTypeYunpian ProviderType = "yunpian"
)

var (
	// stateManagerInstance 当前使用的服务提供商实例
	// 默认aliyun
	stateManagerInstance *StateManager
)

// StateManager 状态管理
type StateManager struct {
	// CurrentProviderType 当前使用的服务提供商类型
	// 默认aliyun
	currentProviderType ProviderType // 有三个状态

	// CurrentProvider 当前使用的服务提供商实例
	// 默认aliyun
	currentProvider SmsServiceInterface //  短信服务接口  有send方法

	// 更新状态时间间隔
	setStateDuration time.Duration //
	// setState  设置状态  函数
}

// initState 初始化状态
func (m *StateManager) initState(duration time.Duration) {
	// 初始化
	m.setStateDuration = duration
	// 拿到一个服务提供商
	m.setState(time.Now())

	// 定时器更新状态
	go func() {
		for {
			// 每一段时间后根据回调的发送成功率 计算得到当前应该使用的 厂商
			select {
			//  时间滴答  设置 300毫秒 一次   只要能取出数据 就执行  set State  取出的就是当前时间
			case t := <-time.NewTicker(m.setStateDuration).C:
				m.setState(t) //
			}
		}
	}()
}

// setState 设置状态
// 根据短信云商回调的短信发送成功率 得到下阶段发送短信使用哪个厂商的服务
func (m *StateManager) setState(t time.Time) {
	// 这里用随机模拟   一个数组
	ProviderTypeArray := [3]ProviderType{
		ProviderTypeAliyun,
		ProviderTypeTencent,
		ProviderTypeYunpian,
	}
	m.currentProviderType = ProviderTypeArray[rand.Intn(len(ProviderTypeArray))]
	// 选择一个特定的 服务商接口
	switch m.currentProviderType {
	case ProviderTypeAliyun:
		m.currentProvider = &ServiceProviderAliyun{}
	case ProviderTypeTencent:
		m.currentProvider = &ServiceProviderTencent{}
	case ProviderTypeYunpian:
		m.currentProvider = &ServiceProviderYunpian{}
	default:
		panic("无效的短信服务商")
	}

	//   显示一个时间
	fmt.Printf("时间：%s| 变更短信发送厂商为: %s \n", t.Format("2006-01-02 15:04:05"), m.currentProviderType)
}

// getState 获取当前状态
func (m *StateManager) getState() SmsServiceInterface {
	return m.currentProvider
}

// GetState 获取当前状态
func GetState() SmsServiceInterface {
	return stateManagerInstance.getState()
}

func main() {

	// 初始化状态管理
	stateManagerInstance = &StateManager{}
	// 300 毫秒的状态更新   这个方法一直在拿ticket  不会停
	stateManagerInstance.initState(300 * time.Millisecond) //然后一直在set  状态 随机状态   单例statemanager 的当前状态就一直在变

	// 模拟发送短信的接口   拿到的当前状态就 一直不一样   他们的send方法也就不一样   但是消息是一样的
	sendSms := func() {
		// 发送短信  先拿到当前statemanager 单例  然后发消息
		GetState().Send(&Context{
			Tel:        "+8613666666666",
			Text:       "3232",
			TemplateID: "TYSHK_01",
		})
	}

	// 模拟用户调用发送短信的接口   这样不同的状态下 就会发送不一样的 短消息
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
	time.Sleep(1 * time.Second)
	sendSms()
}
