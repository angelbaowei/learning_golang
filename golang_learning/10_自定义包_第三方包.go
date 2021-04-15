package main  // 入口

// 首先要使用 go mod init 创建项目

import (
	T "awesomeProject/calc"
	"fmt"
	//_ "awesomeProject/calc"  这是匿名包 此时这个包就不会参与编译

)

// init 函数 优先于main方法执行
// 且init函数执行顺序是与导入包的顺序相反  (见word文档)
func init() {
	fmt.Println("10_自定义包_第三方包.go")
}

func main() {

	res := T.Add(10, 2)
	res2 := T.Sub(10, 2)
	fmt.Println(res, res2)
	fmt.Println(T.AAA)
	//fmt.Println(T.aaa) // error 私有变量不能访问

	// 第三方包地址： https://pkg.go.dev/
	/*
		安装第三方包  1. go get github.com/shopspring/decimal
			    2. 先初始化 go mod init 再在项目目录下 go mod download 依赖下载到Go_Path
	    		    3. 在项目目录下 go mod vendor 下载且将依赖复制到当前项目目录里面
	 */


}
