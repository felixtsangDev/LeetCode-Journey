package main

import (
	"fmt"
)

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] > height[right] {
			if area := height[right] * (right - left); area > max {
				max = area
			}
			right--
		} else {
			if area := height[left] * (right - left); area > max {
				max = area
			}
			left++
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
}
