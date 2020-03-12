package  main

import "fmt"

func main() {
	toplam :=1
	for toplam <10 {
		toplam +=toplam
		fmt.Println(toplam)
	}
}

