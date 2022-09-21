// Package mmain provides basic usecase of various data types and variables
package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program is going to help the engineer understand different data types")
	// var myInt int32
	// var myFloat float32
	// var myString string
	// const myConstant = 55
	/*for j := 1; j != 17; j++ {
		fmt.Println(j)
	}*/
	/*mySlice := []int32{1, 2, 3, 4, 5}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println(mySlice[i])
	}
	for j := range mySlice {
		fmt.Println(mySlice[j])
	}*/
	mySliceString := []string{"manish", "temp", "you", "here", "while", "for", "temp", "manish", "you"}
	myMap := make(map[string]int, 0)
	for s := range mySliceString {
		tmpString := mySliceString[s]
		if _, ok := myMap[tmpString]; !ok {
			fmt.Println("Found a new key, adding to the map: ", tmpString)
			myMap[tmpString] = 1
		} else {
			tmpvalue := myMap[tmpString]
			myMap[tmpString] = tmpvalue + 1
		}
	}
	fmt.Println("------------------------------------")
	for key, value := range myMap {
		fmt.Printf("The word \"%s\" is found %d times\n", key, value)
	}
}
