package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	cont, readErr := ioutil.ReadFile("./hello.txt")
	if readErr != nil {
		fmt.Printf("ReadFile error %v\n", readErr)
		return
	}
	num, err := fmt.Fprintln(w, string(cont))  // w 发送给client(回复)

	if err != nil {
		fmt.Printf("Responce error %v\n", err)
		return
	}
	fmt.Println(num)
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("http serve failed! %v\n", err)
		return
	}

}
