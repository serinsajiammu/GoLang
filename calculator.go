package main

import "fmt"

func main() {

	var a, b, x int
	fmt.Print("\nEnter a and b and choice\n1.add\n2.subtract\n3.multiply\n4.divide\n5.modulus\n")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("\nThe input is %d and %d\n", a, b)
	fmt.Scanf("%d", &x)
	switch x {
	case 1:
		fmt.Println(a + b)
	case 2:
		fmt.Println(a - b)
	case 3:
		fmt.Println(a * b)
	case 4:
		fmt.Println(a / b)
	case 5:
		fmt.Println(a % b)
	}

}
