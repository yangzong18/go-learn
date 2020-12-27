package main

import "fmt"

func main() {

	a := []int{}

	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
	a = append(a,1)
	fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n",a,len(a),cap(a),a)
}
