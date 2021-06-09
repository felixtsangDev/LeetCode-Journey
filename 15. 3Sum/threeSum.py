from typing import List


def threeSum(nums: List[int]) -> List[List[int]]:
    nums.sort()
    result = []
    added = set()

    for i in range(len(nums)):
        first = i
        left = i + 1
        right = len(nums) - 1
        while left < right:
            sum = nums[first] + nums[left] + nums[right]
            if sum == 0:
                if (nums[first], nums[left], nums[right]) not in added:
                    result.append([nums[first], nums[left], nums[right]])
                    added.add((nums[first], nums[left], nums[right]))
                left += 1
            elif sum > 0:
                right -= 1
            else:
                left += 1

    return result


def threeSumFinal(nums: List[int]) -> List[List[int]]:
    result = set()

    posNum, negNum, zeroNum = [], [], []

    for num in nums:
        if num > 0:
            posNum.append(num)
        elif num < 0:
            negNum.append(num)
        else:
            zeroNum.append(num)

    posNumSet, negNumSet = set(posNum), set(negNum)

    # print({"pos": posNum, "neg": negNum, "zero": zeroNum})
    # 1. check if have three 0 in zeroNum
    if len(zeroNum) >= 3:
        result.add((0, 0, 0))

    # 2. check if have 0 , pos and neg have same
    if len(zeroNum) >= 1:
        for num in posNumSet:
            if -num in negNumSet:
                result.add((-num, 0, num))

    # 3. if no 0, have two sum of pos = neg
    for i in range(len(posNum)):
        for j in range(i + 1, len(posNum)):
            sum = posNum[i] + posNum[j]
            if -sum in negNumSet:
                result.add(tuple(sorted([-sum, posNum[i], posNum[j]])))

    # 4. if no 0, have two sum of neg = pos
    for i in range(len(negNum)):
        for j in range(i + 1, len(negNum)):
            sum = negNum[i] + negNum[j]
            if -sum in posNumSet:
                result.add(tuple(sorted([-sum, negNum[i], negNum[j]])))

    return result


print(threeSumFinal([-1, 0, 1, 2, -1, -4]))
print([[-1, -1, 2], [-1, 0, 1]])
print(threeSumFinal([-2, 0, 0, 2, 2]))
print([[-2, 0, 2]])
