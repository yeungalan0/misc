from typing import List, Optional, Set, Tuple
from utils import parse_input


def solve1(problem_input: List[str], include_diagonals: bool = False) -> Optional[int]:
    # Approach:
    # Generate sets of coordinates given each pair
    # Have a global set, where if a duplicate is found, add it to dupes
    # Return dupes count
    unique: Set[Tuple[int, int]] = set()
    duplicates: Set[Tuple[int, int]] = set()

    for line in problem_input:
        points = generate_points(line, include_diagonals)
        for point in points:
            if point in unique:
                duplicates.add(point)
            else:
                unique.add(point)

    return len(duplicates)


def generate_points(points_str: str, include_diagonals: bool = False) -> List[Tuple[int, int]]:
    first: Tuple[int, int]
    second: Tuple[int, int]

    first, second = [tuple(map(int, point_str.split(",")))    # type: ignore
                     for point_str in points_str.split(" -> ")]

    if first[0] != second[0] and first[1] != second[1] and not include_diagonals:
        return []

    # What to add to turn first into second point
    first_row_add = 0 if first[0] == second[0] else 1 if second[0] > first[0] else -1
    first_col_add = 0 if first[1] == second[1] else 1 if second[1] > first[1] else -1

    points: List[Tuple[int, int]] = [first]
    current = first

    while current != second:
        current = (current[0] + first_row_add,
                   current[1] + first_col_add)
        points.append(current)

    return points


def solve2(problem_input: List[str]) -> Optional[int]:
    return solve1(problem_input, True)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day5.txt")
    print(solve1(problem_input))
    print(solve2(problem_input))
