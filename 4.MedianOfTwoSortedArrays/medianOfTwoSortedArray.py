from typing import List
import math


def findMedianSortedArrays(nums1: List[int], nums2: List[int]) -> float:
    joint_nums = nums1 + nums2
    joint_nums.sort()
    if (len(joint_nums) % 2 == 0):
        return (joint_nums[int(len(joint_nums)/2)-1] + joint_nums[int(len(joint_nums)/2)])/2
    else:
        return joint_nums[math.floor(len(joint_nums)/2)]

#
# def findMedianSortedArrays2(nums1: List[int], nums2: List[int]) -> float:
#     joint_nums = nums1 + nums2
#     joint_nums.sort()
#     if (len(joint_nums) % 2 == 0):
#         return (joint_nums[int(len(joint_nums)/2)-1] + joint_nums[int(len(joint_nums)/2)])/2
#     else:
#         return joint_nums[math.floor(len(joint_nums)/2)]


print(findMedianSortedArrays([1, 2, 4], [3, 5]))
