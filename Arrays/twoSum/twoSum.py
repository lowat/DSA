class Solution(object):
    def twoSum(self, nums, target):
        hm = dict()
        for i, num in enumerate(nums):
            if target - num in hm:
                return [hm[target - num], i]
            else:
                hm[num] = i