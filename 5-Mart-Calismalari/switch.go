package main

import "fmt"

func main() {

	var score float64

	fmt.Println("Sınav notu gir")
	fmt.Scanf("%v",&score)
	switch {
	case score <=49:
		fmt.Println("f")
	default:
		fmt.Println("geçerli bir şey gir")
	}
}

