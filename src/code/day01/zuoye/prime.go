package main

import "fmt"
//但因200-10000的素数
func main() {
	for i:=200;i<=10000;i++{
		if(isPrime(i)){
			fmt.Printf("%d\n",i)
		}
	}
}


func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
