from collections import defaultdict

def findLucky(arr):
    maxValue = -1
    freq = defaultdict(int)

    for a in arr:
        freq[a] += 1

    for k, f in freq.items():
        if k == f:
            maxValue = max(maxValue, k)

    return maxValue

test = [
    [
        [2,2,3,4],
        2
    ],
    [
        [1,2,2,3,3,3],
        3
    ],
    [
        [2,2,2,3,3],
        -1
    ]
]

for i in test:
    result = findLucky(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

