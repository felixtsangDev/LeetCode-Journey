# def isPalindrome(x: int) -> bool:
#     if x < 0 or (x > 0 and x % 10 == 0):
#         return False

#     revert = 0

#     while x > revert:
#         revert = revert * 10 + x % 10
#         x = x // 10

#     return x == revert or x == revert // 10


def isPalindrome(x: int) -> bool:
    return True if (x > 0 and str(x) == str(x)[::-1]) else False


print(isPalindrome(96969))
