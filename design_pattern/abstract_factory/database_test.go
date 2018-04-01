package abstract_factory

import (
	"testing"
)

func TestDatabaseFactory(t *testing.T) {
	user := User{}
	mysqlUser := DatabaseFactory{"mysql"}.CreateUser()
	mysqlUser.Insert(user)

	deppartment := Department{}
	accessDepartment := DatabaseFactory{"access"}.CreateDepartment()
	accessDepartment.Delete(deppartment)

}
