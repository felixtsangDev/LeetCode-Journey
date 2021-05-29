def reverse(x: int) -> int:
    isPos = True
    strTemp = ""
    if x < 0:
        isPos = False
        x *= -1

    if isPos:
        strTemp = str(x)[::-1]
    else:
        strTemp = "-" + str(x)[::-1]

    if x == 0 or int(strTemp).bit_length() > 31:
        return 0
    else:
        return int(strTemp)


print(reverse(3210))
