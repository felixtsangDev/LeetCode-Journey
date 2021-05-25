
def lengthOfLongestSubstring(s: str) -> int:
    curr = 0
    longest = 0
    char_map = {}

    for i in range(len(s)):
        print(f"{s[i]}")
        if s[i] in char_map and curr <= char_map[s[i]]:
            print(f"curr {curr}  char_map[s[i]] {char_map[s[i]]}")
            curr = char_map[s[i]] + 1
            print(f"curr {curr}")
        else:
            print(
                f"i {i} curr {curr} longest {longest} i - curr + 1 {i - curr + 1}")
            longest = max(longest, i - curr + 1)

        char_map[s[i]] = i
    return longest


print(lengthOfLongestSubstring("    "))
