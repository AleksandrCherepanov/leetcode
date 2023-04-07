from collections import defaultdict

def isIsomorphic(s, t):
    smap = defaultdict(str)
    tmap = defaultdict(str)

    for i in range(len(s)):
        if s[i] in smap and smap[s[i]] != t[i]:
            return False

        if t[i] in tmap and tmap[t[i]] != s[i]:
            return False

        smap[s[i]] = t[i]
        tmap[t[i]] = s[i]

    return True

test = [
    [
        "egg",
        "add",
        True
    ],
    [
        "foo",
        "bar",
        False
    ],
    [
        "paper",
        "title",
        True
    ],
    [
        "bbbaaaba",
        "aaabbbba",
        False
    ],
    [
        "badc",
        "baba",
        False
    ]
]

for i in test:
    result = isIsomorphic(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


