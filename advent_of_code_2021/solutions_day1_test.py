from solutions_day1 import solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [199, 200, 208, 210, 200, 207, 240, 269, 260, 263],
            7
        ),
        utils.TestCase(  # Only increasing results
            [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
            9
        ),
        utils.TestCase(  # Only decreasing results
            [10, 9, 8, 7, 6, 5, 4, 3, 2, 1],
            0
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [199, 200, 208, 210, 200, 207, 240, 269, 260, 263],
            5
        ),
        utils.TestCase(  # Not enough numbers
            [1, 2],
            0
        ),
        utils.TestCase(  # Only increasing results
            [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
            7
        ),
        utils.TestCase(  # Only decreasing results
            [10, 9, 8, 7, 6, 5, 4, 3, 2, 1],
            0
        )
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
