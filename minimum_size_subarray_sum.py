def minSubArrayLen(target, nums):
    l = 0
    r = 0
   
    s = 0
    f = False
    result = len(nums)
    while r < len(nums):
        s += nums[r]
        while s >= target:
            f = True
            s -= nums[l]
            result = min(result, r - l + 1)
            l += 1

        r += 1

    return result if f else 0

test = [
    [
        7,
        [2,3,1,2,4,3],
        2
    ],
    [
        4,
        [1,4,4],
        1
    ],
    [
        11,
        [1,1,1,1,1,1,1,1],
        0
    ],
]

for i in test:
    result = minSubArrayLen(i[0], i[1])

    if result != i[2]:
        print("Error. ", result, i[2])
    else:
        print("OK")


