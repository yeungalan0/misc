from solutions_day17 import lands_in_target, solve2
import utils


def test_lands_in_target():
    test_cases = [
        utils.TestCase(  # Given test case
            (7, 2, 20, 30, -10, -5),
            True
        ),
        utils.TestCase(  # Given test case
            (9, 0, 20, 30, -10, -5),
            True
        ),
        utils.TestCase(  # Given test case
            (29, -10, 20, 30, -10, -5),
            True
        ),
        utils.TestCase(  # Given test case
            (7, 5, 20, 30, -10, -5),
            True
        ),
        utils.TestCase(  # Given test case
            (7, 12, 20, 30, -10, -5),
            False
        )
    ]

    for test_case in test_cases:
        actual_output = lands_in_target(
            test_case.case[0], test_case.case[1], test_case.case[2], test_case.case[3], test_case.case[4], test_case.case[5])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            ["target area: x=20..30, y=-10..-5"],
            112
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
