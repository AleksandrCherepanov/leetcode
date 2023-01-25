from collections import defaultdict

def containsDuplicate(nums):
    d = defaultdict(int)
    for i in nums:
        d[i] += 1

        if d[i] > 1:
            return True

    return False

test = [
    [
        [1,2,3,1],
        True
    ],
    [
        [1,2,3,4],
        False
    ],
    [
        [1,1,1,3,3,4,3,2,4,2],
        True
    ]
]

for i in test:
    result = containsDuplicate(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

