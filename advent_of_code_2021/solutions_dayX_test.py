from solutions_dayX import solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [],
            None
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [],
            None
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
