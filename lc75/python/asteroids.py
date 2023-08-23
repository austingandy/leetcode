from typing import List


class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        found_pos = False
        last_neg = 0
        care: int = -1
        i = 0
        while i < len(asteroids):
            if not found_pos and asteroids[i] > 0:
                found_pos = True
            if found_pos and asteroids[i] < 0:
                asteroids, numPopped = move_left(asteroids, i)
                i -= numPopped
            i += 1
        return asteroids


# moves the presumed-negative value at index i left until it reaches a stop
# or explodes. Returns the modified list, the updated index of the left-most
# negative value
def move_left(a: List[int], i: int) -> (List[int], int):
    j = i - 1
    num_popped = 0
    while j >= 0 and a[j] > 0:
        curr = a[j]
        if curr <= a[i] * -1:
            a.pop(j)
            i -= 1
            num_popped += 1
        if a[i] * -1 <= curr:
            a.pop(i)
            num_popped += 1
            break
        j -= 1
    return a, num_popped


s = Solution()
print(s.asteroidCollision([8,-8]))
