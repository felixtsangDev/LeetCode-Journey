package main

import "fmt"

func intToRoman(num int) string {
	result := ""
	for num != 0 {
		switch {
		case num < 10:
			for num != 0 {
				if num == 4 {
					result += "IV"
					num = 0
					break
				}
				if num == 9 {
					result += "IX"
					num = 0
					break
				}
				if num > 4 {
					result += "V"
					num -= 5
					if num == 0 {
						break
					}
				}
				result += "I"
				num--
			}
		case num < 100:
			times := int(num / 10)
			num = num - (10 * times)
			for times != 0 {
				if times == 9 {
					result += "XC"
					times = 0
					break
				}
				if times == 4 {
					result += "XL"
					times = 0
					break
				}
				if times > 4 {
					result += "L"
					times -= 5
					if times == 0 {
						break
					}
				}
				result += "X"
				times--
			}
		case num < 1000:
			times := int(num / 100)
			num = num - (100 * times)
			for times != 0 {
				if times == 9 {
					result += "CM"
					times = 0
					break
				}
				if times == 4 {
					result += "CD"
					times = 0
					break
				}
				if times > 4 {
					result += "D"
					times -= 5
					if times == 0 {
						break
					}
				}
				result += "C"
				times--
			}
		case num >= 1000:
			times := int(num / 1000)
			num = num - (1000 * times)
			for times != 0 {
				result += "M"
				times--
			}

		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(1000))
	fmt.Println(intToRoman(100))
	fmt.Println(intToRoman(99994))
}
