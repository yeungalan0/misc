from typing import List
from utils import parse_input


def solve1(problem_input: List[str]) -> int:
    # Array of counts for each binary digit
    places_count = [[0, 0] for x in problem_input[0]]

    for binary_number_string in problem_input:
        for i, digit_string in enumerate(binary_number_string):
            if digit_string == "0":
                places_count[i][0] += 1
            else:
                places_count[i][1] += 1

    gamma = 0
    epsilon = 0

    # Let's start with least significant places
    places_count.reverse()

    for i, binary_count in enumerate(places_count):
        if binary_count[1] > binary_count[0]:
            gamma |= 1 << i
        else:
            epsilon |= 1 << i

    return gamma * epsilon


def solve2(problem_input: List[str]) -> int:
    return get_gas_rating(True, problem_input) * get_gas_rating(False, problem_input)


def get_gas_rating(oxygen: bool, problem_input: List[str]):
    relevant_rating = "1" if oxygen else "0"
    other_rating = "0" if relevant_rating == "1" else "1"
    gas_list = problem_input
    index = 0
    while len(gas_list) > 1:
        count_1 = 0

        for binary_str in gas_list:
            if binary_str[index] == "1":
                count_1 += 1

        count_0 = len(gas_list) - count_1

        if count_0 > count_1:
            gas_list = [
                binary_str for binary_str in gas_list if binary_str[index] == other_rating]
        else:
            gas_list = [
                binary_str for binary_str in gas_list if binary_str[index] == relevant_rating]

        index += 1

    return int(gas_list[0], 2)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day3.txt")

    # problem_input = [int(num) for num in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
