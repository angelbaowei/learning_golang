package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1 创建模版 创建一个.tmpl文件

	// 2 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("template parse error %v\n", err)
		return
	}
	// 3 渲染模版
	my := Person{
		Name: "mzb",
		Age: 25,
	}
	err = t.Execute(w, my)  //name 替换了hello.tmpl中的 "."

	if err != nil {
		fmt.Printf("template execute error %v\n", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("http serve failed! %v\n", err)
		return
	}

}
