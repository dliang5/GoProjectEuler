package main

import (
	"fmt" 
	"math"
)

func main(){
	ans := 0
	num := make([]int, 101)
	tempAns1 := 0
	tempAns2 := 0
	for i := 0; i < 101; i++ {
		num[i] = i
	}

	// time to find the difference between 1-100 ^2 and each individual 
	for i := 0; i < 101; i++ {
		tempAns1 += num[i]
		tempAns2 += int(math.Pow(float64(num[i]), 2))
	}

	ans = int(math.Pow(float64(tempAns1), 2)) - tempAns2
	

	fmt.Println(ans)
}