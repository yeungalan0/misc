from typing import List


class Solution:
    # Brute force solution
    # Time complexity: O(n^2)
    # Space complexity: O(n)
    def twoSum_slow(self, nums: List[int], target: int) -> List[int]:
        for i, num in enumerate(nums):
            compliment = self.find_match(i+1, nums[i+1:], target - num)
            if compliment != None:
                return [i, compliment[0]]

    def find_match(self, start_index: int, nums: List[int], target: int) -> (int, int):
        for i, num in enumerate(nums, start=start_index):
            if num == target:
                return (i, num)

        return None

    # Most optimal big O solution
    # Time complexity: O(n)
    # Space complexity: O(n)
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        compliments = {}

        for i, num in enumerate(nums):
            compliment = target - num

            if compliment in compliments:
                return [compliments[compliment], i]

            compliments[num] = i

        return None


if __name__ == "__main__":
    out = Solution().twoSum([2, 7, 11, 15], 9)
    print(out)
