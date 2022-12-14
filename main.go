// Package mmain provides basic usecase of various data types and variables
package main

import (
	"fmt"
)

func main() {
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
