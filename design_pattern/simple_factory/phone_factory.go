package simple_factory

import "fmt"

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
