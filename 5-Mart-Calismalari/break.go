package main

import "fmt"

func main() {


		for i:=0; i<7; i++ {
			if i== 3 {
				continue
			} else if i ==5 {
				break
			}
			fmt.Println(i)
		}

}

