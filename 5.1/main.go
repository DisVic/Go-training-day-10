package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		_, _ = fmt.Scan(&number)
		array[i] += number
	}
	num := array[n-1]
	if n > 1 {
		for i := n - 1; i > 0; i-- {
			array[i] = array[i-1]
		}
		array[0] = num
	}
	k := 0
	for k < n {
		fmt.Print(array[k], " ")
		k++
	}
}
