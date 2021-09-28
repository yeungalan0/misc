from typing import List

class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = True
        last_offset = len(digits) - 1

        while carry and last_offset > -1:
            new_val = digits[last_offset] + 1

            carry = new_val > 9

            digits[last_offset] = new_val % 10

            last_offset -= 1

        # Hit case where there's a carry left after the most significant digit
        if carry:
            digits.insert(0, 1)

        return digits

if __name__ == "__main__":
    out = Solution().plusOne([9, 9, 9])

    print(out)



        