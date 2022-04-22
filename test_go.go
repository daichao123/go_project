package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hello world")
	//
	//name := "test"
	//
	//fmt.Println(name)
	//var (
	//	name  string
	//	age   int
	//	check bool
	//	money big.Float
	//)
	//var a int
	//var b int
	const (
		NAME string = "tom"
		AGE  int    = 29
	)
	//	fmt.Printf("NAME: %v \n", AGE)
	//	var str_test string = `
	//	asdasd1
	//asdasd2
	//asdasd3
	//`
	//println(str_test)
	a := 0100
	b := 20
	a = 20
	b = 222
	if a > b {
		fmt.Printf("a:%v\n", a)
	}
	var arr1 []int
	var arr = [...]int{1, 2, 2}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i=: %v\n", i)
		fmt.Printf("v=: %v\n", arr[i])
	}
	arr1 = append(arr1, 1, 2, 3)
	fmt.Printf("arr1= %v\n", arr1)

	fmt.Printf("arr:%T\n", arr)
	fmt.Printf("arr:%T\n", arr1)
	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)
	//name := "tom"
	//age := "asd"
	//s := strings.Join([]string{name, age}, "")
	//println(s)
	for _, v := range arr1 {
		println(v)
	}
}
