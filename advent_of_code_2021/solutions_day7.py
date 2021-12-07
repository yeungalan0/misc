from typing import List
from utils import parse_input


def solve1(problem_input: List[int]) -> int:
    unique_hor_positions = set(problem_input)
    least_cost = 10000000000000000000000000000000000000000000

    curr_cost = 0
    for hor_position in unique_hor_positions:
        for crab_hor_position in problem_input:
            diff = abs(crab_hor_position - hor_position)
            curr_cost += diff

        if curr_cost < least_cost:
            least_cost = curr_cost

        curr_cost = 0

    return least_cost


def solve2(problem_input: List[int]) -> int:
    unique_hor_positions = set(problem_input)
    least_cost = 10000000000000000000000000000000000000000000

    curr_cost = 0
    for hor_position in range(max(unique_hor_positions)):
        for crab_hor_position in problem_input:
            x = abs(crab_hor_position - hor_position)
            diff = x * (x + 1) // 2
            curr_cost += diff

        if curr_cost < least_cost:
            least_cost = curr_cost

        curr_cost = 0

    return least_cost


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day7.txt")

    problem_input_int = [int(num) for num in problem_input[0].split(",")]

    print(solve1(problem_input_int))
    print(solve2(problem_input_int))
