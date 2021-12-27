from typing import List
from utils import parse_input


def lands_in_target(x: int, y: int, min_x: int, max_x: int, min_y: int, max_y: int) -> bool:
    x_velocity = x
    y_velocity = y
    while True:
        if x > max_x or y < min_y:
            return False
        elif min_x <= x <= max_x and min_y <= y <= max_y:
            return True
        else:
            if x_velocity > 0:
                x_velocity -= 1
                x += x_velocity
            elif x_velocity < 0:
                x_velocity += 1
                x -= x_velocity
            y_velocity -= 1
            y += y_velocity


def solve2(problem_input: List[str]) -> int:
    x_range, y_range = problem_input[0][13:].split(", ")

    min_x, max_x = map(int, x_range[2:].split(".."))
    min_y, max_y = map(int, y_range[2:].split(".."))

    print(min_x, max_x, min_y, max_y)

    count = 0

    for x in range(0, max_x+1):
        for y in range(min_y-1, abs(min_y-1)):
            if lands_in_target(x, y, min_x, max_x, min_y, max_y):
                count += 1

    return count


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day17.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    # print(solve1(problem_input))
    print(solve2(problem_input))
