from typing import List, Optional, Set, Tuple
from utils import parse_input


def solve1(problem_input: List[str]) -> Optional[int]:
    # Approach:
    # Generate sets of coordinates given each pair
    # Have a global set, where if a duplicate is found, add it to dupes
    # Return dupes count
    unique: Set[Tuple[int, int]] = set()
    duplicates: Set[Tuple[int, int]] = set()

    for line in problem_input:
        points = generate_points(line)
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

    points: List[Tuple[int, int]] = []

    if first[0] != second[0] and first[1] != second[1]:
        # Order them so first is smaller than second
        if first[0] > second[0]:
            first, second = second, first

        current: Tuple[int, int] = first
        add = 1 if second[1] > first[1] else -1

        while current != second:
            points.append(current)
            current = (current[0] + 1, current[1] + add)

        points.append(second)
    elif first[0] != second[0]:
        col = first[1]

        # Order them so first is smaller than second
        if first[0] > second[0]:
            first, second = second, first

        for number in range(first[0], second[0]+1):
            points.append((number, col))
    else:
        row = first[0]

        # Order them so first is smaller than second
        if first[1] > second[1]:
            first, second = second, first

        for number in range(first[1], second[1]+1):
            points.append((row, number))

    return points


def solve2(problem_input: List[str]) -> Optional[int]:
    unique: Set[Tuple[int, int]] = set()
    duplicates: Set[Tuple[int, int]] = set()

    for line in problem_input:
        points = generate_points(line, True)
        for point in points:
            if point in unique:
                duplicates.add(point)
            else:
                unique.add(point)

    return len(duplicates)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day5.txt")

    # problem_input = [int(num) for num in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
