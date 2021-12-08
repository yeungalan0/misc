from solutions_day8 import get_value, is_valid, solve1, solve2
import utils


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
                "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
                "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
                "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
                "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
                "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
                "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
                "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
                "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
                "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
            ],
            26
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_get_value():
    index_to_letter_dict = {"a": 0, "b": 1, "c": 2,
                            "d": 3, "e": 4, "f": 5, "g": 6}
    test_cases = [
        utils.TestCase(  # Given
            (["cdfeb", "fcadb", "cdfeb", "cdbaf"], [2, 11, 13, 1, 3, 5, 7]),
            5353
        )
    ]

    for test_case in test_cases:
        actual_output = get_value(
            test_case.case[0], test_case.case[1], index_to_letter_dict)

        assert test_case.expected_output == actual_output


def test_is_valid():
    index_to_letter_dict = {"a": 0, "b": 1, "c": 2,
                            "d": 3, "e": 4, "f": 5, "g": 6}
    test_cases = [
        utils.TestCase(  # Given
            ([1, 2, 3, 5, 7, 11, 13], [
                "acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"]),
            False
        ),
        utils.TestCase(  # Given
            ([1, 2, 13, 5, 7, 11, 3], [
                "acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"]),
            False
        ),
        utils.TestCase(  # Given
            ([2, 11, 13, 1, 3, 5, 7], [
                "acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"]),
            True
        ),
    ]

    for test_case in test_cases:
        actual_output = is_valid(
            test_case.case[0], index_to_letter_dict, test_case.case[1])

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
                "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
                "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
                "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
                "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
                "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
                "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
                "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
                "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
                "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
            ],
            61229
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
