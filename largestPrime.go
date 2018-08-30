package main

import "fmt"

func main() {
	num := 600851475143
	newNum := num
	largestPrime := 0
	counter := 2

	for counter*counter <= newNum {
		if newNum%counter == 0 {
			newNum = newNum / counter
			largestPrime = counter
		} else {
			counter++
		}
	}
	if newNum > largestPrime {
		largestPrime = newNum
	}
	fmt.Println(num, newNum, largestPrime)
}
