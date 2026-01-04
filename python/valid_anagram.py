class Solution:
    def is_valid_anagram(self, s: str, t:str) -> bool:
        if len(s) != len(t):
            return False
        s_char_count = [0 for i in range(26)]
        t_char_count = [0 for i in range(26)]
        for i in range(len(s)):
            s_char_idx = ord(s[i]) - ord('a')
            t_char_idx = ord(t[i]) - ord('a')
            s_char_count[s_char_idx] += 1
            t_char_count[t_char_idx] += 1

        for i in range(26):
            if s_char_count[i] != t_char_count[i]:
                return False

        return True
