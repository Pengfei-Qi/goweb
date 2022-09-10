package main

import "fmt"

func zeroval(ival int) {
	ival = 5
	fmt.Println("修改变量:",ival)
}

func zeroptr(iptr *int) {
	*iptr = 3
}

func main() {
	i := 2
	fmt.Println("initial:", i)//1

	zeroval(i)
	fmt.Println("zeroval:", i)//0

	zeroptr(&i)
	fmt.Println("zeroptr:", i)//0

	fmt.Println("pointer:", &i)//0
}