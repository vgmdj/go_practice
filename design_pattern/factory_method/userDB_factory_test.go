package factory_method

import "testing"

func TestDBFactory(t *testing.T) {
	factory := new(MysqlUserFactory)
	mysqlUser := factory.CreateFactory()
	mysqlUser.Insert(user{})
	mysqlUser.Select(user{})
}
