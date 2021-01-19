package main

import "fmt"

// 多维数组
func main() {
	var b [3][2]int
	b = [3][2]int{
		[2]int{1,2},
		[2]int{2,3},

	}

	fmt.Println(b)

	c := [3][2]int{
		{1,2},
		{2,4},
	}

	fmt.Println(c)

	d := [2][2]string{
		{"php","java"},
		{"javascript","html"},
	}


	fmt.Printf("%v",d)

	//注意： 多维数组只有第一层可以使用...来让编译器推导数组长度

	//支持的写法
	a1 := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println()
	fmt.Println(a1)


	// 多位数组的便利
	for _,dd := range a1{
		for _,dd2 := range dd{
			fmt.Printf("%v\n",dd2)
		}
	}

	n := [2]int{1,2}
	m := n
	m[0] = 8
	fmt.Println(n)
	fmt.Println(m)
}
