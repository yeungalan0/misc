from typing import List
from utils import parse_input


def solve1(problem_input: List[str]) -> int:
    horizontal = 0
    depth = 0

    for instruction in problem_input:
        direction, magnitude_str = instruction.split(" ")
        magnitude_int = int(magnitude_str)
        if direction == "forward":
            horizontal += magnitude_int
        elif direction == "up":
            depth -= magnitude_int
        else:
            depth += magnitude_int

    return horizontal * depth


def solve2(problem_input: List[str]) -> int:
    horizontal = 0
    depth = 0
    aim = 0

    for instruction in problem_input:
        direction, magnitude_str = instruction.split(" ")
        magnitude_int = int(magnitude_str)
        if direction == "forward":
            horizontal += magnitude_int
            depth += aim * magnitude_int
        elif direction == "up":
            aim -= magnitude_int
        else:
            aim += magnitude_int

    return horizontal * depth


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day2.txt")

    print(solve1(problem_input))
    print(solve2(problem_input))
