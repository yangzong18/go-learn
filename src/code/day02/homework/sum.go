package main

import "fmt"

// 求所有数组的和
func main() {

	//1 求数组[1, 3, 5, 7, 8]所有元素的和
	a := [5]int{1,3,5,7,8}
	var sum = 0
	for _,val := range a{
		sum += val
	}

	fmt.Println(sum)

	//2 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	sum2 := 8
	for key2,val2 := range a{
		for key3 := key2+1;key3 < len(a);key3++ {
			if val2+a[key3] == sum2 {
				fmt.Println(key2,key3)
			}

		}

	}
}
