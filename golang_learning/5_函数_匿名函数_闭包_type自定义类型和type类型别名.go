package main

import "fmt"

// 可变参数
func solve1(x ...int) int{  // x是 多个参数 被转换成的一个切片数组
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// return 多个值
func solve2(x int, y ...int) (int, int) {
	sum := 0
	for _, v := range y {
		sum += v
	}
	return x, sum
}

// 返回值命名
func solve3(x int) (y int) {
	y = x + 1
	return
}

// 切片是引用类型
func solve4(x []int) []int {
	y := x
	return y
}

// type 自定义类型
type calc func(int) int
type myint int64
// type 类型别名  类型别名有"="号
type myint2 = int64

type calctype func(int, int) int
func add(x int, y int) int {
	return x + y
}
func solve5(x int, y int, F calctype) int {
	return F(x, y)
}

// 函数作为返回值
func solve6(op string) calctype {
	switch op {
	case "+":
		return add
	case "-":
		return func(x int, y int) int {  // 匿名函数
			return x - y
		}
	default:
		fmt.Println("nil")
		return nil
	}
}

// 闭包: 可理解为定义在一个函数内部的函数 本质上 闭包是将函数内部和函数外部连接起来的桥梁  闭包是有权访问另一个函数作用域中的变量的函数
/*
	全局变量： 1. 常驻内存 	2. 污染全局(全局变量定义一个a int 局部变量仍可以定义一个a int)
	局部变量： 1. 不常驻内存 	2. 不污染全局
	闭包：	 可以让一个变量常驻内存且不污染全局
*/
// 闭包的写法  函数里面嵌套一个函数 最后返回嵌套的里面的那个函数
// 这里面的 i 是个局部变量 不会污染全局  但这个 i 会常驻内存 (*)
func solve7() func() int {
	var i = 1
	return func() int {
		return i + 1
	}
}

func solve8() func() int {
	var i = 1
	return func() int {
		i++
		return i + 1
	}
}

func main() {
	/* 	func 函数名(参数) 返回值 {
		函数体
	}
	func solve(a int) string {
		return "a"
	}
	*/


	var sum = solve1(1, 2, 3, 4)  // 不固定参数个数
	fmt.Println(sum)
	x, sum := solve2(10, 1, 2, 3, 4, 5)
	fmt.Println(x, sum)
	x = solve3(11)
	fmt.Println(x)

	var slice1 = []int{1, 2, 3}
	slice2 := solve4(slice1)
	fmt.Println(slice1, slice2)
	slice2[1] = 22
	fmt.Println(slice1, slice2)

	var f calc  // calc是一个函数类型的别名  自定义了函数类型
	f = solve3
	fmt.Printf("%v %#v %T\n", f, f, f) // 0xe277c0 (main.calc)(0xe277c0) main.calc
	g := solve3
	fmt.Printf("%v %#v %T\n", g, g, g)  // 0xe277c0 (func(int) int)(0xe277c0) func(int) int
	var c = f(9)  // 调用方法
	fmt.Println(c)

	var h myint
	h = 5
	fmt.Printf("%v %#v %T\n", h, h, h)  // 5 5 main.myint
	var m myint2
	m = 5
	fmt.Printf("%v %#v %T\n", m, m, m)  // 5 5 int64  注意type自定义类型和type类型别名的区别

	var i int64 = 6
	//j := h + i  error 虽然h的myint就是int64 但是仍然不能与int64类型的i直接相加
	j := myint(i) + h
	fmt.Printf("%v %#v %T\n", j, j, j)  // 11 11 main.myint

	// 类似函数指针的用法
	v := solve5(3, 4, add)  // add是一个函数方法  这个solve5函数就实现了一个类似C语言中函数指针的作用
	fmt.Println(v)

	// 函数不能直接定义在另一个函数体内 因此需要匿名函数
	// 匿名函数
	v = solve5(3, 4, func(x int, y int) int {
		return x * y
	})
	fmt.Println(v)

	// 匿名自执行函数
	func(x int) {
		fmt.Println(x, "匿名自执行函数")
	}(10) // ()表示执行

	// 函数作为返回值
	op := solve6("+")
	v = op(3, 4)
	fmt.Println(v)
	op = solve6("-")
	v = op(3, 4)
	fmt.Println(v)

	// 闭包
	op2 := solve7()  // 执行solve7方法 返回了一个嵌套的那个方法
	fmt.Println(op2())  // 2
	fmt.Println(op2())  // 2
	fmt.Println(op2())  // 2  i的值不变
	op3 := solve8()  // 执行solve8方法 返回了一个嵌套的那个方法
	fmt.Println(op3())  // 3
	fmt.Println(op3())  // 4
	fmt.Println(op3())  // 5  i的值改变了
}
