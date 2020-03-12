package main

import "fmt"

func main() {

	var colors = []string {"Red","green","blue"}
	fmt.Println(colors)
	colors=append(colors,"Purple")
	fmt.Println(colors)
}

