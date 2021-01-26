package main

import "fmt"

type student struct {
	id int //学号是唯一的
	name string
	class string
}

// student 类型的构造函数
func newStudent(id int,name string,class string) *student{
	return &student{
		id:id,
		name:name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

// newStudentMagr 是studentMgr 的构造函数
func newStudentMagr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student,0,100),
	}
}

// 1.添加学生
func (s *studentMgr)addStudent(newStu *student)  {
	exist := true
	for _,stu := range s.allStudents{
		if newStu.id == stu.id {
			fmt.Printf("学号已经存在，请重新输入\n")
			exist = false
			break
		}
	}
	if exist{
		s.allStudents = append(s.allStudents,newStu)
		fmt.Print("添加成功\n")
	}
}
// 2.编辑学生
func (s *studentMgr)editStudent(newStu *student)  {
	for key,student := range s.allStudents{
		if newStu.id == student.id { // 学号相同
			s.allStudents[key] = newStu //根据切片的索引直接吧新学生赋值
			fmt.Print("编辑成功\n")
			return
		}
	}
	fmt.Printf("输入的学生信息有误，系统中没有学号是：%d的学生\n",newStu.id)
}
// 3.展示学生
func (s *studentMgr)showStudent()  {
	for _,student := range s.allStudents{
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n",student.id,student.name,student.class)
	}
}

// 4.删除学生
func (s *studentMgr)delStudent(id int)  {
	newStudents := make([]*student,0,200)
	for k,student := range s.allStudents{
		if id == student.id {
			fmt.Printf("del stu id: %#v\n", id)
			continue
		}
		newStudents = append(newStudents, s.allStudents[k])
		s.allStudents = newStudents
	}
	fmt.Println("删除成功！！")
}