package main

import "fmt"

func main() {
	var a = 100
	b := "红藕"
	c := false
	fmt.Println(a,b,c)
	fmt.Printf("%v\n",a)
	fmt.Printf("%t\n",a)
	fmt.Printf("%d%%\n",a)

	fmt.Printf("%5d\n",a)
	fmt.Printf("%10d\n",a)

	d := 3.14159264
	// 小数点后几位
	fmt.Printf("%.2f\n",d)
	// 一共几位数字
	fmt.Printf("%.2g\n",d) // 3.1
	fmt.Printf("%.8g\n",d) // 3.1415926

}
