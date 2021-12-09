from solutions_day9 import find_low_point_coordinates, get_basin_size, solve1, solve2
import utils


def test_find_low_point_coordinates():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                [2, 1, 9, 9, 9, 4, 3, 2, 1, 0],
                [3, 9, 8, 7, 8, 9, 4, 9, 2, 1],
                [9, 8, 5, 6, 7, 8, 9, 8, 9, 2],
                [8, 7, 6, 7, 8, 9, 6, 7, 8, 9],
                [9, 8, 9, 9, 9, 6, 5, 6, 7, 8]
            ],
            [(0, 1), (0, 9), (2, 2), (4, 6)]
        )
    ]

    for test_case in test_cases:
        actual_output = find_low_point_coordinates(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                [2, 1, 9, 9, 9, 4, 3, 2, 1, 0],
                [3, 9, 8, 7, 8, 9, 4, 9, 2, 1],
                [9, 8, 5, 6, 7, 8, 9, 8, 9, 2],
                [8, 7, 6, 7, 8, 9, 6, 7, 8, 9],
                [9, 8, 9, 9, 9, 6, 5, 6, 7, 8]
            ],
            15
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_get_basin_size():
    heightmap = [
        [2, 1, 9, 9, 9, 4, 3, 2, 1, 0],
        [3, 9, 8, 7, 8, 9, 4, 9, 2, 1],
        [9, 8, 5, 6, 7, 8, 9, 8, 9, 2],
        [8, 7, 6, 7, 8, 9, 6, 7, 8, 9],
        [9, 8, 9, 9, 9, 6, 5, 6, 7, 8]
    ]

    test_cases = [
        utils.TestCase(  # Given test case
            ((0, 1), heightmap),
            3
        ),
        utils.TestCase(  # Given test case
            ((0, 9), heightmap),
            9
        ),
        utils.TestCase(  # Given test case
            ((2, 2), heightmap),
            14
        )
    ]

    for test_case in test_cases:
        actual_output = get_basin_size(test_case.case[0], test_case.case[1])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                [2, 1, 9, 9, 9, 4, 3, 2, 1, 0],
                [3, 9, 8, 7, 8, 9, 4, 9, 2, 1],
                [9, 8, 5, 6, 7, 8, 9, 8, 9, 2],
                [8, 7, 6, 7, 8, 9, 6, 7, 8, 9],
                [9, 8, 9, 9, 9, 6, 5, 6, 7, 8]
            ],
            1134
        )
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
