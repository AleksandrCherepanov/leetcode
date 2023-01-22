from collections import defaultdict

def largestUniqueNumber(nums):
    occur = defaultdict(int)

    for n in nums:
        occur[n] += 1

    result = -1 
    for k, v in occur.items():
        if v == 1:
            result = max(result, k)

    return result

test = [
    [
        [5,7,3,9,4,9,8,3,1],
        8
    ],
    [
        [9,9,8,8],
        -1
    ]
]

for i in test:
    result = largestUniqueNumber(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

