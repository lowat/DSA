class Solution(object):
    def isPalindrome(self, s):
        if not s:
            return False
        cleaned_string = self.remove_non_alpha_num_chars(s.lower())
        print(cleaned_string)
        l, r = 0, len(cleaned_string) - 1
        while l < r:
            if cleaned_string[l] != cleaned_string[r]:
                return False
            l += 1
            r -= 1
        return True
    

    def remove_non_alpha_num_chars(self, s):
        res = []
        for c in s:
            if (ord(c) >= ord('a') and ord(c) <= ord('z')) or (ord(c) >= ord('0') and ord(c) <= ord('9')):
                res.append(c)
        return ''.join(res)