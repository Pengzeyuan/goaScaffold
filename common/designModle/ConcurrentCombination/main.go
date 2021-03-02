package main

/*
// Component 组件接口
type Component interface {
	// 添加一个子组件
	Mount(c Component, components ...Component) error
	// 移除一个子组件
	Remove(c Component) error
	// 执行当前组件业务和执行子组件
	// ctx 业务上下文
	// currentConponent 当前组件
	Do(ctx *Context, currentConponent Component) error
	// 执行当前组件业务业务逻辑
	BusinessLogicDo(ctx *Context) error
	// 执行子组件
	ChildsDo(ctx *Context) error
}

再来看看「并发组合模式」的Component`组件接口，如下(重点看和「组合模式」的区别)：

// Component 组件接口
type Component interface {
	// 添加一个子组件
	Mount(c Component, components ...Component) error
	// 移除一个子组件
	Remove(c Component) error
	// 执行当前组件业务:`BusinessLogicDo`和执行子组件:`ChildsDo`
	// ctx 业务上下文
	// currentConponent 当前组件
	// wg 父组件的WaitGroup对象
	// 区别1：增加了WaitGroup对象参数，目的是等待并发子组件的执行完成。
	Do(ctx *Context, currentConponent Component, wg *sync.WaitGroup) error
	// 执行当前组件业务逻辑
	// resChan 回写当前组件业务执行结果的channel
	// 区别2：增加了一个channel参数，目的是并发组件执行逻辑时引入了超时机制，需要一个channel接受组件的执行结果
	BusinessLogicDo(resChan chan interface{}) error
	// 执行子组件
	ChildsDo(ctx *Context) error
}

我们详细再来看，相对于「组合模式」，引入并发之后需要着重关注如下几点：

并发子组件需要设置超时时间：防止子组件执行时间过长，解决方案关键字context.WithTimeout

区分普通组件和并发组件：合成复用基础组件，封装为并发基础组件

拥有并发子组件的父组件需要等待并发子组件执行完毕(包含超时)，解决方案关键字sync.WaitGroup

并发子组件执行自身业务逻辑是需检测超时：防止子组件内部执行业务逻辑时间过长，解决方案关键字select和<-ctx.Done()

第一点：并发子组件需要设置超时时间
// Context 业务上下文
type Context struct {
	// context.WithTimeout派生的子上下文
	TimeoutCtx context.Context
	// 超时函数
	context.CancelFunc
}

第二点：区分普通组件和并发组件
增加新的并发基础组件结构体BaseConcurrencyComponent，并合成复用「组合模式」中的基础组件BaseComponent，如下：

// BaseConcurrencyComponent 并发基础组件
type BaseConcurrencyComponent struct {
	// 合成复用基础组件
	BaseComponent
	// 当前组件是否有并发子组件
	HasChildConcurrencyComponents bool
	// 并发子组件列表
	ChildConcurrencyComponents []Component
	// wg 对象
	*sync.WaitGroup
	// 当前组件业务执行结果channel
	logicResChan chan interface{}
	// 当前组件执行过程中的错误信息
	Err error
}

第三点：拥有并发子组件的父组件需要等待并发子组件执行完毕(包含超时)
修改「组合模式」中的ChildsDo方法，使其支持并发执行子组件，主要修改和实现如下：

通过go关键字执行子组件

通过*WaitGroup.Wait()等待子组件执行结果

// ChildsDo 执行子组件
func (bc *BaseConcurrencyComponent) ChildsDo(ctx *Context) (err error) {
	if bc.WaitGroup == nil {
		bc.WaitGroup = &sync.WaitGroup{}
	}
	// 执行并发子组件
	for _, childComponent := range bc.ChildConcurrencyComponents {
		bc.WaitGroup.Add(1)
		go childComponent.Do(ctx, childComponent, bc.WaitGroup)
	}
	// 执行子组件
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx, childComponent, nil); err != nil {
			return err
		}
	}
	if bc.HasChildConcurrencyComponents {
		// 等待并发组件执行结果
		bc.WaitGroup.Wait()
	}
	return
}

第四点：并发子组件执行自身业务逻辑是需检测超时
select关键字context.WithTimeout()派生的子上下文Done()方案返回的channel，发生超时该channel会被关闭。具体实现代码如下：

// Do 执行子组件
// ctx 业务上下文
// currentConponent 当前组件
// wg 父组件的waitgroup对象
func (bc *BaseConcurrencyComponent) Do(ctx *Context, currentConponent Component, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	// 初始化并发子组件channel
	if bc.logicResChan == nil {
		bc.logicResChan = make(chan interface{}, 1)
	}

	go currentConponent.BusinessLogicDo(bc.logicResChan)

	select {
	// 等待业务执行结果
	case <-bc.logicResChan:
		// 业务执行结果
		fmt.Println(runFuncName(), "bc.BusinessLogicDo wait.done...")
		break
	// 超时等待
	case <-ctx.TimeoutCtx.Done():
		// 超时退出
		fmt.Println(runFuncName(), "bc.BusinessLogicDo timeout...")
		bc.Err = ErrConcurrencyComponentTimeout
		break
	}
	// 执行子组件
	err = currentConponent.ChildsDo(ctx)
	return
}


*/
