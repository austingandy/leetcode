class Solution:
    # each char will eject the closest opposite char to its left
    # or to its right if none are available to its left
    def predictPartyVictory(self, senate: str) -> str:
        # number of this color that we have encountered so far, with a vote still
        # to be cast (will decrement after a ban)
        color_count = {
            "D": 0,
            "R": 0
        }
        s = list(senate)
        while True:
            found = {
                "D": False,
                "R": False
            }
            i = 0
            while i < len(s):
                other = opposite(s[i])
                if color_count[other] > 0:
                    color_count[other] -= 1
                    s.pop(i)
                    continue
                else:
                    color_count[s[i]] += 1
                    found[s[i]] = True
                    i += 1
            if not found["D"] or not found["R"]:
                break
        return "Dire" if found["D"] else "Radiant"


def opposite(c: str) -> str:
    if c == "D":
        return "R"
    else:
        return "D"
    

s = Solution()
print(s.predictPartyVictory("DDRRR"))
