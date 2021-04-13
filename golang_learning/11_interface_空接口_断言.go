package main

import "fmt"

// 接口 interface 接口中不能包含任何变量  接口定义了行为规范 只有方法声明 没有方法体
// 接口是一个规范
type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话 必须通过 结构体或者自定义类型 实现接口里面的方法
type Phone struct {
	Name string
}
// Phone要实现Usber接口的话 必须实现其接口里的所有方法
// 值类型接受者
func (p Phone) start() {
	fmt.Println(p.Name, "start!")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "stop!")
}

func (p Phone) run() {
	fmt.Println(p.Name, "RUN!")
}

type Com struct {

}

func (c Com) solve(usb Usber) {
	usb.start()
	usb.stop()
}

// 空接口 表示没有任何约束 则任意类型都可以实现空接口
type Kong interface {

}

func show(any interface{}) {
	fmt.Printf("%v %T\n", any, any)
}

type Camera struct {
	Name string
}

// 指针类型接收者
func (c *Camera) start() {
	fmt.Println(c.Name, "start!")
}
// 值类型接收者
func (c Camera) stop() {
	fmt.Println(c.Name, "stop!")
}

// 结构体实现多个接口
type A interface {
	SetName(string)
}
type B interface {
	GetName() string
}
type C struct {
	Name string
}
func (c *C) SetName(name string) {
	c.Name = name
}
func (c C) GetName() string{
	return c.Name
}

// 嵌套接口 必须实现所有接口(A B Q)里的方法
type Q interface {
	A
	B
}

func main() {
	var p Phone
	p.Name = "HUAWEI"
	fmt.Println(p)
	p.start()
	p.stop()

	var p1 Usber = p  // 手机实现usb接口   这个写法注意理解
	p1.start()
	p1.stop()
	// p1.run() error p1中没有run方法
	p.run()

	var c Com
	c.solve(p)  // p对象所属的类必须实现usb接口内的所有方法才能调用成功  此调用参数相当于 var usb Usber = p
	c.solve(p1)

	// 空接口
	var str = "abc"
	var p2 Kong = str  // string实现空接口
	fmt.Println(p2)
	var id = 3
	var p3 Kong = id  // int实现空接口
	fmt.Println(p3)

	var a interface{}  // 空接口表示任意类型
	a = 1
	show(a)  // 1 int
	a = "asdsad"
	show(a)  // asdsad string
	a = true
	show(a)  // true bool

	var m1 = make(map[interface{}]interface{})
	m1[0] = 1
	m1[1] = "abc"
	m1[2] = 1.23
	m1["asd"] = 2
	fmt.Println(m1)

	// 对切片也适用

	// 断言  用于 接口类型的断言
	var b interface{} = "abc"
	v, ok := b.(string) // 断言 b 是 string 类型
	if ok {
		fmt.Println(v)
	}

	// x.(type) 只用于switch语句中 用于判断x的类型
	switch b.(type) {
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}

	// 值类型接收者 和 指针类型接收者
	// 值类型接收者 (第18和22行 的 p Phone p是值类型)
	// 指针类型接收者 (第53和57行 的 c *Camera c是指针类型)
	var pp Phone
	pp.Name = "xiaomi"
	var usb1 Usber = pp
	usb1.start()
	usb1.stop()

	//var cc Camera
	//cc.Name = "suoni"
	// var usb2 Usber = cc  error Camera类实现了指针类型接收者 因此cc必须是指针(引用)类型

	//var cc = &Camera{}  // 引用类型
	var cc = new(Camera)  // 或者这个
	cc.Name = "suoni"
	var usb2 Usber = cc
	usb2.start()  // 指针类型接受者
	usb2.stop()  // 值类型接受者
	// 综上： 只要出现了指针类型接受者 一律要用 引用类型

	// 一个结构体实现多个接口
	var ccc = &C{}
	var bbb B = ccc  // C实现B接口
	var aaa A = ccc  // C实现A接口  A接口的SetName方法是引用类型接受者
	aaa.SetName("C实现A和B两个接口")
	fmt.Println("-----", bbb.GetName())  // 因为ccc是引用类型 所以setName aaa 也就setName了bbb

	// 空接口类型不支持索引

	type node struct {
		Name string
		Id int
	}
	var maze = make(map[int]interface{})
	maze[1] = 1
	maze[3] = "abc"
	maze[5] = []string{"aaa", "bbb"}
	fmt.Println(maze[1])
	fmt.Println(maze[3])
	fmt.Println(maze[5])
	// fmt.Println(maze[5][1]) error  type interface {} does not support indexing

	var newnode = node{
		Name: "miaozb",
		Id: 123456,
	}
	maze[4] = newnode
	fmt.Println(maze[4])
	//fmt.Printf(maze[4].Name) error type interface {} is interface with no methods

	// 通过类型断言解决上述问题
	vv, ok := maze[5].([]string)  // 断言
	if ok{
		fmt.Println(vv[0])
		fmt.Println(vv[1])  // 解决的第182行问题
	}

	vv2, ok := maze[4].(node)  // 断言
	if ok{
		fmt.Println(vv2.Name)
		fmt.Println(vv2.Id)  // 解决的第190行问题
	}

}