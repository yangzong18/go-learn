package main
type student struct {
	id int //学号是唯一的
	name string
	class string
}

type studentMgr struct {
	allStudent []*student
}

// student 类型的构造函数
func newStudent(id int,name string,class string) *student{
	return &student{
		id:id,
		name:name,
		class: class,
	}
}
func main() {

}