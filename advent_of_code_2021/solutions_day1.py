from typing import List
from utils import parse_input


def solve1(problem_input: List[int]) -> int:
    previous = problem_input[0]
    count = 0

    for elem in problem_input[1:]:
        if elem > previous:
            count += 1
        previous = elem

    return count


def solve2(problem_input: List[str]) -> int:
    first = sum(problem_input[:3])
    second = sum(problem_input[1:4])
    count = 1 if second > first else 0

    next_index = 4
    last_index = 0

    while next_index < len(problem_input):
        first -= problem_input[last_index]
        first += problem_input[next_index - 1]
        second -= problem_input[last_index + 1]
        second += problem_input[next_index]

        if second > first:
            count += 1

        next_index += 1
        last_index += 1

    return count


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day1.txt")

    problem_input_ints = [int(num) for num in problem_input]

    print(solve1(problem_input_ints))
    print(solve2(problem_input_ints))
