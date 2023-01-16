def moveZeroes(nums):
    i = 0
    j = 0
    while i < len(nums):
        if nums[i] != 0:
            t = nums[i - j]
            nums[i - j] = nums[i]
            nums[i] = t
        else:
            j += 1

        i += 1

test = [
    [
        [0,1,0,3,12],
        [1,3,12,0,0]
    ],
    [
        [0],
        [0]
    ],
    [
        [0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0]
    ],
    [
        [1, 2, 0, 3, 4],
        [1, 2, 3, 4, 0]
    ]
]

for i in test:
    nums = i[0]
    moveZeroes(nums)

    if nums != i[1]:
        print("Error. ", i[0], i[1])
    else:
        print("OK")


