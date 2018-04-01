package Strategy

import "fmt"

/*
商场折扣策略
- 固定折扣
- 满减

输入商品单价，输出找完折后应付的价格
*/

type CashSuper interface {
	Accept(float64) float64
}

//CashNormal normal strategy
type CashNormal struct{}

func (normal CashNormal) Accept(money float64) float64 {
	return money
}

//CashRebate
type CashRebate struct {
	moneyRebate float64
}

func (re CashRebate) Accept(money float64) float64 {
	return money * re.moneyRebate
}

//CashReturn
type CashReturn struct {
	moneyCondition float64
	moneyReturn    float64
}

func (re CashReturn) Accept(money float64) float64 {
	if money >= re.moneyCondition {
		return money - re.moneyReturn
	}
	return money
}

//use simple factory optimize strategy
type CashContext struct {
	strategy CashSuper
}

func NewCashContext(acceptType string) (cashFactory CashContext) {
	switch acceptType {
	default:
		fmt.Println("wrong type")

	case "normal":
		cashFactory.strategy = CashNormal{}

	case "0.8rebate":
		cashFactory.strategy = CashRebate{moneyRebate: 0.8}

	case "300return100":
		cashFactory.strategy = CashReturn{300, 100}

	}
	return
}

func (cashFactory CashContext) Accept(money float64) float64 {
	return cashFactory.strategy.Accept(money)
}
