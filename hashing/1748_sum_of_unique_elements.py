from collections import defaultdict

def sumOfUnique(nums):
    uniq = defaultdict(int) 

    for n in nums:
        uniq[n] += 1

    s = 0
    for k, u in uniq.items():
        if u == 1:
            s += k

    return s

test = [
    [
        [1,2,3,2],
        4
    ],
    [
        [1,1,1,1,1],
        0
    ],
    [
        [1,2,3,4,5],
        15
    ]
]

for i in test:
    result = sumOfUnique(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

