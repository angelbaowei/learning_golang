package main

import "fmt"

// error l := 2 冒号定义只能用于局部变量 不能用于全局变量
var l = 20201120

func getName(name string) (string, int) {
	return name, 25
}

func main() {

	fmt.Println(l)

	// 定义
	a := 2
	var b = 2
	var c string = "abc"
	var d = "abcd"
	fmt.Println("hello world")
	fmt.Printf("%v %v %v %v\n", a, b, c, d)
	fmt.Printf("%+v %#v %T %+v\n", a, b, c, d)

	p, q := 1, 2
	fmt.Println(p, q)

	// 声明
	var name string  // name 为空
	fmt.Println(name)
	name = "name"
	fmt.Println(name)

	var name2, name3 string
	name2 = "123"
	name3 = "456"
	fmt.Println(name2, name3)

	var (
		username string
		age int
	)
	username = "miaozb"
	age = 25
	fmt.Println(username, age)

	var myname, myage = getName("miaozhibin")
	fmt.Println(myname, myage)
	var _, myage2 = getName("miaozhibin")  // "_" 表示匿名变量 不占用空间 不会分配内存
	fmt.Println(myage2)

	// 常量
	const pi = 3.1415926
	fmt.Printf("%.2f\n", pi)
	const (
		pi1 = 2.5
		pi2       // 不赋值 则默认是=pi1
	)
	fmt.Printf("%v %v\n", pi1, pi2)

	const ca1 = iota  // 0
	const ca2 = iota  // 0
	fmt.Println(ca1, ca2)
	const (
		ca = iota	//0	表示const()内的下标为8
		cb = 2		//2
		cc = iota	//2	下标为2
		cd			//3	跟随cc cd = iota
		ce			//4 跟随cc ce = iota
		cf = 1		//1
		cg			//1
		ch			//1
		ci = iota	//8	下标为8
	)
	fmt.Println(ca, cb, cc, cd, ce, cf, cg, ch, ci)
	const (
		n0 = 1		//1
		n1 = iota	//1
		_			//空 表示这个iota下标用匿名对象 实际意义就是这个iota下标跳过
		n3 = iota	//3
		n4 = 0		//4
		_			//空
		n6, n7, n8, n9 = iota, iota, iota+3, iota+1	//6 6 9 7  这一行的iota=6
	)
	fmt.Println(n0, n1, n3, n4, n6, n7, n8, n9)

	// go变量命名规则
	/*
		常量——全大写字母
		变量名——见名思意 驼峰命名规则 如 age maxAge MaxAge
		特有名词就按原样命名 如 DNS
		变量单独封装成包时 maxAge(首字母小写) 表示私有 MaxAge(首字母大写)表示公有
	*/
}
