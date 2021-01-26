package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)
// 学员信息管理系统
//1.添加学员信息
//2.编辑学员信息
//3/

type Students struct {
	ID int
	Name string
	Age int
	Score int
}

type ClassDemo struct {
	Map map[int]*Students
}

// 展示学生信息
func (c *ClassDemo)showStudents()  {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "ID", "姓名", "年龄", "分数")
	sortID := make([]int, 0)
	for k := range c.Map {
		sortID = append(sortID, k)
	}
	sort.Ints(sortID)
	for _, k := range sortID {
		s := c.Map[k]
		fmt.Printf("\t%d\t%s\t%d\t%d\n", s.ID, s.Name, s.Age, s.Score)
	}
}


// 获取用户的输入信息 返回一个map 指针
func getInputs() *Students {
	var (
		id,age,score int
		name string
	)
	fmt.Print("输入id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("输入信息有误，请重新输入")
	}
	student := &Students{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	return student
}

// 保存用户信息
func (sm *ClassDemo)addStudent(s * Students)  {
	_, isSave := sm.Map[s.ID]
	if isSave {
		fmt.Println("学生ID已存在！")
		return
	}
	sm.Map[s.ID] = s
	fmt.Println("保存成功！")
}
// 编辑用户信息
func (sm *ClassDemo)editStudent(s * Students)  {
	_, isSave := sm.Map[s.ID]
	if !isSave {
		fmt.Println("学生ID不存在！")
		return
	}
	sm.Map[s.ID] = s
	fmt.Println("保存成功！")
}

// 删除用户
func (sm *ClassDemo)delStudent(s * Students)  {
	_, isSave := sm.Map[s.ID]
	if !isSave {
		fmt.Println("学生ID不存在！")
		return
	}
	delete(sm.Map,s.ID)
	fmt.Println("删除成功！")
}
// 展示用户
func (sm *ClassDemo)showList()  {
	fmt.Printf("\t%s\t%s\t%s\t%s\n", "ID", "姓名", "年龄", "分数")
	sortID := make([]int, 0)
	for k := range sm.Map {
		sortID = append(sortID, k)
	}
	sort.Ints(sortID)
	for _,k := range sortID{
		student := sm.Map[k]
		fmt.Printf("\t%d\t%s\t%d\t%d\n",student.ID,student.Name,student.Age,student.Score)
	}
}
func main() {
	sm := &ClassDemo{}
	sm.Map = make(map[int]*Students,200)
	for  {
		showMenu()
		var input int
		fmt.Println("请输入序号：")
		fmt.Scanf("%d\n",&input)
		fmt.Println("用户输入的是：",input)
		switch input {
		case 1:
			// add
			stu := getInputs()
			sm.addStudent(stu)
		case 2:
			// edit
			stu := getInputs()
			sm.editStudent(stu)
		case 3:
			// del
			stu := getInputs()
			sm.delStudent(stu)
		case 4:
			// show
			sm.showList()
		case 5:
			// quit 系统
			os.Exit(0)
		}
	}

}


// 生成整数的随机数
func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func showMenu()  {
	fmt.Println("欢迎来到walker学习系统")
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生")
	fmt.Println("3.删除学生")
	fmt.Println("4.展示学生")
	fmt.Println("5.退出系统")
}
