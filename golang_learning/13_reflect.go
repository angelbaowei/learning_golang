package main

import (
	"fmt"
	"reflect"
)

type myint int64
type myint2 = int64
type Person struct {
	Name string
	age int
}

// 类型反射
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)  // 获取变量的类型
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n", v, v.Name(), v.Kind())
}

// 值反射
func reflectValue(x interface{}) {
	//var num = x + 10 // error interface + int
	// 1
	b, _ := x.(int)  // assert
	var num = b + 10
	fmt.Println(num)
	// 2
	var r = reflect.ValueOf(x)  // reflect.ValueOf  获取变量的值
	fmt.Println(r.Kind())  // int  反射得到值可能获取king底层类型
	if r.Kind() == reflect.Int {
		fmt.Println("r is int type")
		//r.SetInt(6)  // error 值类型不能用 要用指针(引用)  见 reflectValue_ref
	}
	var num2 = r.Int() + 10
	fmt.Println(num2)
}
// 引用类型反射
func reflectValue_ref(x interface{}) {
	//var num = *x + 10 // error interface + int
	// 1
	b, _ := x.(*int)  // assert
	var num = *b + 10
	fmt.Println(num)
	// 2
	var r = reflect.ValueOf(x)  // reflect.ValueOf  获取变量的值
	fmt.Println(r.Kind(), r.Elem().Kind())  // ptr int   引用类型使用r.Elem().Kind()
	if r.Kind() == reflect.Ptr {
		fmt.Println("r is ptr type")
	}
	if r.Elem().Kind() == reflect.Int {
		fmt.Println("r.Elem() is int type")
		r.Elem().SetInt(8)  // 修改指针指向的元素的值
	}
	var num2 = r.Elem().Int() + 10
	fmt.Println(num2)
}

type Node struct {
	name string `json:"name1" form:"h1"`
	age int
	Score float32
}

func (n Node) Print() {
	fmt.Println(n.name, n.age, n.Score)
}

func (n *Node) SetInfo(name string, age int, score float32) {
	n.name = name
	n.age = age
	n.Score = score
}

// 结构体反射
func PrintStructFiled(s interface{}) {
	t := reflect.TypeOf(s)  // 获取类型
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体类型")
		return
	}
	field0 := t.Field(0)  // Field index
	fmt.Println(field0.Name, field0.Type, field0.Tag.Get("json"), field0.Tag.Get("form"))  // name string name1 h1  ;
	// index=0的字段的类型名是 name 类型是 string  Tag标签 json标签是 name1 form标签是 h1  第58行

	field1, ok := t.FieldByName("age")  // FieldByName
	if ok {
		fmt.Println(field1.Name, field1.Type)  // age int
	}

	var fieldCount = t.NumField()
	fmt.Println("这个结构体的字段个数为：", fieldCount)  // 3

	// 获取值
	v := reflect.ValueOf(s)  // 值
	score := v.FieldByName("Score")
	if score.Kind() == reflect.Float32 {
		fmt.Println(score)
	}
	field2 := v.Field(0)  // 获取index=0的值
	fmt.Println(field2.Type(), field2)  // string abc


}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)  // 类型
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("不是结构体类型")
		return
	}

	// 获取结构体内的方法
	method0 := t.Method(0)  // index 和结构体的顺序没有关系  和结构体方法名的ASCII有关系  所以不推荐这种方式
	fmt.Println(method0.Name, method0.Type)  // Print func(main.Node)

	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name, method1.Type)  // Print func(main.Node)
	}

	// 执行方法
	v := reflect.ValueOf(s)
	v.MethodByName("Print").Call(nil) // abc 11 90.4 ； nil表示参数为空
	var params []reflect.Value
	params = append(params, reflect.ValueOf("123"))
	params = append(params, reflect.ValueOf(12))
	params = append(params, reflect.ValueOf(float32(23.0)))
	v.MethodByName("SetInfo").Call(params)
	v.MethodByName("Print").Call(nil) // 123 12 23 ； nil表示参数为空
	// 方法数量
	fmt.Println(v.NumMethod(), t.NumMethod())  // 2 2

	// 修改结构体字段的值
	//v.Elem().FieldByName("name").SetString("def")  // error panic: reflect: reflect.Value.SetString using value obtained using unexported field 因为name第一个字母小写 是私有的 不能在结构体外部修改
	v.Elem().FieldByName("Score").SetFloat(100.0)
	v.MethodByName("Print").Call(nil) // 123 12 100 ； nil表示参数为空

}


func main() {
	// Go语言提供了一种机制在 运行时 更新和检查变量的值、调用变量的方法和变量支持的内在操作，但是在 编译时并不知道这些变量的具体类型 ，这种机制被称为反射
	// 反射的代码较脆弱 bug不容易发现  代码本身可读性也较差
	a := 10
	b := 10.0
	c := "10.0"
	reflectType(a)	//int int int
	reflectType(b)	//float64 float64 float64
	reflectType(c)	//string string string
	person := Person{
		"miaozb",
		24,
	}
	reflectType(person)	//main.Person Person struct
	var d myint = 4
	reflectType(d) // main.myint myint int64
	var e myint2 = 4
	reflectType(e)  // int64 int64 int64
	var p = new(int)
	reflectType(p)  // *int 空 ptr

	var i = [3]int{1, 2, 3}
	var j = []int{1, 2, 3}
	reflectType(i)  // 类型：[3]int 类型名称： 类型种类：array
	reflectType(j)  // 类型：[]int 类型名称： 类型种类：slice

	reflectValue(12)  // 值类型
	var f = 2
	reflectValue_ref(&f)  // 引用类型
	fmt.Println(f)  // 8


	// 结构体反射
	node := Node{
		"abc",
		11,
		90.4,
	}
	PrintStructFiled(node)
	PrintStructFn(&node)  // 传入引用类型  传入值类型会129行会报错panic: reflect: call of reflect.Value.Call on zero Value

}