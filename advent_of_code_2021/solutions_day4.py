from typing import Dict, List, Match, Optional, Tuple
from utils import parse_input
from itertools import chain

# Rewrote my solutions for understanding after being inspired by https://www.reddit.com/r/adventofcode/comments/r8i1lq/comment/hn7jm8x/?utm_source=share&utm_medium=web2x&context=3


def solve1(problem_input: List[str]) -> Optional[int]:
    # tables is a list of 2-D lists representing bingo tables
    # bingo_numbers is a list of numbers that will be called out for bingo
    bingo_numbers, tables = parse_bn_and_tables(problem_input)

    called = []

    for number in bingo_numbers:
        called.append(number)
        for table in tables:
            rows_and_cols = chain(table, zip(*table))

            if any(set(row_or_col) <= set(called) for row_or_col in rows_and_cols):
                unmarked_numbers = [
                    n for row in table for n in row if n not in called]
                return sum(unmarked_numbers) * number

    return None


def parse_bn_and_tables(problem_input: List[str]) -> Tuple[List[int], List[List[List[int]]]]:
    bingo_numbers = list(map(int, problem_input[0].split(",")))
    tables: List[List[List[int]]] = []
    curr_table: List[List[int]] = []

    for line in problem_input[2:]:
        if line == "":
            tables.append(curr_table)
            curr_table = []
        else:
            numbers = line.split(" ")
            row = [int(number_str)
                   for number_str in numbers if number_str.isdigit()]

            curr_table.append(row)

    if len(curr_table) > 0:
        tables.append(curr_table)

    return bingo_numbers, tables


def solve2(problem_input: List[str]) -> Optional[int]:
    # tables is a list of 2-D lists representing bingo tables
    # bingo_numbers is a list of numbers that will be called out for bingo
    bingo_numbers, tables = parse_bn_and_tables(problem_input)

    called = bingo_numbers.copy()

    while len(called) > 0:
        last_called = called.pop()
        for table in tables:
            rows_and_cols = chain(table, zip(*table))

            if not any(set(row_or_col) <= set(called) for row_or_col in rows_and_cols):
                # Remark the last_called number once we've found the last winning table
                called.append(last_called)
                unmarked_numbers = [
                    n for row in table for n in row if n not in called]
                return sum(unmarked_numbers) * last_called

    return None


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day4.txt")

    # problem_input = [int(num) for num in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
