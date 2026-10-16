package main

import "fmt"

/// there is no while keyword

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)

		i = i + 1

		/// i++ is showing the error
	}

}
