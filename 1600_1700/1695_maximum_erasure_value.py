from collections import defaultdict

def maximumUniqueSubarray(nums):
    d = defaultdict(int)
    l = r = 0
    s = 0
    ms = 0

    while (r < len(nums)):
        d[nums[r]] += 1
        s += nums[r]
        
        while d[nums[r]] > 1:
            d[nums[l]] -= 1
            s -= nums[l]
            l += 1
        
        print(s)
        ms = max(ms, s)
        r += 1
        
    return ms

test = [
    [
        [4,2,4,5,6],
        17
    ],
    [
        [5,2,1,2,5,2,1,2,5],
        8
    ]
]

for i in test:
    result = maximumUniqueSubarray(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

