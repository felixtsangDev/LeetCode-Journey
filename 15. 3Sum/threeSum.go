package main

import (
	"fmt"
	"time"
)

func getCombination(result [][]int, arr []int, searchMap map[int]bool) [][]int {
	temp := [][]int{}
	checkedMap := make(map[int]bool)
	for i, num := range arr {
		if !checkedMap[num] {
			for _, num2 := range arr[i+1:] {
				revert := (num + num2) * -1
				if searchMap[revert] {
					isExist := false
					for _, r := range temp {
						if (num == r[0] && num2 == r[2]) || (num == r[2] && num2 == r[0]) {
							isExist = true
							break
						}
					}
					if !isExist {
						temp = append(temp, []int{num, revert, num2})
						result = append(result, []int{num, revert, num2})
					}
				}
			}
		}
		checkedMap[num] = true
	}
	return result
}
func threeSum(nums []int) (result [][]int) {
	start := time.Now()
	defer fmt.Println("Total removed visitors: Processing time: ", +time.Since(start))
	positiveNums, negativeNums, zeroNums := []int{}, []int{}, []int{}
	numMap := make(map[int]bool)
	for _, num := range nums {
		switch {
		case num > 0:
			numMap[num] = true
			positiveNums = append(positiveNums, num)
		case num < 0:
			numMap[num] = true
			negativeNums = append(negativeNums, num)
		default:
			zeroNums = append(zeroNums, num)
		}
	}
	if len(zeroNums) >= 3 {
		result = append(result, []int{0, 0, 0})
	}
	temp := make(map[int]bool)
	for _, num := range negativeNums {
		pos := num * -1
		if len(zeroNums) > 0 {
			if numMap[pos] && !temp[pos] {
				temp[pos] = true
				result = append(result, []int{num, 0, pos})
			}
		}
	}
	result = getCombination(result, negativeNums, numMap)
	result = getCombination(result, positiveNums, numMap)
	return result
}

func main() {
	fmt.Println(threeSum([]int{0, 3, 0, 1, 1, -1, -5, -5, 3, -3, -3, 0}))
	fmt.Println(threeSum([]int{-3, -3, 6, 0, 3}))
	fmt.Println(threeSum([]int{1, 1, -2, -1, -1, 2, -3, 4, -4}))
	// fmt.Println(threeSum([]int{-1}))
}
