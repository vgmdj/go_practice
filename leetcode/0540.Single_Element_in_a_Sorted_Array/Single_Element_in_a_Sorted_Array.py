from typing import List


class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        left,right = 0, len(nums)-1
        while left < right:
            mid = (left+right)//2
            mid -= mid&1
            if nums[mid] == nums[mid+1] :
                left = mid +2
            else:
                right = mid
        return nums[right]
