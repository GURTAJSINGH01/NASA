package main

import "fmt"

func main() {
	for i := range 11 {
		if i < 3 {
			fmt.Println("it is less than 3")
		}

	}
	fmt.Println("Greater than 3 ")
}
