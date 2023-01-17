def maxVowels(s, k):
    vowels = ["a","e","i","o","u"]
    l = 0
    r = 0

    maxV = 0
    cur = 0
    while r < len(s):
        if s[r] in vowels:
            cur += 1
      
        while r - l + 1 > k:
            if s[l] in vowels:
                cur -= 1
            l += 1

        maxV = max(maxV, cur)
        if maxV == k:
            return maxV

        r += 1
        
    return maxV

test = [
    [
        "abciiidef",
        3,
        3
    ],
    [
        "aeiou",
        2,
        2
    ],
    [
        "leetcode",
        3,
        2
    ],
]

for i in test:
    result = maxVowels(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


