package abstract_factory

type User struct{}

type IUser interface {
	Insert(user User)
	Delete(user User)
}

type AccessUser struct{}

func (au AccessUser) Insert(user User) {

}

func (au AccessUser) Delete(user User) {

}

type MysqlUser struct{}

func (mu MysqlUser) Insert(user User) {

}

func (mu MysqlUser) Delete(user User) {

}
