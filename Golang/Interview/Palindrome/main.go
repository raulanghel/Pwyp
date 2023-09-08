package main

import (
	"fmt"
	"strings"
)

func main() {

	var originalString string = "Madam"
	if isPalindrome(originalString) {
		fmt.Println("The given string is Palindrome")
	} else {
		fmt.Println("The given string is NOT a Palindrome")
	}
}

func isPalindrome(originalString string) bool {
	var reverseString string = ""
	var length = len(originalString)

	for i := length - 1; i >= 0; i-- {
		reverseString = reverseString + string(originalString[i])
	}

	// Case in-sensitive comparision
	if strings.EqualFold(strings.ToLower(originalString), strings.ToLower(reverseString)) {
		return true
	} else {
		return false
	}
}
