package factory

import (
	"testing"
)

//OperationUsage 工厂操作
func TestFactory(t *testing.T) {
	factory := new(FactoryAdd)
	operation := factory.CreateOperation()
	operation.SetNumber(1, 2)
	t.Logf("this is add operation, 1+2=%v\n", operation.GetResult())

}
