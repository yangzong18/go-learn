package main

import "fmt"
// 打印9*9乘法表
func main() {
	for i:=1;i<=9;i++ {
		for j:=1;j<=i;j++ {
			fmt.Printf("%dX%d=%d\t",j,i,i*j)
		}
		fmt.Println()
	}


	for i:=1;i<=9;i++ {
		for j:=i;j<=9;j++ {
			fmt.Printf("%dX%d=%d\t",i,j,i*j)
		}
		fmt.Println()
	}
}
