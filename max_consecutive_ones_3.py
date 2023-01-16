def longestOnes(nums, k):
    zeros = 0
    i = 0
    j = 0
    result = 0

    while i < len(nums):
        if nums[i] == 0:
            zeros += 1

        while zeros > k and j <= i:
            if nums[j] == 0:
                zeros -= 1
            j += 1

        result = max(result, i - j + 1)
        i += 1

    return result

print(longestOnes([1,1,1,0,0,0,1,1,1,1,0], 2))
print(longestOnes([0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], 3))
print(longestOnes([0,0,0,0], 0))
print(longestOnes([0,0,0,0], 1))

