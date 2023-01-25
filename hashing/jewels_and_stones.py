from collections import defaultdict

def numJewelsInStones(jewels, stones):
    j = defaultdict(int)

    for jew in jewels:
        j[jew] = 1

    result = 0
    for s in stones:
        result += j[s]
    
    return result

test = [
    [
        "aA",
        "aAAbbbb",
        3
    ],
    [
        "z",
        "ZZ",
        0
    ]
]

for i in test:
    result = numJewelsInStones(i[0], i[1])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

