package main

import (
	"fmt"
	"reflect"
)

var (
	s = "hello"
	i = 34
	a = true
)

func main() {

	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(a))
}
