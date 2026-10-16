package main

import "fmt"

func main() {

	if age := 19; age > 18 {
		fmt.Println("adult")
	} else if age > 10 {
		fmt.Println("is teenageer")
	}
}
