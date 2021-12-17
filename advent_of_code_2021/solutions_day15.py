from typing import List
from utils import parse_input
from copy import deepcopy


def tabulate_min_path(grid: List[List[int]]):
    memo = deepcopy(grid)

    memo[0][0] = 0

    for r in range(len(memo)):
        for c in range(len(memo[0])):
            if r == 0 and c == 0:
                continue

            top_val = memo[r-1][c] if r-1 >= 0 else 10000000000000000
            left_val = memo[r][c-1] if c-1 >= 0 else 10000000000000000

            memo[r][c] = grid[r][c] + \
                min(top_val, left_val)

    for row in memo:
        print(row)

    return memo[-1][-1]


def solve1(problem_input: List[List[int]]) -> int:
    return tabulate_min_path(problem_input)


def solve2(problem_input: List[List[int]]) -> int:
    orig_height = len(problem_input)
    orig_width = len(problem_input[0])

    for r in range(orig_height):
        for _ in range(4):
            for _ in range(orig_width):
                value = problem_input[r][len(problem_input[r])-orig_width]
                new_value = value + 1 if value < 9 else 1
                problem_input[r].append(new_value)

    for _ in range(orig_height):
        for _ in range(4):
            new_row = []
            for c in range(len(problem_input[0])):
                value = problem_input[len(problem_input)-orig_height][c]
                new_value = value + 1 if value < 9 else 1
                new_row.append(new_value)

            problem_input.append(new_row)

    # for row in problem_input:
    #     print(row)

    return solve1(problem_input)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day15.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    problem_input_2d_ints = [[int(num) for num in line]
                             for line in problem_input]

    # print(solve1(problem_input_2d_ints))
    print(solve2(problem_input_2d_ints))  # 2973 too high
