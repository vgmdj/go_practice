package simple_factory

/*
设计一个计算器
*/

//Operation 操作接口
type Operation interface {
	SetNumber(float64, float64)
	GetResult() float64
}

//BaseOperation 基础类
type BaseOperation struct {
	Operation
	NumberA float64
	NumberB float64
}

func (bo *BaseOperation) SetNumber(numberA, numberB float64) {
	bo.NumberA = numberA
	bo.NumberB = numberB
}

//OperationAdd 加法运算
type OperationAdd struct {
	BaseOperation
}

func (oa OperationAdd) GetResult() float64 {
	return oa.NumberA + oa.NumberB
}

//OperationSub 减法运算类
type OperationSub struct {
	BaseOperation
}

func (os OperationSub) GetResult() float64 {
	return os.NumberA - os.NumberB
}

//OperationFactory 工厂类
type OperationFactory struct{}

func (of OperationFactory) CreateOperation(oper string) Operation {
	switch oper {
	default:
		return nil
	case "+":
		return new(OperationAdd)
	case "-":
		return new(OperationSub)

	}
}
