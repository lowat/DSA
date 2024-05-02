class Solution(object):
    def isAnagram(self, s, t):
        sCounts = self.charCount(s)
        tCounts = self.charCount(t)
        for i in range(len(sCounts)):
            if sCounts[i] != tCounts[i]:
                return False
        return True
    
    def charCount(self, word):
        result = [0] * 26
        for c in word:
            result[ord('a') - ord(c)] += 1
        return result
        