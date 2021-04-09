import sys
import math
import string

# Actual solution logic should go here
ALIVE = 1
DEAD = 0


def solve(iterations, grid):
    int_grid = convert_to_int_grid(grid)
    evolved_grid = int_grid
    for _ in range(iterations):
        evolved_grid = evolve(evolved_grid)

    print_grid(evolved_grid)


def evolve(grid):
    new_grid = copy_grid(grid)

    for row in range(len(grid)):
        for col in range(len(grid[row])):
            living_neighbors = count_adjacent_neighbors(row, col, grid)
            if grid[row][col] == ALIVE and (
                    living_neighbors < 2 or
                    living_neighbors > 3
            ):
                new_grid[row][col] = 0
            if grid[row][col] == DEAD and living_neighbors == 3:
                new_grid[row][col] = 1

    return new_grid


def print_grid(grid):
    for row in grid:
        print("".join([str(x) for x in row]))


def convert_to_int_grid(grid):
    int_grid = []
    for line in grid:
        int_grid.append([int(character) for character in line])
    return int_grid


def count_adjacent_neighbors(row, col, grid):
    neighbors = 0
    min_row = max(0, row - 1)
    max_row = min(10, row + 2)  # TODO
    min_col = max(0, col - 1)
    max_col = min(10, col + 2)

    for check_row in range(min_row, max_row):
        for check_col in range(min_col, max_col):
            if check_row == row and check_col == col:
                continue
            if grid[check_row][check_col] == ALIVE:
                neighbors += 1

    return neighbors


def copy_grid(grid):
    grid_clone = []

    for row in grid:
        grid_clone.append(row.copy())

    return grid_clone


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
