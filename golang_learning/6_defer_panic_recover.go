package main

import "fmt"

func solve1() {
	fmt.Println("bbb")
	defer fmt.Println("defer")
	fmt.Println("aaa")
}

func solve2() int{   // 匿名返回值
	var a int
	defer func() {
		a++
	}()
	return a  // 1. ret = a  2. a++  3. return   此时ret=0 a=1
}

func solve3() (a int) {  // 命名返回值
	defer func() {
		a++
	}()
	return a  // 命名返回值时solve2中的ret就是这个a了  所以 1. ret(a) = a  2. a++ 其实就是ret(a)++ 3. return  此时ret=1 a=1 且ret就是这个a了
}

func solve4() {
	defer func(x int) {  // defer 注册顺序是先 上面这个再下面这个  执行顺序相反 先下面这个再上面这个
		fmt.Println(x)
	}(3)
	defer func(x int) {
		fmt.Println(x)
	}(4)
	fmt.Println("solve4")
}

func calc(s string, x int, y int) int {
	z := x + y
	fmt.Println(s, x, y, z)
	return z
}

func solve5() {  // (*)
	x := 1
	y := 2
	fmt.Println("begin defer1")
	defer calc("AA", x, calc("A", x, y))  // 注册这条defer语句时必须确定参数的值 因此需要先执行calc("A", 1, 2) 得到返回值 作为注册函数的第三个参数值
	// 并且此时注册了函数calc("AA", 1, 3) 参数已经确定
	fmt.Println("after defer1 begin defer2")
	x = 10
	defer calc("BB", x, calc("B", x, y))  // 注册时必须先执行calc("B", 10, 2) 同上
	// 并且此时注册了函数calc("BB", 10, 12) 参数已经确定
	y = 20
	fmt.Println("after defer2")
}

func fn1() {
	panic("抛出一个异常")
}

func fn2() {
	defer func() {
		err := recover()
		if err != nil {  // err == nil 表示recover没有接受到异常 即没有异常  err != nil表示有异常
			fmt.Println(err)
		}
	}()
	panic("抛出一个异常")
}

func main() {

	// defer 延迟处理
	solve1()  // bbb aaa defer
	var a = solve2()
	fmt.Println(a)  // 0
	a = solve2()
	fmt.Println(a)  // 0
	a = solve3()
	fmt.Println(a)  // 1
	a = solve3()
	fmt.Println(a)  // 1
	// 导致上述两种情况不一样的原因是 return 语句分为: 1. 返回值ret=a 2. 返回  有了defer语句后： 1. 返回值ret=a  2. 运行defer 3. 返回

	// 注册顺序和执行顺序
	solve4()
	// defer在注册要延迟执行的函数时该函数的所有参数的值必须确定其值
	solve5()
	/*
	begin defer1
	A 1 2 3
	after defer1 begin defer2
	B 10 2 12
	after defer2
	BB 10 12 22
	AA 1 3 4
	 */

	// panic revocer 处理异常  panic可以抛出异常 recover可以获取异常 且recover只有在defer函数中才有效
	//fn1()  // 抛出异常 程序终止执行 后续的语句就不会执行下去
	fn2()  // recover后程序能继续执行
	fmt.Println("程序继续执行")

}