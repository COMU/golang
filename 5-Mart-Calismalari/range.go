package main

import "fmt"

func main() {

	var pow = [] int {1,2,4,8}
	for i,v := range pow {
		fmt.Printf("2 ** %d = %d\n",i,v)
	}
}

