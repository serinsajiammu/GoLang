package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a, b int
	fmt.Print("\nEnter a and b:\n")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Println("Input is " + strconv.Itoa(a) + " and " + strconv.Itoa(b))
	var c float64 = float64(a) / float64(b)
	fmt.Println(c)
	fmt.Printf("Answer is %.2f\n", c)
	fmt.Println("Output is " + fmt.Sprint(c))
}
