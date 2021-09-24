class Solution:
    def lengthOfLongestSubstring_old(self, s: str) -> int:
        longest_substring = 0

        for i in range(len(s)):
            substring_length = self.get_substring_length(i, s)
            if substring_length > longest_substring:
                longest_substring = substring_length

        return longest_substring

    def get_substring_length(self, index: int, string: str) -> int:
        substring_set = set()
        for i in range(index, len(string)):
            if string[i] not in substring_set:
                substring_set.add(string[i])
            else:
                break

        return len(substring_set)

    def get_unique_substring_chars(self, index: int, substring_set: set, string: str) -> set:
        farthest_index_thus_far = index + len(substring_set)

        for i in range(farthest_index_thus_far, len(string)):
            if string[i] not in substring_set:
                substring_set.add(string[i])
            else:
                break

        return substring_set

    # Optimizations:
    # longest_substring > len(s[i:]) end
    # skip and start at the longest chain

    def lengthOfLongestSubstring_old2(self, s: str) -> int:
        longest_substring = 0
        substring_char_set = set()

        for i in range(len(s)):  # there's repeating work here and you can skip cetain indexes
            if longest_substring >= len(s) - i:
                break

            if i > 0 and s[i-1] in substring_char_set:
                substring_char_set.remove(s[i-1])

            substring_char_set = self.get_unique_substring_chars(
                i, substring_char_set, s)
            if len(substring_char_set) > longest_substring:
                longest_substring = len(substring_char_set)

        return longest_substring

    # Internalized and made this solution more readable (to me): https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/236649/Python-solution
    def lengthOfLongestSubstring(self, s: str) -> int:
        last_seen_index_dic = {}
        longest_substring = 0
        head_pointer = 0

        for tail_pointer in range(len(s)):
            if s[tail_pointer] in last_seen_index_dic:
                # move window down if it includes the character at the tail
                head_pointer = max(
                    head_pointer, last_seen_index_dic[s[tail_pointer]] + 1)

            last_seen_index_dic[s[tail_pointer]] = tail_pointer

            longest_substring = max(
                longest_substring, tail_pointer - head_pointer + 1)

        return longest_substring


if __name__ == "__main__":
    out = Solution().lengthOfLongestSubstring("pwwkew")
    print(out)
