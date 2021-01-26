package main

import (
	"code/day07/01package/calc"
	"code/day07/01package/snow"
	"fmt"
)
func main()  {

	ret := calc.Add(2,3)
	name := calc.Name
	snow.Snow()
	fmt.Print(name)
	fmt.Print(ret)
}
