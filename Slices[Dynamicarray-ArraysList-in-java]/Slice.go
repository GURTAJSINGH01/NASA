package main

import (
	"fmt"
)

func main() {
	// var num = make([]int, 2, 5)
	// fmt.Println(num)
	// num = append(num, 1)
	// num = append(num, 2)
	// num = append(num, 3)
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// num = append(num, 4)

	// num = append(num, 5)
	// fmt.Println(num)
	// fmt.Println(cap(num))

	// num := []int{1, 3, 2, 44, 3}
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// copying of one source slice to the another

	// var num1 = make([]int, 0, 5)
	// num1 = append(num1, 4)

	// var num2 = make([]int, len(num1))

	// copy(num2, num1)

	// fmt.Println(num1, num2)

	// num := []int{1, 3, 1, 43, 4, 2, 4}
	// print the range of the number from the given slice

	// fmt.Println(num[2:4])

	/// Slice package
	// we can directly compare the two slices with this method
	// var num = []int{1, 3, 1, 43, 4, 2, 4}
	// var num2 = []int{1, 3, 1, 43, 4, 2, 4}

	// fmt.Println(slices.Equal(num, num2))

	/// 2d Slices

	var num = [2][2]int{{1, 2}, {1, 2}}
	fmt.Println(num)

}
