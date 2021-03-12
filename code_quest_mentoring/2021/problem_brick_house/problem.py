import sys
import math
import string

SHORT_BRICK_LENGTH = 1
LARGE_BRICK_LENGTH = 5

# Actual solution logic should go here


def solve(case):
    running_sum = 0
    short_num, large_num, target = map(lambda x: int(x), case.split(" "))
    if short_num + large_num * LARGE_BRICK_LENGTH < target:
        return False

    while large_num > 0 and running_sum < target:
        running_sum += LARGE_BRICK_LENGTH
        large_num -= 1

    if running_sum == target:
        return True
    elif running_sum > target:
        running_sum -= LARGE_BRICK_LENGTH

    while short_num > 0 and running_sum < target:
        running_sum += SHORT_BRICK_LENGTH
        short_num -= 1

    if running_sum == target:
        return True

    return False

# Logic to read to input and pass to our solving function


def main():
    cases = int(sys.stdin.readline().rstrip())
    for _ in range(cases):
        case = sys.stdin.readline().rstrip()
        if solve(case):
            print("true")
        else:
            print("false")


if __name__ == "__main__":
    main()
