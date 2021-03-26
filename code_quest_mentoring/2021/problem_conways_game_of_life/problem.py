import sys
import math
import string

# Actual solution logic should go here
ALIVE = 1


def solve(iterations, grid):
    print(f"Iterations: {iterations}")
    print(f"Grid: {len(grid)}")
    int_grid = convert_to_int_grid(grid)
    print(count_adjacent_neighbors(3, 5, int_grid))


def convert_to_int_grid(grid):
    int_grid = []
    for line in grid:
        int_grid.append([int(character) for character in line])
    return int_grid


def count_adjacent_neighbors(row, col, grid):
    neighbors = 0
    min_row = max(0, row - 1)
    max_row = min(10, row + 2)
    min_col = max(0, col - 1)
    max_col = min(10, col + 2)

    for check_row in range(min_row, max_row):
        for check_col in range(min_col, max_col):
            if check_row == row and check_col == col:
                continue
            if grid[check_row][check_col] == ALIVE:
                neighbors += 1
            print(check_row, check_col)

    return neighbors


# Logic to read to input and pass to our solving function
def main():
    cases = int(sys.stdin.readline().rstrip())
    for _ in range(cases):
        iterations = int(sys.stdin.readline().rstrip())
        grid = []

        for _ in range(10):
            grid.append(sys.stdin.readline().rstrip())

        solve(iterations, grid)


if __name__ == "__main__":
    main()
