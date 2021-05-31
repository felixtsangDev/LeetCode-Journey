package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	length := int(math.Log10(float64(x)))
	temp := x
	for i := 0; i < length+1; i++ {
		pow := int(math.Pow(10, float64(length-i)))
		if x/pow != temp%10 {
			return false
		}
		x -= pow * (x / pow)
		temp /= 10
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
