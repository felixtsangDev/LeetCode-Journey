def romanToInt(s: str) -> int:
    result = 0
    dict = {
        "I": 1,
        "IV": 4,
        "V": 5,
        "IX": 9,
        "X": 10,
        "XL": 40,
        "L": 50,
        "XC": 90,
        "C": 100,
        "CD": 400,
        "D": 500,
        "CM": 900,
        "M": 1000,
    }
    while len(s) > 0:
        if s[0:2] in dict:
            result += dict[s[0:2]]
            s = s[2::]
        else:
            result += dict[s[0:1]]
            s = s[1::]

    return result


def romanToInt2(s: str) -> int:
    result = 0
    dict = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    temp = 0
    for i in reversed(s):
        if dict[i] >= temp:
            result = result + dict[i]
        else:
            result = result - dict[i]
        temp = dict[i]

    return result


print(romanToInt2("IV"))
print("ANS: 4")
print(romanToInt2("IX"))
print("ANS: 9")
print(romanToInt2("LVIII"))
print("ANS: 58")
print(romanToInt2("MCMXCIV"))
print("ANS: 1994")
