from typing import Dict, List
from utils import parse_input


def create_timer_dict(fish: List[int]) -> Dict[int, int]:
    timer_dict: Dict[int, int] = {}

    for fish_timer in fish:
        timer_dict[fish_timer] = timer_dict.get(fish_timer, 0) + 1

    return timer_dict


def solve1(problem_input: List[int], days: int = 80) -> int:
    # Approach use hash map to store number of fish timers at a specific day
    # Each day "subtract 1" from each timer by moving fish values from timer_day
    # key to timer_day - 1
    # Sum all values at end
    fish_timer_dict = create_timer_dict(problem_input)
    max_timer_day = 8
    reset_timer_day = 6

    for _ in range(days):
        tmp_next = fish_timer_dict.get(max_timer_day, 0)
        tmp_curr = -1
        for day in reversed(range(max_timer_day)):  # start with day 7
            if day == 0:
                fish_timer_dict[max_timer_day] = fish_timer_dict.get(0, 0)
                fish_timer_dict[reset_timer_day] = fish_timer_dict.get(
                    reset_timer_day, 0) + fish_timer_dict.get(0, 0)
                fish_timer_dict[0] = tmp_next
            else:
                tmp_curr = fish_timer_dict.get(day, 0)
                fish_timer_dict[day] = tmp_next
                tmp_next = tmp_curr

    return sum(fish_timer_dict.values())


def solve2(problem_input: List[int]) -> int:
    return solve1(problem_input, 256)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day6.txt")

    problem_input_int = [int(num) for num in problem_input[0].split(",")]

    print(solve1(problem_input_int))
    print(solve2(problem_input_int))
