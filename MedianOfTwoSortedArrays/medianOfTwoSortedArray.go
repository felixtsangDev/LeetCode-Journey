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
	for _, a:= nums2 {
		mid := math.Round(len(nums1))
		nums1 := append(num1)
		if a > mid{
			if a < nums1[len(nums1)-1]{
				// append between
				copy()
			}else{
				// append after
			}
		}else{
			if a > nums1[0]{
				// append between
			}else{
				// append after
			}
		}
	}
}

func main() {
	testcases := []test{
		test{array1: []int{1,5, 9}, array2: []int{2,4, 7}},
		// test{array1: []int{1, 2}, array2: []int{3, 4}},
		// test{array1: []int{0, 0}, array2: []int{0, 0}},
		// test{array1: []int{}, array2: []int{1}},
		// test{array1: []int{2}, array2: []int{}},
	}
	for _, testcase := range testcases {
		fmt.Println(findMedianSortedArrays(testcase.array1, testcase.array2))
	}
}
