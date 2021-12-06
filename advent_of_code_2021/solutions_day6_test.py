from solutions_day6 import create_timer_dict, solve1, solve2
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


def test_create_timer_dict():
    test_cases = [
        utils.TestCase(  # Given test case
            [3, 4, 3, 1, 2],
            {1: 1, 2: 1, 3: 2, 4: 1}
        )
    ]

    for test_case in test_cases:
        actual_output = create_timer_dict(test_case.case)

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
