from dataclasses import dataclass
from typing import Any, List


@dataclass
class TestCase:
    """Class to keep track of test case and expected output"""
    __test__ = False  # Needed for pytest to ignore this class
    case: List[str]
    expected_output: Any


def parse_input(input_file) -> List[str]:
    with open(input_file, "r", encoding="UTF-8") as file:
        lines = file.readlines()
        return [line.rstrip() for line in lines]
