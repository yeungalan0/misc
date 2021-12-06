from solutions_day6 import create_count_list, solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [3, 4, 3, 1, 2],
            5934
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_create_count_list():
    test_cases = [
        utils.TestCase(  # Given test case
            ([3, 4, 3, 1, 2], 8),
            [0, 1, 1, 2, 1, 0, 0, 0, 0]
        )
    ]

    for test_case in test_cases:
        actual_output = create_count_list(
            fish_times=test_case.case[0], max_days=test_case.case[1])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given test case
            [3, 4, 3, 1, 2],
            26984457539
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
