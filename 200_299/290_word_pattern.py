from collections import defaultdict

def wordPattern(pattern, s):
    d = defaultdict(str)
    wd = defaultdict(str)
    w = s.split(' ')

    if len(pattern) != len(w):
        return False

    for i in range(len(pattern)):
        if pattern[i] in d and d[pattern[i]] != w[i]:
            return False

        if w[i] in wd and w[i] != d[pattern[i]]:
            return False

        d[pattern[i]] = w[i]
        wd[w[i]] = pattern[i]

    return True

test = [
    [
        "abba",
        "dog cat cat dog",
        True
    ],
    [
        "abba",
        "dog cat cat fish",
        False
    ],
    [
        "aaaa",
        "dog cat cat dog",
        False
    ],
    [
        "abba",
        "dog dog dog dog",
        False
    ],
]

for i in test:
    result = wordPattern(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


