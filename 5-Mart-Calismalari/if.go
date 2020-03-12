package main

import "fmt"

func main() {

	a:=10
	b:=10
	if b>a {
		fmt.Println("b büyük a")
	} else if b==a {
		fmt.Println("eşittir")
	} else {
		fmt.Println("Küçüktür")
	}
}

