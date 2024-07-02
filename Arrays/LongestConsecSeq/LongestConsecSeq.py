class Solution(object):
    def longestConsecutive(self, nums):
        nums_set, max_seq_length, curr_seq_length = set(nums), 0, 0
        for num in nums:
            curr_seq_length = 0
            if num - 1 in nums_set:
                continue
            curr_seq_length += 1
            curr_num = num + 1
            while curr_num in nums_set:
                curr_seq_length += 1
                curr_num += 1
            max_seq_length = max(max_seq_length, curr_seq_length)
        return max(max_seq_length, curr_seq_length)