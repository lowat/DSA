class Solution(object):
    def containsDuplicate(self, nums):
        res = set()
        for num in nums:
            if num in res:
                return True
            res.add(num)
        return False