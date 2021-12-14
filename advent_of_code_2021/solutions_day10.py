from typing import List
from utils import parse_input


def solve1(problem_input: List[str]) -> int:
    matching_dict = {"(": ")", "[": "]", "{": "}", "<": ">"}
    points_dict = {")": 3, "]": 57, "}": 1197, ">": 25137}
    closing_symbols = matching_dict.values()
    syntax_error_score = 0

    for line in problem_input:
        stack: List[str] = []
        for symbol in line:
            # try:
            #     if symbol in closing_symbols:
            #         opening = stack.pop()

            if symbol in closing_symbols:
                if len(stack) > 0:
                    opening = stack.pop()
                    if matching_dict[opening] != symbol:
                        syntax_error_score += points_dict[symbol]
                        break
                else:
                    syntax_error_score += points_dict[symbol]
                    break
            else:
                stack.append(symbol)

    return syntax_error_score


def solve2(problem_input: List[str]) -> int:
    matching_dict = {"(": ")", "[": "]", "{": "}", "<": ">"}
    closing_symbols = matching_dict.values()
    incomplete_sequences = []

    for line in problem_input:
        stack: List[str] = []
        is_valid = True
        for symbol in line:
            if symbol in closing_symbols:
                if len(stack) > 0:
                    opening = stack.pop()
                    if matching_dict[opening] != symbol:
                        is_valid = False
                else:
                    is_valid = False
            else:
                stack.append(symbol)

        if len(stack) > 0 and is_valid:
            incomplete_sequences.append(stack)

    return calculate_score(incomplete_sequences)


def calculate_score(incomplete_sequences: List[List[str]]) -> int:
    points_dict = {"(": 1, "[": 2, "{": 3, "<": 4}
    scores = []

    for incomplete_sequence in incomplete_sequences:
        score = 0
        for symbol in reversed(incomplete_sequence):
            score *= 5
            score += points_dict[symbol]

        scores.append(score)

    scores.sort()
    median = len(scores) // 2

    return scores[median]


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day10.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
