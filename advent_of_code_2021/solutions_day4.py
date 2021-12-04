from typing import Dict, List, Match, Optional, Tuple
from utils import parse_input


def solve1(problem_input: List[str]) -> int:
    # tables is a list of 2-D lists representing bingo tables
    # bingo_numbers is a list of numbers that will be called out for bingo
    bingo_numbers, tables = parse_bn_and_tables(problem_input)

    # dict of each number to a list of tuples representing (table index, row, col)
    number_dict = generate_number_dict(tables)

    for number in bingo_numbers:
        for match in number_dict.get(number, []):
            table = tables[match[0]]
            row = match[1]
            col = match[2]

            table[row][col] = "-" + table[row][col]

            solution = check_winner(table, row, col)

            if solution:
                return solution


def parse_bn_and_tables(problem_input: List[str]) -> Tuple[List[str], List[List[List[str]]]]:
    bingo_numbers = problem_input[0].split(",")
    tables = []
    curr_table = []

    for line in problem_input[2:]:
        if line == "":
            tables.append(curr_table)
            curr_table = []
        else:
            row = []
            numbers = line.split(" ")

            for number in numbers:
                if number.isdigit():
                    row.append(number)

            curr_table.append(row)

    if len(curr_table) > 0:
        tables.append(curr_table)

    return bingo_numbers, tables


def generate_number_dict(tables: List[List[List[str]]]) -> Dict[str, List[Tuple[int, int, int]]]:
    number_dict = {}

    for table_number, table in enumerate(tables):
        for row in range(len(table)):
            for col in range(len(table[0])):
                number_string = table[row][col]
                new_value = (table_number, row, col)

                if number_string not in number_dict:
                    number_dict[number_string] = []

                number_dict[number_string].append(new_value)

    return number_dict


def check_winner(table: List[List[str]], row: int, col: int) -> Optional[int]:
    def get_unmarked_sum(table: List[List[str]]) -> int:
        unmarked_sum = 0
        for row in table:
            for number_str in row:
                number = int(number_str)
                if number > 0:
                    unmarked_sum += number

        return unmarked_sum
    # Check row
    row_bingo = True
    for number_str in table[row]:
        if "-" not in number_str:
            row_bingo = False
            break

    if row_bingo:
        return get_unmarked_sum(table) * abs(int(table[row][col]))

    # Check col
    col_bingo = True
    for row_index in range(len(table)):
        number_str = table[row_index][col]
        if "-" not in number_str:
            col_bingo = False
            break

    if col_bingo:
        return get_unmarked_sum(table) * abs(int(table[row][col]))

    return None


def solve2(problem_input: List[str]) -> int:
    # tables is a list of 2-D lists representing bingo tables
    # bingo_numbers is a list of numbers that will be called out for bingo
    bingo_numbers, tables = parse_bn_and_tables(problem_input)

    # dict of each number to a list of tuples representing (table index, row, col)
    number_dict = generate_number_dict(tables)
    winning_tables = set(range(len(tables)))

    for number in bingo_numbers:
        for match in number_dict.get(number, []):
            table = tables[match[0]]
            row = match[1]
            col = match[2]

            table[row][col] = "-" + table[row][col]

            solution = check_winner(
                table, row, col) if match[0] in winning_tables else None

            if solution:
                winning_tables.remove(match[0])
                if len(winning_tables) == 0:
                    return solution


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day4.txt")

    # problem_input = [int(num) for num in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
