def isMatch(s, p):
    s, p = " " + s, " " + p
    lenS, lenP = len(s), len(p)
    dp = [[False] * (lenP) for i in range(lenS)]
    dp[0][0] = True

    # Write the first row of the dp
    for j in range(1, lenP):
        if p[j] == "*":
            dp[0][j] = dp[0][j - 2]
    print(dp[0])

    for i in range(1, lenS):
        for j in range(1, lenP):
            if p[j] in {s[i], "."}:
                dp[i][j] = dp[i - 1][j - 1]
            elif p[j] == "*":
                # 1. check the result of 2 char backward
                # 2. check the result above

                dp[i][j] = dp[i][j - 2] or dp[i - 1][j] and p[j - 1] in {s[i], "."}

        print(dp[i])

    return bool(dp[-1][-1])


def isMatchArray(s, p):
    s, p = " " + s, " " + p
    lenS, lenP = len(s), len(p)
    prev = [True]
    for j in range(1, lenP):
        prev.append(p[j] == "*" and prev[j - 2])

    for i in range(1, lenS):
        curr = [False]
        for j in range(1, lenP):

            if p[j] in {s[i], "."}:
                curr.append(prev[j - 1])
            elif p[j] == "*":
                curr.append(curr[j - 2] or prev[j] and p[j - 1] in {s[i], "."})
            else:
                curr.append(False)

        print(curr)
        prev = curr

    return prev[-1]


print(isMatch("aa", "a*baaaaaaaab*.*"))
print(isMatchArray("aa", "a*baaaaaaaab*.*"))
