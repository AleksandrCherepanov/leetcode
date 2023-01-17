def pivotIndex(nums):
    _sum = 0
    leftsum = 0
    for i in range(0, len(nums)):
        _sum += nums[i]

    for i in range(0, len(nums)):
        if leftsum == _sum - leftsum - nums[i]:
            return i
        leftsum += nums[i] 

    return -1

test = [
    [
        [1,7,3,6,5,6],
        3
    ],
    [
        [1,2,3],
        -1
    ],
    [
        [2,1,-1],
        0
    ],
]

for i in test:
    result = pivotIndex(i[0])

    if result != i[-1]:
        print("Error. ", result, i[-1])
    else:
        print("OK")

