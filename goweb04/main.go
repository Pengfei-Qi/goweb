package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2. 调用模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("template parse file failed! err:%v \n", err)
		return
	}
	//3. 写入数据
	name := "王老五"
	t.Execute(w, name)
}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed! err:%v \n", err)
		return
	}
}
