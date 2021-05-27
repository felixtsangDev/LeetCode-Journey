package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	max := 0
	result := [2]int{}
	for i, c := range s {
		for j, d := range s[i:] {
			actualIndex := j + i
			if c == d {
				l := 0
				diff := false
				for k := actualIndex; k > i; k-- {
					if s[i+l] != s[k] {
						diff = true
						break
					}
					l++
				}
				if !diff && actualIndex-i > max {
					max = actualIndex - i
					result = [2]int{i, actualIndex}
				}
			}
		}
	}
	return s[result[0] : result[1]+1]
}

func main() {
	temp := longestPalindrome("baaaabd")
	fmt.Println(temp)
}
