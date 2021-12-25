from typing import Callable, List
from utils import parse_input
import math


# Solution pretty much entirely comes from the below amazing redit pos/code that solves this problem way better than I could currently.
https: // topaz.github.io/paste/  # XQAAAQAMBwAAAAAAAAAxHMAC0B2Tuh+Zc7rcnUmLYPQvNsm4Tg4i3Zjx2q0nDfSbeEz3deAe3DypB+rhSDyH0m+Kv7tpnOVvFjs43k2hXGyyF/u5PnCx9/MR2UouZnZsl1Fvf5kjeG4kH+D5e9tT6UwvsL2h1WJ3G/6unBKk7qJxlmqmBAfpC3krCj8l88/MjzqjXQVtmwwYN6Kq5/OKQvoJMdXZKdsKPZOnVxHMMD2qB1Ck2NZIq2r+ZHQIWBjnj/2BVUA6u4c67BjiSlLMNQXagsy9+wllPDINTujXIkgLo/3hmApsn7/E04vXucfWYdhDWh8f+Ovkp91w6Wj6lmS3kcmoVgLcBdflDoi7EWnA5bc3cOjSaJ9zMEIcu2BRDtehR8Mz0T65fxSkMtfp9sjkQ9/t0kMcIJMxAdy9sFKFLIO3VM0X24f51aAsfnkUvmC/9c1J8bCsx6gs6qtFEiSJ94aZ7nRfy1m7YG+eR4BycNihtQxXyAHL0mw4bEMstvaCGrZCxt285XbpkWnNYKoF5nTZ7Qfh/JBT65019W7DNMIQJAEBkaxsIBCdsV+xLBSaR+ZbkN54HB13Z69MROYoZwwkazNGSuAScmUFpZs96Z6BfwWITtDLwKeAE0CBIedtWAlc0RJRkHh9z+rIgnanFKU/FiVUyoMAB40xnv6sUty11jWCyRTqDX3H8lNlEvukHsy1JCUiBzGontCjV6WzcDdN2wh4n83SclzuWUprbaeZgwxzPRHPzgTa3uh81QGLAD4TMw8gYZyBfAMmz6NI/P+4YpGw


def convert_to_binary(hex_str: str) -> str:
    # Clever trick to convert to hex from https://stackoverflow.com/a/17157819/5910564
    return bin(int("1"+hex_str, 16))[3:]


def solve1_helper(start_index: int, bin_str: str):
    i = start_index
    total_versions = int(bin_str[i:i+3], 2)
    type_id = int(bin_str[i+3:i+6], 2)
    i += 6

    if type_id == 4:
        while True:
            i += 5
            if bin_str[i-5] == "0":
                break
    else:
        if bin_str[i] == "0":
            # Figure out end based on 15 bit number
            end_i = i + 16 + int(bin_str[i+1:i+16], 2)
            i += 16

            while i < end_i:
                i, versions = solve1_helper(i, bin_str)
                total_versions += versions
        else:
            sub_packets = int(bin_str[i+1:i+12], 2)
            i += 12

            for _ in range(sub_packets):
                i, versions = solve1_helper(i, bin_str)
                total_versions += versions

    return i, total_versions


def solve2_helper(start_index: int, bin_str: str):
    i = start_index
    type_id = int(bin_str[i+3:i+6], 2)
    i += 6

    values: List[int] = []

    ops: List[Callable[[List[int]], int]] = [
        sum,
        math.prod,
        min,
        max,
        lambda value_list: value_list[0],  # literal value
        lambda value_list: 1 if value_list[0] > value_list[1] else 0,
        lambda value_list: 1 if value_list[0] < value_list[1] else 0,
        lambda value_list: 1 if value_list[0] == value_list[1] else 0
    ]

    if type_id == 4:
        bin_value_str = ""
        while True:
            bin_value_str += bin_str[i+1:i+5]
            i += 5
            if bin_str[i-5] == "0":
                break

        values.append(int(bin_value_str, 2))
    else:
        if bin_str[i] == "0":
            # Figure out end based on 15 bit number
            end_i = i + 16 + int(bin_str[i+1:i+16], 2)
            i += 16

            while i < end_i:
                i, value = solve2_helper(i, bin_str)
                values.append(value)
        else:
            sub_packets = int(bin_str[i+1:i+12], 2)
            i += 12

            for _ in range(sub_packets):
                i, value = solve2_helper(i, bin_str)
                values.append(value)

    return i, ops[type_id](values)


def solve1(hex_str: str) -> int:
    bin_str = convert_to_binary(hex_str)

    return solve1_helper(0, bin_str)[1]


def solve2(hex_str: str) -> int:
    bin_str = convert_to_binary(hex_str)

    return solve2_helper(0, bin_str)[1]


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day16.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input[0]))
    print(solve2(problem_input[0]))  # Too high: 5390819598101
