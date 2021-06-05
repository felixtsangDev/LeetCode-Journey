package main

import "fmt"

func romanToInt(s string) int {
	result := 0
	for i, c := range s {
		switch c {
		case 'I':
			result++
		case 'V':
			if i > 0 && s[i-1] == 'I' {
				result += 3
			} else {
				result += 5
			}
		case 'X':
			if i > 0 && s[i-1] == 'I' {
				result += 8
			} else {
				result += 10
			}
		case 'L':
			if i > 0 && s[i-1] == 'X' {
				result += 30
			} else {
				result += 50
			}
		case 'C':
			if i > 0 && s[i-1] == 'X' {
				result += 80
			} else {
				result += 100
			}
		case 'D':
			if i > 0 && s[i-1] == 'C' {
				result += 300
			} else {
				result += 500
			}
		case 'M':
			if i > 0 && s[i-1] == 'C' {
				result += 800
			} else {
				result += 1000
			}
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
