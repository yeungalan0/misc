from typing import Coroutine, List, Set, Tuple
from utils import parse_input
import heapq
import math


def djikstras_min_path(grid: List[List[int]]) -> int:
    bests = [[math.inf] * len(row) for row in grid]
    bests[0][0] = 0
    queue: List[Tuple[int, int, int]] = []
    heapq.heappush(queue, (0, 0, 0))
    while True:
        _, x0, y0 = heapq.heappop(queue)
        c = bests[y0][x0]
        if y0 == len(grid) - 1 and x0 == len(grid[y0]) - 1:
            return c
        for x1, y0 in ((x0 - 1, y0), (x0, y0 - 1), (x0, y0 + 1), (x0 + 1, y0)):
            if y0 not in range(len(grid)) or x1 not in range(len(grid[y0])):
                continue
            d = c + grid[y0][x1]
            if d < bests[y0][x1]:
                bests[y0][x1] = d
                heapq.heappush(queue, (d, x1, y0))


def solve1(problem_input: List[List[int]]) -> int:
    return djikstras_min_path(problem_input)


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

    print(solve1(problem_input_2d_ints))
    print(solve2(problem_input_2d_ints))  # 2967, 2973 too high
