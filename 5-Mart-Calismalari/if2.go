package main

import "fmt"

func main() {

	if foo:=2; foo==1 {
		fmt.Println("bar")
	} else {
		fmt.Println("buz")
		fmt.Println(foo)
	}
}

