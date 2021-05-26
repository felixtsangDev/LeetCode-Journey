def longestPalindrome(s: str) -> str:
    longest = ""
    for i in range(len(s)):
        temp1 = findPalindrome(s, i, i)
        temp2 = findPalindrome(s, i, i+1)
        longest = max([longest, temp1, temp2], key=lambda x: len(x))
        # print(f"longest {longest}")
    return longest


def findPalindrome(s: str, left: int, right: int) -> str:
    while left >= 0 and right < len(s) and s[left] == s[right]:
        # print(f"left: {left} {s[left]} right: {right} {s[right]}")
        left -= 1
        right += 1

    # print(f"result: {s[left+1:right]}")

    return s[left+1:right]


print(f"longestPalindrome: {longestPalindrome('abcdefgfedcba')}")
