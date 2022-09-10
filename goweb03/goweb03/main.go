package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义模板
	m2 := func(name string) (string, error) {
		return name + "真好看", nil
	}

	//解析模板
	files := template.New("v1.tmpl")

	files.Funcs(template.FuncMap{
		"m2": m2,
	})
	_, err := files.ParseFiles("./v1.tmpl")
	if err != nil {
		fmt.Printf("parse file failed! err:%v ", err)
		return
	}
	name := "小王子"
	//渲染模板
	files.Execute(w, name)
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	files, err := template.ParseFiles("./t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("parse file failed! err:%v ", err)
		return
	}
	//渲染模板
	name := "小王子"
	files.Execute(w, name)
}

func main() {

	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", tmpl)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed! err:%v ", err)
		return
	}
}
