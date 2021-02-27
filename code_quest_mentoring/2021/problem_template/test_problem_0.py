import pytest
import problem_0

expected_output = None

with open('expected_output.txt', 'r') as file:
    expected_output = file.read().rstrip("\n")

def test_main(capsys):
    problem_0.main()
    actual_output = capsys.readouterr().out.rstrip("\n")
    print(f"Actual output:\n{actual_output}\n")
    print(f"Expected output:\n{expected_output}")
    assert actual_output == expected_output