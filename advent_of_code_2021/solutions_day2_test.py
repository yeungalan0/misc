from solutions_day2 import solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "forward 5",
                "down 5",
                "forward 8",
                "up 3",
                "down 8",
                "forward 2",
            ],
            150
        ),
        utils.TestCase(  # Only decreasing depth
            [
                "forward 5",
                "up 5",
                "forward 8",
                "up 3",
                "up 8",
                "forward 2",
            ],
            -240
        ),
        utils.TestCase(  # Only decreasing depth
            [
                "forward 5",
                "down 5",
                "forward 8",
                "down 3",
                "down 8",
                "forward 2",
            ],
            240
        ),
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "forward 5",
                "down 5",
                "forward 8",
                "up 3",
                "down 8",
                "forward 2",
            ],
            900
        )
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
