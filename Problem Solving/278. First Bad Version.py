# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int: # n as right
        left   = 0
        right  = n
        result = 0
        
        while right >= left:
            middle = (left + right) // 2
            
            if isBadVersion(middle) is True:
                result = middle
                right = middle - 1
            else:
                left = middle + 1
        
        return result
        