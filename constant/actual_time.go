package constant

//拉取第三方大厅排队办事实时数据的三个方法--对应三张表

const (
	HallWindowInfo = iota + 1
	HallTakeNumber
	HallCallNumber
	HallTransactionCompleted
	HallEvaluate
)
