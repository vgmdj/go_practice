package abstract_factory

type Department struct {
}

type IDepartment interface {
	Insert(department Department)
	Delete(department Department)
}

type AccessDepartment struct{}

func (ad AccessDepartment) Insert(department Department) {

}

func (ad AccessDepartment) Delete(department Department) {

}

type MysqlDepartment struct{}

func (md MysqlDepartment) Insert(department Department) {

}

func (md MysqlDepartment) Delete(department Department) {

}
