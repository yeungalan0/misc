import sys
import math
import string

# Actual solution logic should go here
def solve(case):
    print(case)

# Logic to read to input and pass to our solving function
def main():
    cases = int(sys.stdin.readline().rstrip())
    for case_num in range(cases):
        case = sys.stdin.readline().rstrip()
        solve(case)

if __name__ == "__main__":
    main()