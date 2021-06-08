from typing import List


def longestCommonPrefix(strs: List[str]) -> str:
    if strs == None:
        return ""

    if len(strs[0]) == 0:
        return ""

    result = ""

    for i in range(len(strs[0])):
        for j in range(1, len(strs)):
            if i >= len(strs[j]) or strs[0][i] != strs[j][i]:
                return result
        result += strs[0][i]

    return result


print(longestCommonPrefix(["flower", "flow", "flight"]))
print("ANS: fl")
print(longestCommonPrefix(["dog", "racecar", "car"]))
print("ANS: ''")
print(longestCommonPrefix(["ab", "a"]))
print("ANS: a")
