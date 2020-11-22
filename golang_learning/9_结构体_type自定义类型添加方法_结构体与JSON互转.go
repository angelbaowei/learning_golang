package main

import (
	"encoding/json"
	"fmt"
)

// golang没有"类"的概念 只有结构体struct
// 结构体名字Node 首字母大写表示这个结构体是 公有 的  如果定义成node 首字母小写表示这个结构体的 私有 的
// 对于结构体里面的属性 value next 同样适用 首字母大小的规则
type Node struct {
	value int
	next *int
}

// 给结构体结构体方法
func (n Node) printInfo() {  // 这里的n是Node的this对象
	fmt.Println(n)
}

func (n *Node) setInfo(x int) {  // 这里的n是指针类型 即引用类型 改变n就是改变调用对象
	n.value = x
}

// 同样可以给自定义类型 添加方法
type myInt int64
func (p myInt) prints() {
	fmt.Println(p)
}

func main() {

	var n1 Node  // 实例化结构体方式1
	n1.value = 1
	n1.next = new(int)
	fmt.Printf("%v %#v %T\n", n1, n1, n1)

	var n2 = new(Node)  // // 实例化结构体方式2
	// 虽然n2是指针类型 一般来说应当用 *n2.value 访问 但是golang中允许对于结构体对象是指针类型时用直接用"."而不用"*."访问
	n2.value = 2
	n2.next = new(int)
	fmt.Printf("%v %#v %T\n", n2, n2, n2)
	(*n2).value = 3
	fmt.Printf("%v %#v %T\n", n2, n2, n2)

	var n3 = &Node{} // 实例化结构体方式3
	n3.value = 4
	n3.next = nil
	fmt.Printf("%v %#v %T\n", n3, n3, n3)

	var n4 = Node{  // 实例化结构体方式4
		value: 5,
		next: nil,
	}
	fmt.Printf("%v %#v %T\n", n4, n4, n4)

	// 结构体是值类型

	// 结构体方法
	n1.printInfo()
	n1.setInfo(10)
	n1.printInfo()

	// 自定义类型添加方法
	var p myInt
	p = 3
	p.prints()

	// 结构体嵌套
	type Info struct {
		Age int
	}
	type Person struct {
		Name string
		If Info
	}
	var P Person
	P.Name = "a"
	P.If.Age = 12
	fmt.Println(P)

	// 结构体的继承是通过结构体的嵌套来实现的
	type Person2 struct {
		Name string
		Info  // Person2 继承了 Info
	}
	var P2 Person2
	P2.Name = "b"
	P2.Age = 13
	fmt.Println(P2)

	// JSON 与 结构体 相互转换  序列化与反序列化

	// 结构体转JSON
	jsonByte1, _ := json.Marshal(P2)
	jsonStr1 := string(jsonByte1)
	fmt.Println(jsonStr1)

	// JSON转结构体
	var str = `{"Name":"c","Age":14}`  // 反引号定义
	var P3 Person2
	err := json.Unmarshal([]byte(str), &P3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(P3)
	}

	// 结构体转JSON时 结构体的字段必须是公有(大写)的 否则转换不了

	// 结构体标签Tag
	type Stu struct {
		Name string `json:"NAME"`
		Id int `json:"ID"`
	}
	var stu Stu
	stu.Name = "miaozb"
	stu.Id = 1234
	jsonByte1, _ = json.Marshal(stu)
	jsonStr1 = string(jsonByte1)
	fmt.Println(jsonStr1)  // {"NAME":"miaozb","ID":1234}

	

}
