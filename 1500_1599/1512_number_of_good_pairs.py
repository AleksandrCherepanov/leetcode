from collections import defaultdict

def numIdenticalPairs(nums):
    d = defaultdict(int)

    for n in nums:
        d[n] += 1

    s = 0 
    for v in d.values():
        s += int((v * (v - 1)) / 2)

    return s

test = [
    [
        [1,2,3,1,1,3],
        4
    ],
    [
        [1,1,1,1],
        6
    ],
    [
        [1,2,3],
        0
    ]
]

for i in test:
    result = numIdenticalPairs(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

