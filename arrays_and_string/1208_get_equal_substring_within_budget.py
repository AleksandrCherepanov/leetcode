def equalSubstring(s, t, maxCost):
    l = 0
    r = 0
    
    maxLen = 0
    cost = 0
    while r < len(s):
        cost += abs(ord(s[r]) - ord(t[r]))
          
        while cost > maxCost:
            cost -= abs(ord(s[l]) - ord(t[l]))
            l += 1

        maxLen = max(r - l + 1, maxLen)

        r += 1 

    return maxLen

test = [
    [
        "abcd",
        "bcdf",
        3,
        3
    ],
    [
        "abcd",
        "cdef",
        3,
        1
    ],
    [
        "abcd",
        "acde",
        0,
        1
    ],
]

for i in test:
    result = equalSubstring(i[0], i[1], i[2])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

