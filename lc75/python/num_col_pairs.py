from typing import List


def transpose(grid: List[List[int]]) -> List[List[int]]:
    g: List[List[int]] = [[0]*len(grid) for _ in range(len(grid))]
    g[0][0] = 4
    for i in range(len(grid)):
        for j in range(len(grid)):
            g[i][j] = grid[j][i]
    return g


def row_is_equal(this: List[int], that: List[int]):
    if len(this) != len(that):
        return False
    for i in range(len(this)):
        if this[i] != that[i]:
            return False
    return True


def helper(l: List[str]):
    i = -1
    for j in range(len(l)):
        if l[j] == "*":
            i = j
            break
    if i == -1:
        return l
    return helper(l[0:i-1] + l[i+1:])


class Solution:
    def __init__(self):
        pass
    
    def equalPairs(self, grid: List[List[int]]) -> int:
        g = transpose(grid)
        d = {}
        d.setdefault(0)
        count: int = 0
        for row in grid:
            s = str(row)
            if s in d:
                d[s] += 1
            else:
                d[s] = 1
        for col in g:
            s = str(col)
            if s not in d:
                continue
            count += d[s]
        return count

    def removeStars(self, s: str) -> str:
        return helper(list(s))


s: Solution = Solution()
print(s.equalPairs([[3, 2, 1], [1, 7, 6], [2, 7, 7]]))


