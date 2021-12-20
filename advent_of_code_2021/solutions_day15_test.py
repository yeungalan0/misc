from solutions_day15 import solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                [1, 1, 6, 3, 7, 5, 1, 7, 4, 2],
                [1, 3, 8, 1, 3, 7, 3, 6, 7, 2],
                [2, 1, 3, 6, 5, 1, 1, 3, 2, 8],
                [3, 6, 9, 4, 9, 3, 1, 5, 6, 9],
                [7, 4, 6, 3, 4, 1, 7, 1, 1, 1],
                [1, 3, 1, 9, 1, 2, 8, 1, 3, 7],
                [1, 3, 5, 9, 9, 1, 2, 4, 2, 1],
                [3, 1, 2, 5, 4, 2, 1, 6, 3, 9],
                [1, 2, 9, 3, 1, 3, 8, 5, 2, 1],
                [2, 3, 1, 1, 9, 4, 4, 5, 8, 1],
            ],
            40
        ),
        utils.TestCase(  # Given
            [
                [1, 9, 9, 9, 9],
                [1, 9, 1, 1, 1],
                [1, 1, 1, 9, 1],
                [9, 9, 9, 9, 1],
                [9, 9, 9, 9, 1],
            ],
            10
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


# def test_solve2():
#     test_cases = [
#         utils.TestCase(  # Given
#             [
#                 [1, 1, 6, 3, 7, 5, 1, 7, 4, 2],
#                 [1, 3, 8, 1, 3, 7, 3, 6, 7, 2],
#                 [2, 1, 3, 6, 5, 1, 1, 3, 2, 8],
#                 [3, 6, 9, 4, 9, 3, 1, 5, 6, 9],
#                 [7, 4, 6, 3, 4, 1, 7, 1, 1, 1],
#                 [1, 3, 1, 9, 1, 2, 8, 1, 3, 7],
#                 [1, 3, 5, 9, 9, 1, 2, 4, 2, 1],
#                 [3, 1, 2, 5, 4, 2, 1, 6, 3, 9],
#                 [1, 2, 9, 3, 1, 3, 8, 5, 2, 1],
#                 [2, 3, 1, 1, 9, 4, 4, 5, 8, 1],
#             ],
#             315
#         )
#     ]

#     for test_case in test_cases:
#         actual_output = solve2(test_case.case)

#         assert test_case.expected_output == actual_output
