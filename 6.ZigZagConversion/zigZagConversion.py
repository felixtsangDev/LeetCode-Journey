def convert(s: str, numRows: int) -> str:
    if numRows == 1 or len(s) < numRows:
        return s
    row_num = 1
    step = 1  # decide go down(1) or up(-1)
    rows = {}
    for char in s:
        if row_num not in rows:
            rows[row_num] = char
        else:
            rows[row_num] += char
        row_num += step

        # Change direction when hit top or bottom
        if row_num == 1 or row_num == numRows:
            step *= -1
        # print(rows)

    result = "".join(rows.values())
    return result


print(convert("PAYPALISHIRING", 4))
