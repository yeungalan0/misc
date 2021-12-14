from typing import List, Set, Tuple
from utils import parse_input
from copy import deepcopy


def increase_energy(energy_grid: List[List[int]]):
    for row in energy_grid:
        for i in range(len(row)):
            row[i] += 1

    return energy_grid


def flash_octopuses(energy_grid: List[List[int]]) -> Tuple[List[List[int]], int]:
    adj_options = [(-1, -1), (-1, 0), (-1, 1), (0, -1),
                   (0, 1), (1, -1), (1, 0), (1, 1)]
    flashed_coordinates: Set[Tuple[int, int]] = set()
    queue = []
    height = len(energy_grid)
    width = len(energy_grid[0])

    # print_grid(energy_grid, 0)

    for r in range(height):
        for c in range(width):
            if energy_grid[r][c] > 9:
                queue.append((r, c))

    while len(queue) > 0:
        flashing_octopus_coord = queue.pop(0)
        # Add 1 to each adjacent neighbor
        for adj_option in adj_options:
            adj_coord = (
                flashing_octopus_coord[0] + adj_option[0], flashing_octopus_coord[1] + adj_option[1])
            adj_r, adj_c = adj_coord

            is_valid = 0 <= adj_r < height and 0 <= adj_c < width and adj_coord not in flashed_coordinates and adj_coord not in queue

            if is_valid:
                energy_grid[adj_r][adj_c] += 1

                if energy_grid[adj_r][adj_c] > 9:
                    queue.append(adj_coord)

        r, c = flashing_octopus_coord
        flashed_coordinates.add(flashing_octopus_coord)
        energy_grid[r][c] = 0

    # print_grid(energy_grid, 1)
    return (energy_grid, len(flashed_coordinates))


def print_grid(energy_grid: List[List[int]], iteration: int = -1):
    print(f"Grid iteration: {iteration}")
    for row in energy_grid:
        print(row)


def solve1(problem_input: List[List[int]], steps: int = 100) -> Tuple[List[List[int]], int]:
    energy_grid = deepcopy(problem_input)
    flash_count = 0
    for _ in range(steps):
        # Add 1 to all octopuses
        energy_grid = increase_energy(energy_grid)
        energy_grid, new_flashes = flash_octopuses(energy_grid)
        flash_count += new_flashes

        # print_grid(energy_grid, _)

    return (energy_grid, flash_count)


def solve2(problem_input: List[List[int]]) -> int:
    energy_grid = deepcopy(problem_input)
    octopus_count = len(energy_grid) * len(energy_grid[0])
    step = 0

    while True:
        step += 1
        # Add 1 to all octopuses
        energy_grid = increase_energy(energy_grid)
        energy_grid, new_flashes = flash_octopuses(energy_grid)

        if new_flashes == octopus_count:
            break

        # print_grid(energy_grid, _)

    return step


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day11.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    problem_input_2d_ints = [[int(num) for num in line]
                             for line in problem_input]

    # Stupidity 1: python allows negative indexing, so my try catch was dumb
    # Stupidity 2: solution 1 in place changes problem_input_2d_ints, destroying my answer to 2
    print(solve1(problem_input_2d_ints))
    print(solve2(problem_input_2d_ints))  # 300 (too low)
