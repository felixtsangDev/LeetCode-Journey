def myAtoi(s: str) -> int:
    sign = 1
    index = 0
    num = 0
    stripped = s.strip()
    if not stripped:
        return 0
    if stripped[index] == "+":
        index += 1
    elif stripped[index] == "-":
        sign = -1
        index += 1

    while index < len(stripped):
        if not stripped[index].isdigit():
            break
        num = num * 10 + ord(stripped[index]) - ord("0")
        if num > 2 ** 31:
            break
        index += 1
    return min(max(sign * num, -(2 ** 31)), 2 ** 31 - 1)


print(myAtoi("-91283472332"))
