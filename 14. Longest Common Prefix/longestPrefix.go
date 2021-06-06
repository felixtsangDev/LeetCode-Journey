package main

import "fmt"

type result struct {
	key   string
	value int
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	counter := make(map[string]int)
	result := &result{key: strs[0], value: 1}
	temp := ""
	for _, s := range strs[0] {
		temp += string(s)
		counter[temp] += 1
	}
	for _, s := range strs[1:] {
		if s == "" {
			return ""
		}
		temp := ""
		for i, c := range s {
			temp += string(c)
			if counter[temp] == 0 {
				if i == 0 {
					return ""
				}
				break
			} else {
				counter[temp] += 1
				if counter[temp] >= result.value {
					result.key = temp
					result.value = counter[temp]
				}
			}
		}
	}
	if result.value == 1 {
		return ""
	}
	return result.key
}

func main() {
	// fmt.Println(longestCommonPrefix([]string{"int", "aaa", "na"}))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"b", "a"}))
	fmt.Println(longestCommonPrefix([]string{"aaa", "aa", "aaa"}))
}
