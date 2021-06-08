def intToRoman(num: int) -> str:
    result = ""
    value = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
    romans = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
    temp = 0
    while num > 0:
        if value[temp] > num:
            temp += 1
        else:
            result += romans[temp]
            num -= value[temp]

    return result


print(intToRoman(4))
print("ANS: IV")
print(intToRoman(9))
print("ANS: IX")
print(intToRoman(58))
print("ANS: LVIII")
print(intToRoman(1994))
print("ANS: MCMXCIV")
