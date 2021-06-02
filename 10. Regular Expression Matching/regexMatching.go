package main

import "fmt"

type patternType struct {
	IsZeroOrAll bool
	Value       string
}

func extractPattern(p string) (pattern []patternType) {
	temp := patternType{}
	for i, l := range p {
		if l == '*' {
			continue
		}
		if i+1 >= len(p) {
			temp.IsZeroOrAll = false
			temp.Value += string(l)
			pattern = append(pattern, temp)
			continue
		}
		if p[i+1] != '*' {
			if p[i+1] == '.' || p[i] == '.' {
				if temp.Value != "" {
					temp.IsZeroOrAll = false
					pattern = append(pattern, temp)
				}
				temp.IsZeroOrAll = false
				temp.Value = string(l)
				pattern = append(pattern, temp)
				temp.Value = ""
			} else {

				temp.Value += string(l)
			}
		} else {
			if temp.Value != "" {
				temp.IsZeroOrAll = false
				pattern = append(pattern, temp)
			}
			temp.IsZeroOrAll = true
			temp.Value = string(l)
			pattern = append(pattern, temp)
			temp.Value = ""
			if i+1 >= len(p) {
				break
			}
		}
	}
	return
}

func isMatch(s string, p string) bool {
	// oneOrManyIndex := []int{}
	pattern := extractPattern(p)
	index := 0
	check := false
	for z, pt := range pattern {
		for i := index; i < len(s); i++ {
			if pt.IsZeroOrAll {
				if pt.Value == "." {
					check = true
					index = i
					continue
				}
				if string(s[i]) == pt.Value {
					check = true
					index = i

					continue
				} else {
					index = i
					break
				}
			} else {
				if pt.Value == "." {
					check = true
					if index != len(s)-1 && z != len(pattern)-1 {
						index = i + 1
					}
					break
				}
				k := len(pt.Value) + i
				if k > len(s) || string(s[i:k]) != pt.Value {
					return false
				}
				check = true
				if len(pt.Value)+i >= len(s) && z != len(pattern)-1 {
					index = i + len(pt.Value)
				} else if len(pt.Value)+i >= len(s) || z == len(pattern)-1 {
					index = i + len(pt.Value) - 1
				} else {
					index = len(pt.Value) + i
				}
				break
			}
		}
	}
	if index != len(s)-1 || check == false {
		return false
	}
	return true
}

func main() {
	fmt.Println(isMatch("a", "ab*a"))                  //false
	fmt.Println(isMatch("a", "ab*"))                   //true
	fmt.Println(isMatch("aa", "."))                    //false
	fmt.Println(isMatch("mississippi", "mis*is*p*."))  //false
	fmt.Println(isMatch("mississippi", "mis*is*ip*.")) //true
	fmt.Println(isMatch("aaaa", "a*"))                 //true
	fmt.Println(isMatch("aa", "a"))                    // false
	fmt.Println(isMatch("aaa", "a.a"))                 //true
	fmt.Println(isMatch("ab", ".*"))                   //true
}
