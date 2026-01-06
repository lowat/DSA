class Solution:
    def two_sum(self, nums: list[int], target: int) -> list[int]:
        prev = {}
        for i, num in enumerate(nums):
            if target - num in prev:
                return [i, prev[target - num]]
            prev[num] = i
