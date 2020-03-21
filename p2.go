package main

import "fmt"

func main() {
	var i int = 10
	var s string = "hello"
	var a = "this is Go language"
	var fname, lname string = "Serin", "Saji"
	age := 20
	p, q, r := 5, 10, 15
	Subject, marks := "GoLang", 50
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(fname + lname)
	fmt.Println(age)
	fmt.Println(p + q + r)
	fmt.Println(Subject, "-", marks)
}
