class Solution:
    def decodeString(self, s: str) -> str:
        start: int = -1
        end: int = -1
        num_open: int = 0
        num_str: str = ""
        for i, c in enumerate(s):
            if c == "[":
                if start == -1:
                    start = i
                num_open += 1
            if c == "]":
                if num_open == 1:
                    end = i
                num_open -= 1
            if end > start >= 0:
                break
            if start == -1:
                if c.isnumeric():
                    num_str += c
                else:
                    num_str = ""
        if start == -1:
            return s
        new_str: str = s[:start-1] + self.decodeString(s[start+1:end]) * int(num_str)
        if end+1 < len(s):
            new_str += self.decodeString(s[end+1:])
        return new_str
    
    
s = Solution()
print(s.decodeString("3[a2[c]]"))
