class Solution(object):
    def groupAnagrams(self, strs):
        ans = collections.defaultdict(list)
        for word in strs:
            cCount = [0] * 26
            for c in word:
                cCount[ord(c) - ord('a')] += 1
            key = tuple(cCount)
            ans[key].append(word)
        return ans.values()