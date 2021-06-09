package main

import (
	"fmt"
	"time"
)

func threeSum(nums []int) (result [][]int) {
	start := time.Now()
	defer fmt.Println("Total removed visitors: Processing time: ", +time.Since(start))
	for i, num1 := range nums {
		for j, num2 := range nums[i+1:] {
			for k, num3 := range nums[j+i+2:] {
				tempj := j + i + 1
				tempk := k + tempj + 1
				if num1+num2+num3 == 0 && i != tempj && i != tempk && tempj != tempk {
					match := 0
					for _, tl := range result {
						if match == 3 {
							break
						}
						match = 0
						a, b, c := false, false, false
						for _, loopRI := range tl {
							if loopRI == num1 && !a {
								a = true
								match++
								continue
							}
							if loopRI == num2 && !b {
								b = true
								match++
								continue
							}
							if loopRI == num3 && !c {
								c = true
								match++
								continue
							}
						}
					}
					if match != 3 {
						temp := []int{}
						temp = append(temp, num1, num2, num3)
						result = append(result, temp)
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{0, 3, 0, 1, 1, -1, -5, -5, 3, -3, -3, 0}))
}
