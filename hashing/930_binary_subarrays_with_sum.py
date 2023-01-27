from collections import defaultdict

def numSubarraysWithSum(nums, goal):
    d = defaultdict(int)
    d[0] = 1
    ans = curr = 0

    for n in nums:
        curr += n
        ans += d[curr - goal]
        d[curr] += 1

    return ans

test = [
    [
        [1, 0, 1, 0, 1],
        2,
        4
    ],
    [
        [0, 0, 0, 0, 0],
        0,
        15
    ]
]

for i in test:
    result = numSubarraysWithSum(i[0], i[1])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

