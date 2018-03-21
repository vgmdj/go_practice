package simple_factory

import (
	"testing"
)

func TestFactory(t *testing.T) {

	opFactory := new(OperationFactory)
	operation := opFactory.CreateOperation("+")
	operation.SetNumber(1, 2)
	t.Logf("this is add operation, 1+2=%v\n", operation.GetResult())

	operation = opFactory.CreateOperation("-")
	operation.SetNumber(2, 1)
	t.Logf("this is sub operation, 2-1=%v\n", operation.GetResult())

	phonefactory := PhoneFactory{}
	phone := phonefactory.CreatePhone("HW")
	phone.ShowBrand()
}
