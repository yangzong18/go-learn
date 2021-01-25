package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统 需求
//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息
func main() {
	for  {
		// 1.打印系统菜单
		showMenu()
		// 2.等待用户选择要执行的选项
		var input int
		fmt.Print("请输入你要操作的序号：")
		fmt.Scanf("%d\n",&input)
		fmt.Println("用户输入的是：",input)
		// 3.执行用户选择的动作
		switch input {
		case 1:
			// 添加学员
		case 2:
			// 编辑学员
		case 3:
			// 展示
		case 4:
			//退出
			os.Exit(0)


		}
	}



}

func showMenu()  {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.展示学员")
	fmt.Println("4.退出系统")
}

