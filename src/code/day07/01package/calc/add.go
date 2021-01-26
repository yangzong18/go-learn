package calc

import (
	"fmt"
)

// 定义一个package 表示符大写，表示外部可以访问 一般都小写

// Name 全局变量
var Name = "沙河小王子"

// Add 定义一个方法
func Add(x,y int) int {
	return x+y
}
// 包倒入的时候，自动执行init
// 没有参数，没有返回值
func init(){
	fmt.Println("calc.init()")
}

func main()  {
	Sub(200,400)
}