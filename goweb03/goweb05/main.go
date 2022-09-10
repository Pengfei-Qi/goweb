package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/**
html/template 与 text/template 区别
1. 使用方法上面区别不大
2. html/template 可预防 xss 跨站脚本攻击
*/

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板,并且修改默认的标识符
	files, err := template.New("v2.tmpl").Delims("{[", "]}").ParseFiles("./v2.tmpl")
	if err != nil {
		fmt.Printf("parse file failed! err : %v", err)
		return
	}

	//渲染模板
	msg := "小王子"
	files.Execute(w, msg)
}
func xss(w http.ResponseWriter, r *http.Request) {
	//定义模板

	//解析模板,对内容特殊处理
	files, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Printf("parse file failed! err : %v", err)
		return
	}
	//渲染模板
	//html/template针对的是需要返回HTML内容的场景，在模板渲染过程中会对一些有风险的内容进行转义，以此来防范跨站脚本攻击。
	msg1 := "<script>alert('嘿嘿嘿')</script>"
	msg2 := "<a href='http://image.uc.cn/s/wemedia/s/upload/2022/be68d16f804d0f07b0b57a2c9dbd9fcd.png'>显示图片</a>"
	files.Execute(w, map[string]string{
		"msg1": msg1,
		"msg2": msg2,
	})
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("parse file failed, err:%v", err)
		return
	}
}
