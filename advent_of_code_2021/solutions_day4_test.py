from solutions_day4 import check_winner, generate_number_dict, parse_bn_and_tables, solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"
                ""
                "22 13 17 11  0",
                " 8  2 23  4 24",
                "21  9 14 16  7",
                " 6 10  3 18  5",
                " 1 12 20 15 19",
                "",
                " 3 15  0  2 22",
                " 9 18 13 17  5",
                "19  8  7 25 23",
                "20 11 10 24  4",
                "14 21 16 12  6",
                "",
                "14 21 17 24  4",
                "10 16 15  9 19",
                "18  8 23 26 20",
                "22 11 13  6  5",
                " 2  0 12  3  7",
            ],
            4512
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_parse_bn_and_tables():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
                "",
                "22 13 17 11  0",
                " 8  2 23  4 24",
                "21  9 14 16  7",
                " 6 10  3 18  5",
                " 1 12 20 15 19",
                "",
                " 3 15  0  2 22",
                " 9 18 13 17  5",
                "19  8  7 25 23",
                "20 11 10 24  4",
                "14 21 16 12  6",
            ],
            (
                ["7", "4", "9", "5", "11", "17", "23", "2", "0", "14", "21", "24", "10", "16",
                    "13", "6", "15", "25", "12", "22", "18", "20", "8", "19", "3", "26", "1"],
                [
                    [
                        ["22", "13", "17", "11", "0"],
                        ["8", "2", "23", "4", "24"],
                        ["21", "9", "14", "16", "7"],
                        ["6", "10", "3", "18", "5"],
                        ["1", "12", "20", "15", "19"],
                    ],
                    [
                        ["3", "15", "0", "2", "22"],
                        ["9", "18", "13", "17", "5"],
                        ["19", "8", "7", "25", "23"],
                        ["20", "11", "10", "24", "4"],
                        ["14", "21", "16", "12", "6"],
                    ]
                ]
            )
        )
    ]

    for test_case in test_cases:
        actual_output = parse_bn_and_tables(test_case.case)

        assert test_case.expected_output == actual_output


def test_generate_number_dict():
    test_cases = [
        utils.TestCase(
            [
                [
                    ["22", "13"],
                    ["8", "2"],
                ],
                [
                    ["3", "15"],
                    ["8", "13"],
                ]
            ],
            {
                "22": [(0, 0, 0)],
                "13": [(0, 0, 1), (1, 1, 1)],
                "8": [(0, 1, 0), (1, 1, 0)],
                "2": [(0, 1, 1)],
                "3": [(1, 0, 0)],
                "15": [(1, 0, 1)]
            }
        )
    ]

    for test_case in test_cases:
        actual_output = generate_number_dict(test_case.case)

        assert test_case.expected_output == actual_output


def test_check_winner():
    test_cases = [
        utils.TestCase(  # row victory
            [
                [
                    ["22", "-13", "17", "11", "0"],
                    ["8", "2", "-23", "4", "24"],
                    ["21", "9", "-14", "16", "7"],
                    ["-6", "-10", "-3", "-18", "-5"],
                    ["1", "12", "20", "15", "19"]
                ],
                3, 2
            ],
            624
        ),
        utils.TestCase(  # col victory
            [
                [
                    ["22", "-13", "17", "11", "-0"],
                    ["8", "2", "-23", "4", "-24"],
                    ["21", "9", "-14", "16", "-7"],
                    ["6", "-10", "-3", "-18", "-5"],
                    ["1", "12", "20", "15", "-19"]
                ],
                0, 4
            ],
            0
        ),
        utils.TestCase(  # No victory
            [
                [
                    ["22", "-13", "17", "11", "-0"],
                    ["8", "2", "-23", "4", "-24"],
                    ["21", "9", "-14", "16", "-7"],
                    ["6", "-10", "-3", "-18", "5"],
                    ["1", "12", "20", "15", "-19"]
                ],
                0, 4
            ],
            None
        )
    ]

    for test_case in test_cases:
        actual_output = check_winner(
            test_case.case[0], test_case.case[1], test_case.case[2])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1"
                ""
                "22 13 17 11  0",
                " 8  2 23  4 24",
                "21  9 14 16  7",
                " 6 10  3 18  5",
                " 1 12 20 15 19",
                "",
                " 3 15  0  2 22",
                " 9 18 13 17  5",
                "19  8  7 25 23",
                "20 11 10 24  4",
                "14 21 16 12  6",
                "",
                "14 21 17 24  4",
                "10 16 15  9 19",
                "18  8 23 26 20",
                "22 11 13  6  5",
                " 2  0 12  3  7",
            ],
            1924
        )
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
