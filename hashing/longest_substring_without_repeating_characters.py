from collections import defaultdict

def lengthOfLongestSubstring(s):
    chrs = list(s)
    d = defaultdict(int) 
    l = 0
    r = 0
    result = 0

    while r < len(chrs):
        d[chrs[r]] += 1

        if d[chrs[r]] > 1:
            while d[chrs[r]] > 1:
                if d[chrs[l]] == 1:
                    del d[chrs[l]]
                else:
                    d[chrs[l]] -= 1
                l += 1
       
        result = max(result, len(d))
        r += 1

    return result

test = [
    [
        "abcabcbb",
        3
    ],
    [
        "bbbbb",
        1
    ],
    [
        "pwwkew",
        3
    ]
]

for i in test:
    result = lengthOfLongestSubstring(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

