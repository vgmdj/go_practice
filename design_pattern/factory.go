package main

import "fmt"

//Operation 操作接口
type OperationII interface {
	SetNumber(float64, float64)
	GetResult() float64
}

//BaseOperation 基础类
type BaseOperationII struct {
	OperationII
	NumberA float64
	NumberB float64
}

func (bo *BaseOperationII) SetNumberII(numberA, numberB float64) {
	bo.NumberA = numberA
	bo.NumberB = numberB
}

//OperationAdd 加法运算
type OperationAddII struct {
	BaseOperationII
}

func (oa OperationAddII) GetResult() float64 {
	return oa.NumberA + oa.NumberB
}

//OperationSub 减法运算类
type OperationSubII struct {
	BaseOperationII
}

func (os OperationSubII) GetResult() float64 {
	return os.NumberA - os.NumberB
}

//OperationSubII 工厂接口
type IFactory interface {
	CreateOperation() OperationII
}

type FactoryAdd struct{}

func (fa FactoryAdd) CreateOperation() OperationII {
	return new(OperationAddII)
}

//OperationUsage 工厂操作
func OperationUsageII() {
	factory := new(FactoryAdd)
	operation := factory.CreateOperation()
	operation.SetNumber(1, 2)
	fmt.Printf("this is add operation, 1+2=%v\n", operation.GetResult())

}

func main() {
	OperationUsageII()
}
