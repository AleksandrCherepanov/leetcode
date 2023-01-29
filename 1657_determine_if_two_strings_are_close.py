from collections import defaultdict

def closeStrings(word1, word2):
    d1 = defaultdict(int)
    d2 = defaultdict(int)

    for c in word1:
        d1[c] += 1
    
    for c in word2:
        d2[c] += 1

    for k in d1:
        if k not in d2:
            return False

    for k in d2:
        if k not in d1:
            return False

    sl1 = sorted(d1.values())
    sl2 = sorted(d2.values())

    if len(sl1) != len(sl2):
        return False

    for i in range(len(sl1)):
        if sl1[i] != sl2[i]:
            return False

    return True

test = [
    [
        "abc",
        "bca",
        True
    ],
    [
        "a",
        "aa",
        False
    ],
    [
        "cabbba",
        "abbccc",
        True
    ],
    [
        "uau",
        "ssx",
        False
    ]
]

for i in test:
    result = closeStrings(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


