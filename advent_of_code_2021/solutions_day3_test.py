from solutions_day3 import get_gas_rating, solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "00100",
                "11110",
                "10110",
                "10111",
                "10101",
                "01111",
                "00111",
                "11100",
                "10000",
                "11001",
                "00010",
                "01010",
            ],
            198
        ),
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "00100",
                "11110",
                "10110",
                "10111",
                "10101",
                "01111",
                "00111",
                "11100",
                "10000",
                "11001",
                "00010",
                "01010",
            ],
            230
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output


def test_get_gas_rating_oxygen():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "00100",
                "11110",
                "10110",
                "10111",
                "10101",
                "01111",
                "00111",
                "11100",
                "10000",
                "11001",
                "00010",
                "01010",
            ],
            23
        )
    ]

    for test_case in test_cases:
        actual_output = get_gas_rating(
            oxygen=True, problem_input=test_case.case)

        assert test_case.expected_output == actual_output


def test_get_gas_rating_co2():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "00100",
                "11110",
                "10110",
                "10111",
                "10101",
                "01111",
                "00111",
                "11100",
                "10000",
                "11001",
                "00010",
                "01010",
            ],
            10
        )
    ]

    for test_case in test_cases:
        actual_output = get_gas_rating(
            oxygen=False, problem_input=test_case.case)

        assert test_case.expected_output == actual_output
