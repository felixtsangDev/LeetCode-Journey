package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	isPositive := true
	if x < 0 {
		isPositive = false
		x *= -1
	}
	s := strconv.Itoa(x)
	temp := ""
	for _, i := range s {
		temp = string(i) + temp
	}
	if !isPositive {
		temp = "-" + temp
	}
	result, _ := strconv.Atoi(temp)
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func main() {
	result := reverse(1534236469)
	fmt.Println(result)
}
