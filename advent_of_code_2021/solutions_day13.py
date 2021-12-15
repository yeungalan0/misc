from typing import List, Set, Tuple
from utils import parse_input


def parse_points_and_folds(problem_input: List[str]) -> Tuple[Set[Tuple[int, int]], List[Tuple[str, int]]]:
    points: Set[Tuple[int, int]] = set()
    folds: List[Tuple[str, int]] = []
    process_points = True

    for line in problem_input:
        if line == "":
            process_points = False
            continue

        if process_points:
            x, y = line.split(",")
            points.add((int(x), int(y)))
        else:
            _, _, fold_str = line.split(" ")
            axis, magnitude = fold_str.split("=")
            folds.append((axis, int(magnitude)))

    return (points, folds)


def get_fold_points(points: Set[Tuple[int, int]], fold: Tuple[str, int]) -> Set[Tuple[int, int]]:
    x_or_y = 0 if fold[0] == "x" else 1
    fold_distance = fold[1]

    new_points: Set[Tuple[int, int]] = set()

    for point in points:
        point_to_add = point
        if point[x_or_y] > fold_distance:
            orig_point = list(point)
            new_value = fold_distance * 2 - point[x_or_y]
            orig_point[x_or_y] = new_value

            point_to_add = tuple(orig_point)  # type: ignore

        new_points.add(point_to_add)

    return new_points


def solve1(problem_input: List[str]) -> int:
    points, folds = parse_points_and_folds(problem_input)

    for fold in folds[0:1]:
        points = get_fold_points(points, fold)

    return len(points)


def solve2(problem_input: List[str]):
    points, folds = parse_points_and_folds(problem_input)

    for fold in folds:
        points = get_fold_points(points, fold)

    width = next(x for x in reversed(folds) if x[0] == "x")[1]
    height = next(y for y in reversed(folds) if y[0] == "y")[1]

    for r in range(width):
        row = []
        for c in range(height):
            if (r, c) in points:
                row.append("#")
            else:
                row.append(".")
        print(row)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day13.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
