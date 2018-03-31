package factory_method

type user struct{}

type UserDB interface {
	Insert(u user) error
	Select(u user) error
}

type MysqlUserDB struct {
	UserDB
}

func (mysql MysqlUserDB) Insert(u user) error {
	return nil
}
func (mysql MysqlUserDB) Select(u user) error {
	return nil
}

type SqlServerUserDB struct {
	UserDB
}

func (sqlServer SqlServerUserDB) Insert(u user) error {
	return nil
}
func (sqlServer SqlServerUserDB) Select(u user) error {
	return nil
}

type DBFactory interface {
	CreateFactory() UserDB
}

type MysqlUserFactory struct{}

func (mf MysqlUserFactory) CreateFactory() UserDB {
	return new(MysqlUserDB)
}

type SqlServerFactory struct{}

func (sf SqlServerFactory) CreateFactory() UserDB {
	return new(SqlServerUserDB)
}
