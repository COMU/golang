package main

import "fmt"

func main() {
	numbers := make([] int,2,5)
	numbers[0]=1
	numbers[1]=2
	fmt.Println(numbers)

	numbers=append(numbers,3)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
}

