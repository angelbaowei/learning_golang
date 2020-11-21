package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

func main() {
	/*
		基本数据类型： int float bool string  (没有double, float分为float32 和 float64)
		复合数据类型： 数组 切片 struct 函数 map channel interface
	 */

	var a int8 = 1
	var b int16 = 2
	fmt.Printf("%v %T %v %T\n", a, a, b, b)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))  // a占 1 byte  b占 2 bytes

	// var c = a + b  error 因为a与b的类型不同  float同理 需要类型完全一致才可以
	var c = a + int8(b)
	fmt.Println(c)

	var d int = 17
	fmt.Println(unsafe.Sizeof(d)) // 64位机器 是 8 bytes  32位机器是 4 bytes
	fmt.Printf("%d %b %o %x", d, d, d, d)  // %d 十进制  %b 二进制  %o 八进制  %x 十六进制

	var f1 float32 = 3.4
	fmt.Printf("%v %f\n", f1, f1)  // %f 默认保留6位小数
	fmt.Println(unsafe.Sizeof(f1))  // float32 -- 4bytes
	f2 := float64(f1)
	fmt.Printf("%v %.1f\n", f2, f2)
	fmt.Println(unsafe.Sizeof(f2)) // float64 -- 8bytes

	f3 := 3.14e1
	fmt.Printf("%v %T, %v\n", f3, f3, unsafe.Sizeof(f3)) // 默认是64位

	m1 := 8.2
	m2 := 3.9
	fmt.Println(m1 - m2)  // 4.29999999999  浮点数精度丢失问题  可以使用第三方包 https://github.com/shopspring/decimal 来使用高精度的浮点数运算
	// 安装第三方包 go get github.com/shopspring/decimal
	d1 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))
	fmt.Println(d1)  // 4.3

	// 在判同的时候可使用 math.Abs(m1 - m2) < 1e-6 来判断二者相等
	m1 = 3.9
	fmt.Println(math.Abs(m1 - m2) < 1e-6)

	i1 := int(m1)
	fmt.Println(i1) // 3

	/*
		布尔类型默认值为false
		不允许强制类型转换： 其他不能强制转为布尔 布尔不能强制转为其他
		布尔类型无法参与数值运算
	 */

	var flag bool
	fmt.Printf("%v %T %t\n", flag, flag, flag)

	// int 默认值为 0  浮点型默认值为 0
	// 声明string
	var ss string
	fmt.Printf("%v %#v %T\n", ss, ss, ss)  //  "" string
	// 多行字符串
	str1 := `
  asdasd
	hfgh
	opjkop
	nkjn
das`
	fmt.Println(str1)  // 原样打印

	str1 = "aaaa"
	fmt.Println(len(str1)) // 4 bytes
	str1 = "你好"
	fmt.Println(len(str1))  // 6 一个汉字 3 bytes 是因为golang中汉字用utf-8编码

	str2 := "bbbb"
	str3 := str1 + str2  // 直接拼接字符串
	fmt.Println(str3)
	str3 = fmt.Sprintf("%v --- %v", str2, str1)  // 格式化拼接字符串
	fmt.Println(str3)

	// string 包函数
	// Split-分割 strings.Join()-合并 Contains-是否包含
	// HasPrefix-前缀包含 HasSuffix-后缀包含
	// Index-下标 LastIndex()-最后一个的下标  没有则返回-1
	// Replace-替换 ReplaceAll-替换全部

	str4 := strings.Split(str3, "---")
	fmt.Println(str4, len(str4), str4[0], str4[1])  // 切片[]

	str5 := strings.Join(str4, "*")
	fmt.Println(str5) // 将切片添加拼接符拼接在一起

	arr := []string{"a", "b", "c"} 	// 切片
	fmt.Println(strings.Join(arr, "-")) // a-b-c

	str1 = "abcde"
	str2 = "bcd"
	str3 = "ab"
	str6 := "de"
	fmt.Println(strings.Contains(str1, str2))
	fmt.Println(strings.HasPrefix(str1, str3)) // str1 前缀 是否包含str3
	fmt.Println(strings.HasSuffix(str1, str6)) // str1 后缀 是否包含str6

	str1 = "abacabdsag"
	fmt.Println(strings.Index(str1, "ab"))  // 0
	fmt.Println(strings.LastIndex(str1, "ab")) // 4

	str2 = strings.Replace(str1, "ab", "cd", 1)  // n=1表示替换第一处
	fmt.Println(str2)
	str2 = strings.Replace(str1, "ab", "cd", 2)  // n=2表示替换前两处
	fmt.Println(str2)
	str2 = strings.ReplaceAll(str1, "ab", "cd") // 表示替换全部
	fmt.Println(str2)

	// 类型转换
	// 其他转string
	// 方法1
	var i int = 3
	Stri := fmt.Sprintf("%d", i)  // int to string
	var f = 3.16
	Strf := fmt.Sprintf("%f", f)  // float to string
	fmt.Printf("%v %T\n", Stri, Stri)
	fmt.Printf("%v %T\n", Strf, Strf)
	// 方法2
	Stri = strconv.FormatInt(int64(i), 10)  // base表示10进制的int
	fmt.Printf("%v %T\n", Stri, Stri)
	Strf = strconv.FormatFloat(float64(f), 'f', 1, 64)  // 参数：要转换的值、格式化类型('f' 'b' 'e')、保留小数点位数、格式化类型(32 or 64)
	fmt.Printf("%v %T\n", Strf, Strf)  // 3.2 string
	f = 3.15
	Strf = strconv.FormatFloat(float64(f), 'f', 1, 64)
	fmt.Printf("%v %T\n", Strf, Strf)  // 3.1 string
	// string转其他
	ii, err := strconv.ParseInt(Stri, 10, 64)
	fmt.Printf("%v %T %v\n", ii, ii, err)  // 3 int64 nil
	Stri = "3xx"
	ii, err = strconv.ParseInt(Stri, 10, 64)
	fmt.Printf("%v %T %v\n", ii, ii, err)  // 0 int64 strconv.ParseInt: parsing "3xx": invalid syntax
	ff, err := strconv.ParseFloat(Strf, 64)
	fmt.Printf("%v %T %v\n", ff, ff, err)  // 3.1 float64 nil

	// 字符 byte 和 rune  用单引号
	// 字符有两种类型： byte类型-ascii码   rune类型-utf8字符
	var za = 'a'
	fmt.Printf("%v %T %c\n", za, za, za)  // 97 int32 a
	str1 = "abcasfasfasfasfasfasdasd"
	fmt.Printf("%v %T %c\n", str1[2], str1[2], str1[2])  // 99 uint8 c
	fmt.Println(unsafe.Sizeof(str1), len(str1)) // 16 24  unsafe.Sizeof(str1)算的是string的结构体大小：因为Golang中的sring内部实现由两部分组成，一部分是指向字符串起始地址的指针，另一部分是字符串的长度，两部分各是8字节，所以一共16字节
	// 因此 unsafe.Sizeof 无法查看string类型占用的空间

	za = '国'  // 汉字是utf-8编码
	fmt.Printf("%v %T %c\n", za, za, za)  // Unicode编码后的值：22269 int32 a  utf-8是Unicode的实现方式之一

	str1 = "asd"
	for i := 0; i < len(str1); i++ {  // len = 3
		fmt.Printf("%v-%c\t", str1[i], str1[i])
	}
	fmt.Println()
	str1 = "asd法"
	for i := 0; i < len(str1); i++ {  // len = 6
		fmt.Printf("%v-%c\t", str1[i], str1[i]) // 有汉字 此种遍历方法是按ascii解析的 不正确
	}
	fmt.Println()
	for idx, r := range str1 {  // len = 4   range 按照rune类型解析 即utf8解析
		fmt.Printf("%v:%v-%c\t", idx, r, r)
	}
	fmt.Println()

	// golang中string不能直接修改 需要 转换成[]byte/rune切片-修改-转回来
	// str1[1] = 'b' error
	byteStr1 := []byte(str1)
	byteStr1[1] = 'b'
	str1 = string(byteStr1)
	fmt.Println(str1)

	runeStr1 := []rune(str1)
	runeStr1[3] = '哈'
	str1 = string(runeStr1)
	fmt.Println(str1)

	// % 余数=被除数-(被除数/除数)*除数
	fmt.Println(10 % 3, -10 % 3, 10 % -3)
	// golang 中 ++ -- 只能单独使用 不能像C语言那样 且只有i++/i-- 没有++i/--i
	// ^ 表示 异或 或者 按位取反
	fmt.Println(5 ^ 6, ^5)  // 3 -6
}
