package main

import "fmt"

func main() {

	// if 语句
	if age := 30; age > 20 {  // age是只存在于if内的局部变量
		fmt.Println(age)
	}
	// fmt.Println(age) undefined: age

	i := 1
	for i <= 3 {  // 相当于while
		fmt.Printf("%d\t", i)
		i++
	}
	fmt.Println()
	i = 1
	for {  // 相当于while(true)
		if i <= 3 {
			fmt.Printf("%d\t", i)
		} else {
			break
		}
		i++
	}
	fmt.Println()

	for i = 1; i <= 3; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	str := "abc"
	for key, value := range str {
		fmt.Printf("%v %v\t", key, value)
	}
	fmt.Println()

	switch extname := ".htm"; extname {  // switch可以不加break
	case ".html", ".htm":  // 多个值中的一个匹配就可以
		fmt.Println(extname)
		fallthrough  // 能够无条件穿透到下一个紧挨的case内的语句 穿透一层
	case ".css":
		fmt.Println("css")
		fmt.Println("css2")
	case ".cpp":
		fmt.Println("cpp")
	default:
		fmt.Println("default")
	}

label1:
	for i = 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			if i * j > 4 {
				break label1  // 跳出多层循环
			}
			fmt.Println(i, j)
		}
	}
fmt.Println("----------------------------------")
label2:
	for i = 1; i <= 5; i++ {
		for j := 1; j <= 2; j++ {
			if i * j < 4 {
				continue label2  // i = 1, j = 1  continue label2 i = 2, j = 1 continue label2 i = 3, j = 1
			}
			fmt.Println(i, j)
		}
	}

	var n = 30
	if n > 20 {
		goto label3 // goto
	}
	fmt.Println("aaa")
label3:
	fmt.Println("bbb")

}
