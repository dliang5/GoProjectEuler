
package main

import (
	"fmt"
)

func main() {
	sum := 0
	prev := 1
	cur := 2
	var new int
	for cur < 4000000 {
		if cur%2 == 0 {
			sum += cur
		}
		new = cur + prev
		prev = cur
		cur = new
	}
	fmt.Println(sum)
}
