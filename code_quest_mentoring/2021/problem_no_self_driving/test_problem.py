import pytest
import problem
import io
import os

DIR_PATH = os.path.dirname(os.path.realpath(__file__))

expected_output = None
input = None

# Read in our expected output
with open(f"{DIR_PATH}/expected_output.txt", 'r') as file:
    expected_output = file.read().rstrip("\n")

# Read in our expected input
with open(f"{DIR_PATH}/input.txt", 'r') as file:
    input = file.read()

def test_main(capsys, monkeypatch):
    # Mock stdin to pass in our input
    monkeypatch.setattr('sys.stdin', io.StringIO(input))

    # Call our main function to ultimately solve the problem
    problem.main()

    # Capture the program output from stdout/stderr
    actual_output = capsys.readouterr().out.rstrip("\n")

    # Ensure the actual output matches the expected output
    print(f"Actual output:\n{actual_output}\n")
    print(f"Expected output:\n{expected_output}")
    assert actual_output == expected_output