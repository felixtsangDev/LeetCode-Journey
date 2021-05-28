package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/zigzag-conversion/submissions/
func convert(s string, numRows int) string {
	link := numRows - 2
	if link < 0 {
		link = 0
	}
	result := []byte{}
	column := int(math.Round(float64(len(s))/float64(numRows+link) + 0.5))
	for i := 0; i < numRows; i++ {
		n := 0
		count := 0
		for i+n < len(s) {
			if count < column {
				// fmt.Println("Append: ", string(s[i+n]), i+n)
				result = append(result, s[i+n])
				if i != 0 && i != numRows-1 && (i+n)+((numRows-1-i)*2) < len(s) {
					// fmt.Println("Append: ", string(s[(i+n)+((numRows-1-i)*2)]), (i+n)+((numRows-1-i)*2))
					result = append(result, s[i+n+((numRows-1-i)*2)])
				}
			}
			count++
			n += numRows + link
		}

	}

	return string(result)
}

func main() {
	result := convert("PAYPALISHIRING", 4)
	fmt.Println(result)
	fmt.Println("PINALSIGYAHRPI")

}
