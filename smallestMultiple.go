package main

import (
	"fmt"
)
func funBruteForce(input int, num ...int) bool{
	isDivisible := true
	for _, val := range num{
		if input % val != 0{
			isDivisible = false
			break
		}
	}

	return isDivisible

}
func main(){
	smallest := 21
	isSmallest := false
	for {
		isSmallest = funBruteForce(smallest,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,28,19,20)
		if isSmallest == true {
			break
		} 
		smallest++
	}

	fmt.Println(smallest) 
}