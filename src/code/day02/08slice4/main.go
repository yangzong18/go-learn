package main

import "fmt"

//切片删除元素
func main() {

	a := []int{1,2,3,4,5,6}
	a = append(a[:2],a[3:]...)
	fmt.Println(a)

	b := [...]int{1,3,5,7,9,11,13}
	c := b[:]
	fmt.Printf("c:%v,len:%v,cap:%v\n",c,len(c),cap(c))
	c[0] = 100
	fmt.Println(c)
	fmt.Println(b)

	d :=b[2:3]
	fmt.Printf("d:%v,len:%v,cap:%v\n",d,len(d),cap(d))

	e:=b[3:]
	fmt.Printf("e:%v,len:%v,cap:%v\n",e,len(e),cap(e))
	//如果切片的容量小于 1024，则扩容时其容量大小乘以2；一旦容量大小超过 1024，则增长因子变成 1.25，即每次增加原来容量的四分之一。
	//如果扩容之后，还没有触及原数组的容量，则切片中的指针指向的还是原数组，如果扩容后超过了原数组的容量，则开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
	f := append(e,1)
	fmt.Printf("f:%v,len:%v,cap:%v\n",f,len(f),cap(f))

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // [0 1 2 3 4 5 6 7 8 9] len=10,cap=10
	s1 := s[0:5] // [0 1 2 3 4] len=5,cap=10
	s2 := s[5:] // [5 6 7 8 9] len=5,cap=5
	fmt.Printf("s1:%v,len:%v,cap:%v\n",s1,len(s1),cap(s1))
	fmt.Printf("s2:%v,len:%v,cap:%v\n",s2,len(s2),cap(s2))
}
