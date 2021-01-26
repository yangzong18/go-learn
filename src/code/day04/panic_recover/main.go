package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func a()  {
	fmt.Println("func a")
}
/*
recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。*/
func b()  {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic b")
}

func c()  {
	fmt.Println("func c")
}
