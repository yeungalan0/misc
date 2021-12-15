from solutions_day13 import get_fold_points, parse_points_and_folds, solve1, solve2
import utils


def test_parse_points_and_folds():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "8,4",
                "1,10",
                "2,14",
                "",
                "fold along y=7",
                "fold along x=5",
            ],
            (
                {(8, 4), (1, 10), (2, 14)},
                [("y", 7), ("x", 5)]
            )
        )
    ]

    for test_case in test_cases:
        actual_output = parse_points_and_folds(test_case.case)

        assert test_case.expected_output == actual_output


def test_get_fold_points():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                {(3, 0), (6, 0), (9, 0), (2, 2)},
                ("x", 5)
            ],
            {(3, 0), (4, 0), (1, 0), (2, 2)}
        )
    ]

    for test_case in test_cases:
        actual_output = get_fold_points(test_case.case[0], test_case.case[1])

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "6,10",
                "0,14",
                "9,10",
                "0,3",
                "10,4",
                "4,11",
                "6,0",
                "6,12",
                "4,1",
                "0,13",
                "10,12",
                "3,4",
                "3,0",
                "8,4",
                "1,10",
                "2,14",
                "8,10",
                "9,0",
                "",
                "fold along y=7",
                "fold along x=5",
            ],
            17
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


# def test_solve2():
#     test_cases = [
#         utils.TestCase(  # Given
#             [],
#             None
#         ),
#     ]

#     for test_case in test_cases:
#         actual_output = solve2(test_case.case)

#         assert test_case.expected_output == actual_output
