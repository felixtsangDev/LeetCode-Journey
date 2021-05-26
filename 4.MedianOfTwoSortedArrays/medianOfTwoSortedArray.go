package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/

type test struct {
	array1 []int
	array2 []int
}

func findMedianSortedArrays(nums1 []int, nums2 []int) (median float64) {
	l2 := len(nums2)
	l1 := len(nums1)
	isOdd := (l1 + l2) % 2
	length := int(math.Round((float64(l1+l2) + 1) / 2))
	anchor1 := 0
	anchor2 := 0
	check := false
	for i := 0; i < length; i++ {
		if isOdd != 1 && (i == length-2 || i == length-1) {
			check = true
		} else if i == length-1 {
			check = true
		}
		if anchor2+1 > l2 {
			if check {
				median += float64(nums1[anchor1])
			}
			anchor1++
			continue
		}
		if anchor1+1 > l1 {
			if check {
				median += float64(nums2[anchor2])
			}
			anchor2++
			continue
		}
		if nums1[anchor1] >= nums2[anchor2] {
			if check {
				median += float64(nums2[anchor2])
			}
			anchor2++
		} else {
			if check {
				median += float64(nums1[anchor1])
			}
			anchor1++
		}
	}
	if isOdd != 1 {
		median = median / 2
	}
	return
}

func main() {
	testcases := []test{
		// test{array1: []int{1, 5, 9}, array2: []int{2, 4, 7}},
		test{array1: []int{1, 2}, array2: []int{}},
		// test{array1: []int{0, 0}, array2: []int{0, 0}},
		test{array1: []int{}, array2: []int{2, 3}},
		// test{array1: []int{2}, array2: []int{}},
	}
	for _, testcase := range testcases {
		temp := findMedianSortedArrays(testcase.array1, testcase.array2)
		fmt.Println(temp)
	}
}
