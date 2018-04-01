package abstract_factory

import "fmt"

const (
	Mysql  = "mysql"
	Access = "access"
)

type DatabaseFactory struct {
	DB string
}

func (df DatabaseFactory) CreateUser() IUser {
	switch df.DB {
	default:
		fmt.Println("unknown database return default mysql")
		return new(MysqlUser)

	case Mysql:
		return new(MysqlUser)

	case Access:
		return new(AccessUser)
	}

}

func (df DatabaseFactory) CreateDepartment() IDepartment {
	switch df.DB {
	default:
		fmt.Println("unknown database return default mysql")
		return new(MysqlDepartment)

	case Mysql:
		return new(MysqlDepartment)

	case Access:
		return new(AccessDepartment)
	}

}
