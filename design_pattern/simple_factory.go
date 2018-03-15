package main

import "fmt"

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

//OperationUsage 工厂操作
func OperationUsage() {
	factory := new(OperationFactory)
	operation := factory.CreateOperation("+")
	operation.SetNumber(1, 2)
	fmt.Printf("this is add operation, 1+2=%v\n", operation.GetResult())

	operation = factory.CreateOperation("-")
	operation.SetNumber(2, 1)
	fmt.Printf("this is sub operation, 2-1=%v\n", operation.GetResult())

}

/*
设计一个工厂来生产各种厂商的手机
其中初始的厂商有小米，苹果，华为

*/

//Phone interface
type Phone interface {
	ShowBrand()
}

//IPhone apple
type IPhone struct {
}

func (phone IPhone) ShowBrand() {
	fmt.Println("[Phone Brand]: Apple")
}

//HPhone huawei
type HPhone struct {
}

func (phone HPhone) ShowBrand() {
	fmt.Println("[Phone Brand]: Huawei")
}

//XPhone xiaomi
type XPhone struct {
}

func (phone XPhone) ShowBrand() {
	fmt.Println("[Phone Brand]: Xiaomi")
}

type PhoneFactory struct{}

func (factory PhoneFactory) CreatePhone(brand string) Phone {
	switch brand {
	default:
		return nil
	case "HW":
		return new(HPhone)
	case "XM":
		return new(XPhone)
	case "PG":
		return new(IPhone)

	}
}

func PhoneUsage() {
	factory := PhoneFactory{}
	phone := factory.CreatePhone("HW")
	phone.ShowBrand()

}

func main() {
	OperationUsage()
	PhoneUsage()
}
