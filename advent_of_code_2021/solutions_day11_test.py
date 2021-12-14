from solutions_day11 import flash_octopuses, increase_energy, solve1, solve2
import utils


def test_increase_energy():
    test_cases = [
        utils.TestCase(  # Given
            [
                [1, 1, 1, 1, 1],
                [1, 9, 9, 9, 1],
                [1, 9, 1, 9, 1],
                [1, 9, 9, 9, 1],
                [1, 1, 1, 1, 1],
            ],
            [
                [2, 2, 2, 2, 2],
                [2, 10, 10, 10, 2],
                [2, 10, 2, 10, 2],
                [2, 10, 10, 10, 2],
                [2, 2, 2, 2, 2],
            ]
        ),
    ]

    for test_case in test_cases:
        actual_output = increase_energy(test_case.case)

        assert test_case.expected_output == actual_output


def test_flash_octopuses():
    test_cases = [
        utils.TestCase(
            [
                [2, 2, 2, 2, 2],
                [2, 10, 10, 10, 2],
                [2, 10, 2, 10, 2],
                [2, 10, 10, 10, 2],
                [2, 2, 2, 2, 2],
            ],
            ([
                [3, 4, 5, 4, 3],
                [4, 0, 0, 0, 4],
                [5, 0, 0, 0, 5],
                [4, 0, 0, 0, 4],
                [3, 4, 5, 4, 3],
            ], 9)
        ),
        utils.TestCase(
            increase_energy([
                [6, 5, 9, 4, 2, 5, 4, 3, 3, 4],
                [3, 8, 5, 6, 9, 6, 5, 8, 2, 2],
                [6, 3, 7, 5, 6, 6, 7, 2, 8, 4],
                [7, 2, 5, 2, 4, 4, 7, 2, 5, 7],
                [7, 4, 6, 8, 4, 9, 6, 5, 8, 9],
                [5, 2, 7, 8, 6, 3, 5, 7, 5, 6],
                [3, 2, 8, 7, 9, 5, 2, 8, 3, 2],
                [7, 9, 9, 3, 9, 9, 2, 2, 4, 5],
                [5, 9, 5, 7, 9, 5, 9, 6, 6, 5],
                [6, 3, 9, 4, 8, 6, 2, 6, 3, 7],
            ]),
            ([
                [8, 8, 0, 7, 4, 7, 6, 5, 5, 5],
                [5, 0, 8, 9, 0, 8, 7, 0, 5, 4],
                [8, 5, 9, 7, 8, 8, 9, 6, 0, 8],
                [8, 4, 8, 5, 7, 6, 9, 6, 0, 0],
                [8, 7, 0, 0, 9, 0, 8, 8, 0, 0],
                [6, 6, 0, 0, 0, 8, 8, 9, 8, 9],
                [6, 8, 0, 0, 0, 0, 5, 9, 4, 3],
                [0, 0, 0, 0, 0, 0, 7, 4, 5, 6],
                [9, 0, 0, 0, 0, 0, 0, 8, 7, 6],
                [8, 7, 0, 0, 0, 0, 6, 8, 4, 8],
            ], 35)
        ),
    ]

    for test_case in test_cases:
        actual_output = flash_octopuses(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(
            ([
                [1, 1, 1, 1, 1],
                [1, 9, 9, 9, 1],
                [1, 9, 1, 9, 1],
                [1, 9, 9, 9, 1],
                [1, 1, 1, 1, 1],
            ], 1),
            ([
                [3, 4, 5, 4, 3],
                [4, 0, 0, 0, 4],
                [5, 0, 0, 0, 5],
                [4, 0, 0, 0, 4],
                [3, 4, 5, 4, 3],
            ], 9)
        ),
        utils.TestCase(
            ([
                [1, 1, 1, 1, 1],
                [1, 9, 9, 9, 1],
                [1, 9, 1, 9, 1],
                [1, 9, 9, 9, 1],
                [1, 1, 1, 1, 1],
            ], 2),
            ([
                [4, 5, 6, 5, 4],
                [5, 1, 1, 1, 5],
                [6, 1, 1, 1, 6],
                [5, 1, 1, 1, 5],
                [4, 5, 6, 5, 4],
            ], 9)
        ),
        utils.TestCase(  # Given test case
            ([
                [5, 4, 8, 3, 1, 4, 3, 2, 2, 3],
                [2, 7, 4, 5, 8, 5, 4, 7, 1, 1],
                [5, 2, 6, 4, 5, 5, 6, 1, 7, 3],
                [6, 1, 4, 1, 3, 3, 6, 1, 4, 6],
                [6, 3, 5, 7, 3, 8, 5, 4, 7, 8],
                [4, 1, 6, 7, 5, 2, 4, 6, 4, 5],
                [2, 1, 7, 6, 8, 4, 1, 7, 2, 1],
                [6, 8, 8, 2, 8, 8, 1, 1, 3, 4],
                [4, 8, 4, 6, 8, 4, 8, 5, 5, 4],
                [5, 2, 8, 3, 7, 5, 1, 5, 2, 6],
            ], 10),
            ([
                [0, 4, 8, 1, 1, 1, 2, 9, 7, 6],
                [0, 0, 3, 1, 1, 1, 2, 0, 0, 9],
                [0, 0, 4, 1, 1, 1, 2, 5, 0, 4],
                [0, 0, 8, 1, 1, 1, 1, 4, 0, 6],
                [0, 0, 9, 9, 1, 1, 1, 3, 0, 6],
                [0, 0, 9, 3, 5, 1, 1, 2, 3, 3],
                [0, 4, 4, 2, 3, 6, 1, 1, 3, 0],
                [5, 5, 3, 2, 2, 5, 2, 3, 5, 0],
                [0, 5, 3, 2, 2, 5, 0, 6, 0, 0],
                [0, 0, 3, 2, 2, 4, 0, 0, 0, 0],
            ], 204)
        ),
        utils.TestCase(  # Given test case
            ([
                [5, 4, 8, 3, 1, 4, 3, 2, 2, 3],
                [2, 7, 4, 5, 8, 5, 4, 7, 1, 1],
                [5, 2, 6, 4, 5, 5, 6, 1, 7, 3],
                [6, 1, 4, 1, 3, 3, 6, 1, 4, 6],
                [6, 3, 5, 7, 3, 8, 5, 4, 7, 8],
                [4, 1, 6, 7, 5, 2, 4, 6, 4, 5],
                [2, 1, 7, 6, 8, 4, 1, 7, 2, 1],
                [6, 8, 8, 2, 8, 8, 1, 1, 3, 4],
                [4, 8, 4, 6, 8, 4, 8, 5, 5, 4],
                [5, 2, 8, 3, 7, 5, 1, 5, 2, 6],
            ], 100),
            ([
                [0, 3, 9, 7, 6, 6, 6, 8, 6, 6],
                [0, 7, 4, 9, 7, 6, 6, 9, 1, 8],
                [0, 0, 5, 3, 9, 7, 6, 9, 3, 3],
                [0, 0, 0, 4, 2, 9, 7, 8, 2, 2],
                [0, 0, 0, 4, 2, 2, 9, 8, 9, 2],
                [0, 0, 5, 3, 2, 2, 2, 8, 7, 7],
                [0, 5, 3, 2, 2, 2, 2, 9, 6, 6],
                [9, 3, 2, 2, 2, 2, 8, 9, 6, 6],
                [7, 9, 2, 2, 2, 8, 6, 8, 6, 6],
                [6, 7, 8, 9, 9, 9, 8, 7, 6, 6],
            ], 1656)
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case[0], test_case.case[1])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                [5, 4, 8, 3, 1, 4, 3, 2, 2, 3],
                [2, 7, 4, 5, 8, 5, 4, 7, 1, 1],
                [5, 2, 6, 4, 5, 5, 6, 1, 7, 3],
                [6, 1, 4, 1, 3, 3, 6, 1, 4, 6],
                [6, 3, 5, 7, 3, 8, 5, 4, 7, 8],
                [4, 1, 6, 7, 5, 2, 4, 6, 4, 5],
                [2, 1, 7, 6, 8, 4, 1, 7, 2, 1],
                [6, 8, 8, 2, 8, 8, 1, 1, 3, 4],
                [4, 8, 4, 6, 8, 4, 8, 5, 5, 4],
                [5, 2, 8, 3, 7, 5, 1, 5, 2, 6],
            ],
            195
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
