class NumArray:

    def __init__(self, nums: List[int]):
        self.nums = nums
        self.sums = nums[:]
        for i in range(1, len(self.sums)):
            self.sums[i] += self.sums[i-1]

    def sumRange(self, left: int, right: int) -> int:
        leftSum = 0
        for i in range(0, left):
            leftSum += self.nums[i]

        return self.sums[right] - leftSum
