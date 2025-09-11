package utils

import "fmt"

func HelperFunction() {
	fmt.Println("Helper function from public package.")
}

func MakeArray() [10]int {
	var arr [10]int
	return arr
}
