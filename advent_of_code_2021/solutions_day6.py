from typing import Dict, List
from utils import parse_input


# input: list of fish times before spawning
# returns: list where index is days left before spawning and value is count of fish


def create_count_list(fish_times: List[int], max_days: int) -> List[int]:
    """
    Parameters:
    argument1: list of fish times before spawning

    Returns:
    list where index is days left before spawning and value is count of fish
    """
    # want to include the last day itself
    fish_count_list = [0] * (max_days + 1)

    for fish_time_left in fish_times:
        fish_count_list[fish_time_left] += 1

    return fish_count_list


def solve1(problem_input: List[int], days: int = 80) -> int:
    # Approach: use list to store number of fish at a specific day in spawn cycle
    # where the index represents the days left in cycle and value is the count.
    # Each day "subtract 1" from each timer by moving fish values from index to index -1
    # by chopping list and adding appropriate number of spawned fish/reset fish
    # Sum all values at end
    max_timer_day = 8
    reset_timer_day = 6
    fish_timer_list = create_count_list(problem_input, max_timer_day)

    for _ in range(days):
        fish_count_at_0 = fish_timer_list[0]
        fish_timer_list = fish_timer_list[1:]
        fish_timer_list.append(fish_count_at_0)
        fish_timer_list[reset_timer_day] += fish_count_at_0

    return sum(fish_timer_list)


def solve2(problem_input: List[int]) -> int:
    return solve1(problem_input, 256)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day6.txt")

    problem_input_int = [int(num) for num in problem_input[0].split(",")]

    print(solve1(problem_input_int))
    print(solve2(problem_input_int))
