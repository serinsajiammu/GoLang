package main

import "fmt"

func main() {
	var fname, place string
	var age int
	fmt.Print("Enter your first name:")
	fmt.Scanln(&fname)
	fmt.Print("\nEnter your place:")
	fmt.Scanln(&place)
	fmt.Print("\nEnter yoyr age:")
	fmt.Scanf("%d", &age)
	fmt.Printf("I am %s from %s\n", fname, place)
	fmt.Printf("Age:%d", age)
}
