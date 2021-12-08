from typing import Dict, List
from utils import parse_input
import random


def solve1(problem_input: List[str]) -> int:
    count_singles = 0

    for input in problem_input:
        outputs = input.split(" ")[11:]
        for output in outputs:
            if len(output) == 2 or len(output) == 4 or len(output) == 3 or len(output) == 7:
                count_singles += 1

    return count_singles


def solve2(problem_input: List[str]) -> int:
    decoded_sum = 0

    for p_input in problem_input:
        seen = set()
        inputs = p_input.split(" ")[:10]
        outputs = p_input.split(" ")[11:]
        index_to_letter_dict = {"a": 0, "b": 1, "c": 2,
                                "d": 3, "e": 4, "f": 5, "g": 6}
        segment_guess_list = [1, 2, 3, 5, 7, 11, 13]

        while not is_valid(segment_guess_list, index_to_letter_dict, inputs):
            seen.add(tuple(segment_guess_list))

            while tuple(segment_guess_list) in seen:
                random.shuffle(segment_guess_list)

        seen = set()
        decoded_sum += get_value(outputs,
                                 segment_guess_list, index_to_letter_dict)

    return decoded_sum


def is_valid(segment_guess_list: List[int], index_to_letter_dict: Dict[str, int], inputs: List[str]) -> bool:
    valid_values = [13, 28, 32, 21, 33, 40, 14, 42, 35, 37]

    for digit_repr in inputs:
        value = 0
        for letter in digit_repr:
            i = index_to_letter_dict[letter]
            value += segment_guess_list[i]

        if value not in valid_values:
            return False

    return True


def get_value(outputs: List[str], segment_value_list: List[int], letters_to_index: Dict[str, int]) -> int:
    valid_values = [37, 13, 28, 32, 21, 33, 40, 14, 42, 35]

    value = ""
    for digit_repr in outputs:
        digit_value = 0
        for letter in digit_repr:
            i = letters_to_index[letter]
            digit_value += segment_value_list[i]

        value += str(valid_values.index(digit_value))

    return int(value)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day8.txt")

    # problem_input = [int(num) for num in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
