import sys
import math
import string

# Actual solution logic should go here
def solve(case):
    velocity, distance = case.split(":")
    # if float(velocity) == 0:
    #     print("SAFE")
    #     return

    try:
        time = float(distance) / float(velocity)

        if time <= 1:
            print("SWERVE")
        elif time <= 5:
            print("BRAKE")
        else:
            print("SAFE")
    except ZeroDivisionError as e:
        print("gtrgrr")


# Logic to read to input and pass to our solving function
def main():
    cases = int(sys.stdin.readline().rstrip())
    for case_num in range(cases):
        case = sys.stdin.readline().rstrip()
        solve(case)

if __name__ == "__main__":
    main()