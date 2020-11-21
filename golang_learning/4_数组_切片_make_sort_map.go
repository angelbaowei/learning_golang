package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
	1.值类型：变量直接存储值，内存通常在栈中分配。
	值类型：基本数据类型int、float、bool、string、数组、struct
	2.引用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在 堆上分配。通过GC回收。
	引用类型：指针、切片、map、chan等都是引用类型。
	 */

	/*
	声明没有赋值时的默认值：
		bool->false
		int float->0
		string->""

		slices->nil
		maps->nil
		channels->nil
		functions->nil
		pointers->nil
		interface->nil

	注意使用make创建时默认值不同！

	*/

	// 数组长度声明的时候就固定了 长度不可变
	var arr1 [3]int  // 数组
	arr1[0] = 1
	fmt.Printf("%v %#v %T\n", arr1, arr1, arr1)
	var arr2 = [3]string{"c++", "java", "golang"}
	fmt.Printf("%v %#v %T\n", arr2, arr2, arr2)
	var arr5 = [...]float32{0:1.2, 2, 3:1, 3.5, 5.1, 2}
	fmt.Printf("%v %#v %T\n", arr5, arr5, arr5)

	// 多维数组
	var multiarr = [...][2]int{  // 第一维度可以为动态的 "..."  第二维不可以
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(multiarr)

	// 切片 长度可变
	var arr3 = [...]float32{1.2, 1, 3.5, 5.1, 2} // 数组
	fmt.Printf("%v %#v %T\n", arr3, arr3, arr3)
	var arr4 = []float32{1.2, 1, 3.5, 5.1, 2} // 切片
	fmt.Printf("%v %#v %T\n", arr4, arr4, arr4)

	var a = [...]int{1, 2, 3}
	var b = a
	b[1] = 22  // 改变b 不改变 a 值类型 copy过来了
	fmt.Println(a, b)  // [1 2 3] [1 22 3]

	var sa = []int{1, 2, 3}  // 切片 引用类型
	var sb = sa
	sb[1] = 22  // 改变sb的同时也改变了sa 引用类型
	fmt.Println(sa, sb)  // [1 22 3] [1 22 3]

	var slice1 []int
	fmt.Printf("%v %#v %T\n", slice1, slice1, slice1)  // [] []int(nil) []int
	slice1 = append(slice1, 1)  // 添加元素
	fmt.Printf("%v %#v %T\n", slice1, slice1, slice1)  // [1] []int{1} []int

	var Arr1 = [...]int{11, 22, 33, 44, 55}
	var Slice1 = Arr1[1:3]  // 根据数组定义切片  根据切片定义切片类似
	fmt.Printf("%v %#v %T\n", Slice1, Slice1, Slice1)  // [1 2 3 4 5] []int{1, 2, 3, 4, 5} []int

	// 切片的长度和容量
	fmt.Println(Arr1, len(Arr1), cap(Arr1)) // 5, 5  数组
	fmt.Println(Slice1, len(Slice1), cap(Slice1)) // 2, 4  切片  这里的4指的是Slice1从下标1开始到数组Arr1末尾的长度=4  (*)

	// 动态创建 make  make([]T, size, cap)
	var Slice2 = make([]int, 4, 5)
	fmt.Printf("%v %#v %T\n", Slice2, Slice2, Slice2)  // [0 0 0 0] []int{0, 0, 0, 0} []int
	fmt.Println(Slice2, len(Slice2), cap(Slice2)) // [0 0 0 0] 4 5
	Slice2[2] = 2
	fmt.Println(Slice2, len(Slice2), cap(Slice2)) // [0 0 2 0] 4 5
	Slice2 = append(Slice2, 11, 22)  // 给切片扩容元素
	fmt.Println(Slice2, len(Slice2), cap(Slice2)) // [0 0 2 0 11 22] 6 10  扩容 5->10

	Slice3 := []int{9, 8, 7, 6, 5, 4}
	Slice3 = append(Slice2, Slice3...)  // 给切片扩容切片 后面要加"..."
	fmt.Println(Slice3, len(Slice3), cap(Slice3)) // [0 0 2 0 11 22 9 8 7 6 5 4] 12 20

	// 删除切片的元素只能这样操作
	Slice4 := append(Slice3[:2], Slice3[3:]...) // 删除Slice3中的元素2了
	fmt.Println(Slice4, len(Slice4), cap(Slice4)) // [0 0 0 11 22 9 8 7 6 5 4] 11 20

	// golang 内置 sort() 提供了 int float64 string 数组 切片的排序
	sort.Ints(Slice3)  // 原地升序排序
	fmt.Println(Slice3, len(Slice3), cap(Slice3)) // [0 0 0 4 4 5 6 7 8 9 11 22] 12 20
	sort.Sort(sort.Reverse(sort.IntSlice(Slice3)))  // 降序 比较麻烦
	fmt.Println(Slice3, len(Slice3), cap(Slice3)) // [22 11 9 8 7 6 5 4 4 0 0 0] 12 20

	Arr2 := [...]int{1,5,8,2,4,3}  // 数组
	sort.Ints(Arr2[:])  // 原地升序排序
	fmt.Println(Arr2, len(Arr2), cap(Arr2)) // [1 2 3 4 5 8] 6 6
	sort.Sort(sort.Reverse(sort.IntSlice(Arr2[:])))  // 降序 比较麻烦
	fmt.Println(Arr2, len(Arr2), cap(Arr2)) // [8 5 4 3 2 1] 6 6

	Arr3 := []int{1, 4, 3}
	var Arr4 = make([]int, len(Arr3), cap(Arr3))
	copy(Arr4, Arr3)  // 深拷贝
	fmt.Println(Arr3, Arr4)  // [1 4 3] [1 4 3]
	Arr4[1] = 44
	fmt.Println(Arr3, Arr4)  // [1 4 3] [1 44 3]

	// map
	var maze1 = make(map[int]string)  // 使用make创建
	maze1[1] = "abc"
	maze1[3] = "def"
	fmt.Printf("%v %#v %T\n", maze1, maze1, maze1)

	maze2 := map[int]string {  // 直接创建
		1: "abc",
		3: "def",
	}
	fmt.Printf("%v %#v %T\n", maze2, maze2, maze2)

	for k, v := range maze2 {
		fmt.Println(k, v) // 1 abc 3 def 或者 3 def 1 abc  顺序不一定
	}

	maze2a, ok := maze2[1]
	fmt.Println(maze2a, ok)
	maze2b, ok := maze2[2]
	fmt.Println(maze2b, ok)  // 空 和 false

	fmt.Println(maze2)
	delete(maze2, 1)  // map数据的删除
	fmt.Println(maze2)

	var nslice = make([]map[int]string, 2, 2)  // map类型的切片
	if nslice[0] == nil {  // map默认为nil
		nslice[0] = make(map[int]string)
		nslice[0][1] = "qwe"
	}
	fmt.Println(nslice)
	var nmap = make(map[int][]string)  // 值为切片类型的map
	nmap[2] = []string{"a", "s", "d"}
	fmt.Println(nmap)

	// 按key升序排序输出map
	var maze3 = make(map[int]string)
	maze3[4] = "44"
	maze3[8] = "88"
	maze3[1] = "11"
	maze3[3] = "33"
	var slice3 []int
	for k, v := range maze3 {
		fmt.Printf("%v, %v\t", k, v)
		slice3 = append(slice3, k)
	}
	fmt.Println()

	sort.Ints(slice3)
	for _, key := range slice3 {
		fmt.Printf("%v, %v\t", key, maze3[key])
	}
	fmt.Println()

}
