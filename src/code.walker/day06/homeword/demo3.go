package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Students struct {
	ID int
	Name string
	Age int
	Score int
}

type Class struct {
	Map map[int]*Students
}

// 展示学生信息
func (c *Class)showStudents()  {
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

// 添加新的学生信息
func (c *Class)addStudents()  {
	var id int
	var name string
	var age int
	var score int
	fmt.Print("输入id: ")
	_, err := fmt.Scan(&id)
	fmt.Print("输入姓名: ")
	_, err = fmt.Scan(&name)
	fmt.Print("输入年龄: ")
	_, err = fmt.Scan(&age)
	fmt.Print("输入分数: ")
	_, err = fmt.Scan(&score)
	if err != nil {
		fmt.Println("保存出错！")
	}
	_, isSave := c.Map[id]
	if isSave {
		fmt.Println("学生ID已存在！")
		return
	}
	student := &Students{
		ID:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
	c.Map[id] = student
	fmt.Printf("学生信息：%#v\n",c)
	fmt.Println("保存成功！")
}
func main() {
	c := &Class{}
	c.Map = make(map[int]*Students, 50)

	c.addStudents()
}


// 生成整数的随机数
func random(min, max int) int {
	return rand.Intn(max - min) + min
}
