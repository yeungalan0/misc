from solutions_day14 import parse_input_polymers, solve1, solve2
import utils


def test_parse_input_polymers():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "NNCB",
                "",
                "CH -> B",
                "HH -> N",
                "CB -> H"
            ],
            (["N", "N", "C", "B"], {"CH": "B", "HH": "N", "CB": "H"})
        )
    ]

    for test_case in test_cases:
        actual_output = parse_input_polymers(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "NNCB",
                "",
                "CH -> B",
                "HH -> N",
                "CB -> H",
                "NH -> C",
                "HB -> C",
                "HC -> B",
                "HN -> C",
                "NN -> C",
                "BH -> H",
                "NC -> B",
                "NB -> B",
                "BN -> B",
                "BB -> N",
                "BC -> B",
                "CC -> N",
                "CN -> C",
            ],
            1588
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                "NNCB",
                "",
                "CH -> B",
                "HH -> N",
                "CB -> H",
                "NH -> C",
                "HB -> C",
                "HC -> B",
                "HN -> C",
                "NN -> C",
                "BH -> H",
                "NC -> B",
                "NB -> B",
                "BN -> B",
                "BB -> N",
                "BC -> B",
                "CC -> N",
                "CN -> C",
            ],
            2188189693529
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
