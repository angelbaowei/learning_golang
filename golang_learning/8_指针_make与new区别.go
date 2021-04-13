package main

import "fmt"

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 10
}

func main() {

	var a = 1
	fmt.Printf("%v %v\n", a, &a)  // 1 0xc000016068
	var p = &a
	fmt.Printf("%v %T, %v %T, %v %T %v\n", a, a, p, p, *p, *p, &p)
	// 1 int, 0xc000016068 *int, 1 int 0xc00000e030

	fn1(a)
	fmt.Println(a)  // 1
	fn2(&a)
	fmt.Println(a)  // 10

	// 引用数据类型必须分配内存空间(make new)才能赋值
	// var q []int
	// q[0] = 1  error
	// var q = make([]int, 4, 4)  right
	//var q map[int]string
	//q[1] = "a" //error
	//var q = make(map[int]string) //right
	// var q *int
	// *q = 1  error
	// var q = new(int) right	  指针也是引用类型

	// make和new区别
	// new用于指针  make用于slice map chan
	// new返回的类型是指针类型 make返回类型是数据本身的引用类型

	var q = new(int)  // 对于指针类型的内存分配 用new  不能用make
	// var q = make(*int)  error
	*q = 1
	fmt.Println(q, *q)



}