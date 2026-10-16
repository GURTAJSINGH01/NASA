package main

import "fmt"

func main() {

	m := make(map[string]string)

	m["name"] = "gt"
	m["age"] = "20"

	fmt.Println(m["name"], m["age"])
}
