from typing import List, Set, Tuple
from utils import parse_input
import math


def find_low_point_coordinates(heightmap: List[List[int]]) -> List[Tuple[int, int]]:
    low_point_coords = []

    # Top, right, bottom, left
    adjacent_options = [(-1, 0), (0, 1), (1, 0), (0, -1)]
    height = len(heightmap)
    width = len(heightmap[0])

    for r in range(height):
        for c in range(len(heightmap[0])):
            adjacent_cells = []
            for adjacent_option in adjacent_options:
                adjacent_cell = (
                    r + adjacent_option[0], c + adjacent_option[1])
                is_valid = adjacent_cell[0] > - \
                    1 and adjacent_cell[0] < height and adjacent_cell[1] > - \
                    1 and adjacent_cell[1] < width
                if is_valid:
                    adjacent_cells.append(adjacent_cell)

            if all(heightmap[r][c] < heightmap[adj_coord[0]][adj_coord[1]] for adj_coord in adjacent_cells):
                low_point_coords.append((r, c))

    return low_point_coords


def get_basin_size(low_point_coord: Tuple[int, int], heightmap: List[List[int]]) -> int:
    basin_coords: Set[Tuple[int, int]] = set()
    seen: Set[Tuple[int, int]] = set()
    queue = [low_point_coord]
    adjacent_options = [(-1, 0), (0, 1), (1, 0), (0, -1)]
    height = len(heightmap)
    width = len(heightmap[0])

    while len(queue) > 0:
        coord = queue.pop(0)
        r, c = coord
        if heightmap[r][c] < 9 and coord not in seen:
            basin_coords.add(coord)
            for adjacent_option in adjacent_options:
                adjacent_cell = (
                    r + adjacent_option[0], c + adjacent_option[1])
                is_valid = adjacent_cell[0] > - \
                    1 and adjacent_cell[0] < height and adjacent_cell[1] > - \
                    1 and adjacent_cell[1] < width and adjacent_cell not in seen
                if is_valid:
                    queue.append(adjacent_cell)

        seen.add(coord)

    return len(basin_coords)


def solve1(problem_input: List[List[int]]) -> int:
    low_point_coords = find_low_point_coordinates(problem_input)

    risk_values = [problem_input[r][c] + 1 for r, c in low_point_coords]
    return sum(risk_values)


def solve2(problem_input: List[List[int]]) -> int:
    basin_sizes = []

    low_point_coords = find_low_point_coordinates(problem_input)

    for low_point_coord in low_point_coords:
        basin_size = get_basin_size(low_point_coord, problem_input)
        basin_sizes.append(basin_size)

    basin_sizes_sorted = sorted(basin_sizes)

    return math.prod(basin_sizes_sorted[-3:])


if __name__ == "__main__":
    heightmap = parse_input("inputs/input_day9.txt")

    # problem_input_int = [int(num) for num in problem_input[0].split()]

    problem_input_2d_ints = [[int(num) for num in line]
                             for line in heightmap]

    print(solve1(problem_input_2d_ints))
    print(solve2(problem_input_2d_ints))
