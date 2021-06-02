from typing import List


def maxArea(height: List[int]) -> int:
    largestArea = 0
    left = 0
    right = len(height) - 1

    while left != right:
        width = right - left

        if height[left] > height[right]:
            temp = height[right] * width
            right -= 1
        else:
            temp = height[left] * width
            left += 1

        if temp > largestArea:
            largestArea = temp

    return largestArea


print(maxArea([2, 3, 4, 5, 18, 17, 6]))
print("Answer: 17")
print(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]))
print("Answer: 49")
print(maxArea([1, 2, 4, 3]))
print("Answer: 4")
