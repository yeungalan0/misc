from solutions_day5 import generate_points, solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "0,9 -> 5,9",
                "8,0 -> 0,8",
                "9,4 -> 3,4",
                "2,2 -> 2,1",
                "7,0 -> 7,4",
                "6,4 -> 2,0",
                "0,9 -> 2,9",
                "3,4 -> 1,4",
                "0,0 -> 8,8",
                "5,5 -> 8,2",
            ],
            5
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_generate_points():
    test_cases = [
        utils.TestCase(  # row works
            ["0,9 -> 5,9"],
            [(0, 9), (1, 9), (2, 9), (3, 9), (4, 9), (5, 9)]
        ),
        utils.TestCase(  # Ignore vertical lines...
            ["8,0 -> 0,8"],
            []
        ),
        utils.TestCase(  # col works
            ["4,9 -> 4,3"],
            [(4, 3), (4, 4), (4, 5), (4, 6), (4, 7), (4, 8), (4, 9)]
        )
    ]

    for test_case in test_cases:
        actual_output = generate_points(test_case.case[0])

        assert test_case.expected_output == actual_output


# TODO: Need to adjust testing framework, instead of a list probably need an Any or an OR
def test_generate_points_diagonals():
    test_cases = [
        utils.TestCase(  # diagonals up works
            ["9,7 -> 7,9"],
            [(7, 9), (8, 8), (9, 7)]
        ),
        utils.TestCase(  # diagonals up works
            ["0,0 -> 3,3"],
            [(0, 0), (1, 1), (2, 2), (3, 3)]
        ),
        utils.TestCase(  # diagonals up works
            ["3,0 -> 0,3"],
            [(0, 3), (1, 2), (2, 1), (3, 0)]
        )
    ]

    for test_case in test_cases:
        actual_output = generate_points(test_case.case[0], True)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                "0,9 -> 5,9",
                "8,0 -> 0,8",
                "9,4 -> 3,4",
                "2,2 -> 2,1",
                "7,0 -> 7,4",
                "6,4 -> 2,0",
                "0,9 -> 2,9",
                "3,4 -> 1,4",
                "0,0 -> 8,8",
                "5,5 -> 8,2",
            ],
            12
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
