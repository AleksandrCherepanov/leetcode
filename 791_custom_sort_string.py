from collections import defaultdict

def customSortString(order, s):
    d = defaultdict(int)

    for i in range(len(order)):
        d[order[i]] = 0
    
    for c in s:
        d[c] += 1

    result = ''
    for k, v in d.items():
        result += k * v
    
    return result


test = [
    [
        "cba",
        "abcd",
        "cbad"
    ],
    [
        "cbafg",
        "abcd",
        "cbad"
    ],
    [
        "badc",
        "fatbcd",
        "badcft"
    ],
    [
        "kqep",
        "pekeq",
        "kqeep"
    ]
]

for i in test:
    result = customSortString(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


