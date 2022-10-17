package models

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
* 验证内存对齐
* 参考链接: https://geektutu.com/post/hpg-struct-alignment.html
 */
type v1 struct {
	a int8
	b string
	c int8
}

type v2 struct {
	a int8
	c int8
	b string
}

func TestStruct(t *testing.T) {
	s1 := v1{
		a: 10,
		b: "张三三",
		c: 20,
	}
	s2 := v2{
		a: 10,
		b: "张三三",
		c: 20,
	}

	fmt.Println(unsafe.Sizeof(s1), unsafe.Sizeof(s2))
}
