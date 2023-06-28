from collections import defaultdict

def checkInclusion(s1, s2):
    d1 = defaultdict(int)
    d2 = defaultdict(int) 

    for c in s1:
        d1[c] += 1
    
    l = r = 0
    
    s = list(s2)
    while r < len(s2):
        d2[s[r]] += 1

        while r - l + 1 > len(s1):
            d2[s[l]] -= 1
            l += 1
        
        isBreak = False
        for c in d1:
            if d2[c] != d1[c]:
                isBreak = True
                break

        if not isBreak:
            return True

        r += 1

    return False

test = [
    [
        "abdoooo",
        "eidbaooo",
        False
    ],
    [
        "oooa",
        "eidboaoo",
        True
    ],
    [
        "a",
        "ab",
        True
    ]
]

for i in test:
    result = checkInclusion(i[0], i[1])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

