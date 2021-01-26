package main

import "fmt"

func main() {
	var a = make([]string, 5, 10) // [     ]
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i)) // 容量不够会自动申请扩容
	}
	fmt.Println(a) // [     0 1 2 3 4 5 6 7 8 9]
}
