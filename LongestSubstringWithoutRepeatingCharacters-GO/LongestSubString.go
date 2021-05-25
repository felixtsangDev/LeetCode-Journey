package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/submissions/
func lengthOfLongestSubstring(s string) int {
	var result [128]bool
	max := 0
	count := 0
	end := 0
	for i, _ := range s[end:] {
		for j, a := range s[i:] {
			if !result[int(a)] {
				count += 1
				if count > max {
					max = count
				}
				result[int(a)] = true
			} else {
				count = 0
				var reset [128]bool
				result = reset
				end = j
				break
			}

		}
	}
	return max
}

func main() {
	result := lengthOfLongestSubstring("123456789 !@#$%^&*()")
	fmt.Println(result)
}
