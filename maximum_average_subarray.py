def findMaxAverage(nums, k):
    i = 0
    j = 0
    maxAvg = None
    s = 0
    while i < len(nums):
        s += nums[i]
        if i - j + 1 == k:
            avg = s / (i - j + 1)  
            maxAvg = avg if maxAvg is None else max(maxAvg, avg)
            s -= nums[j]
            j += 1
        
        i += 1
    return maxAvg

print(findMaxAverage([1,12,-5,-6,50,3], 4))
print(findMaxAverage([5], 1))
print(findMaxAverage([-1], 1))
