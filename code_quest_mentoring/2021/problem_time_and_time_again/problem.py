import sys
import math
import string
import re


# Actual solution logic should go here
def solve(case):
    hours = extract_time_units(r'(\d{1,2})h', case)
    minutes = extract_time_units(r'(\d{1,2})m', case)
    seconds = extract_time_units(r'(\d{1,2})s', case)

    print(f"{hours}:{minutes}:{seconds}")


def extract_time_units(regex_pattern, case):
    time_unit = "00"
    time_unit_search = re.search(regex_pattern, case)

    if time_unit_search:
        time_unit = time_unit_search.group(1)

    if len(time_unit) == 1:
        time_unit = "0" + time_unit

    return time_unit


# Logic to read to input and pass to our solving function
def main():
    cases = int(sys.stdin.readline().rstrip())
    for _ in range(cases):
        case = sys.stdin.readline().rstrip()
        solve(case)


if __name__ == "__main__":
    main()
