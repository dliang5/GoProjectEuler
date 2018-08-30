package main

import (
	"fmt"
	"strconv"
)

/*
	Finding the largest palindrome product in 3 digits
	ex. in 2 digits
		9009 = 91 * 99
*/
func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// testing here to make sure everything is good to go
	// before algorithm design here hehehe
	answer := 101
	strAnswer := strconv.Itoa(answer)
	fmt.Println(isPalindrome(strAnswer))
	// what i would do is start from 999 and move down
	// cuz that should in theory give me the largest value from two
	// three digits faster than from the bottom up
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			currentValue := i * j
			if isPalindrome(strconv.Itoa(currentValue)) == true {
				if currentValue > answer {
					answer = currentValue
				}
				break
			}
		}
	}

	fmt.Println(answer)
}
