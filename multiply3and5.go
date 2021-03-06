package main

import "fmt"

/*
Find the sum of all the multiples of 3 or 5 below 1000
*/

func main() {
	// example the range is 0..10
	total := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	fmt.Println(total)
}
