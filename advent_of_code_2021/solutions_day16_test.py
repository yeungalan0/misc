from solutions_day16 import convert_to_binary, solve1, solve2
import utils


def test_convert_to_binary():
    test_cases = [
        utils.TestCase(
            "D2FE28",
            "110100101111111000101000"
        ),
        utils.TestCase(
            "38006F45291200",
            "00111000000000000110111101000101001010010001001000000000"
        ),
        utils.TestCase(
            "EE00D40C823060",
            "11101110000000001101010000001100100000100011000001100000"
        )
    ]

    for test_case in test_cases:
        actual_output = convert_to_binary(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(
            "8A004A801A8002F478",
            16
        ),
        utils.TestCase(
            "620080001611562C8802118E34",
            12
        ),
        utils.TestCase(
            "C0015000016115A2E0802F182340",
            23
        ),
        utils.TestCase(
            "A0016C880162017C3686B18A3D4780",
            31
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(
            "C200B40A82",
            3
        ),
        utils.TestCase(
            "04005AC33890",
            54
        ),
        utils.TestCase(
            "880086C3E88112",
            7
        ),
        utils.TestCase(
            "CE00C43D881120",
            9
        ),
        utils.TestCase(
            "D8005AC2A8F0",
            1
        ),
        utils.TestCase(
            "F600BC2D8F",
            0
        ),
        utils.TestCase(
            "9C005AC2F8F0",
            0
        ),
        utils.TestCase(
            "9C0141080250320F1802104A08",
            1
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
