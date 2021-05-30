package main

import (
	"fmt"
	"math"
)

func rangeCheck(num int) bool {
	if num > math.MaxInt32 || num < math.MinInt32 {
		return false
	} else {
		return true
	}
}

func myAtoi(s string) int {
	start := false
	isPositive := true
	resultArr := []int{}
	for i := 0; i < len(s); i++ {
		charInt := int(s[i])
		if charInt >= 48 && charInt <= 57 {
			if !start {
				start = true
			}
			if charInt == 48 && len(resultArr) == 0 {
				continue
			}
			resultArr = append(resultArr, charInt-48)
			continue
		}
		if charInt == 45 || charInt == 32 || charInt == 43 {
			if start {
				break
			}
			if charInt == 45 {
				isPositive = false
				start = true
				continue
			} else if charInt == 43 {
				start = true
				continue
			} else {
				continue
			}
		}
		break
	}
	result := 0
	if len(resultArr) > 19 {
		if isPositive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	for i, d := range resultArr {
		result += d * int(math.Pow(10, float64(len(resultArr))-1-float64(i)))
	}
	if !isPositive {
		result *= -1
	}
	if rangeCheck(result) {
		return result
	} else {
		if isPositive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
}

func main() {
	fmt.Println(myAtoi("00-1"))
}
