from solutions_day12 import solve1, solve2, GraphNode, create_graph
import utils


def test_create_graph():
    node_start = GraphNode("start")
    node_A = GraphNode("A")
    node_b = GraphNode("b")
    node_end = GraphNode("end")

    node_start.adj = [node_A, node_b]
    node_A.adj = [node_start, node_b, node_end]
    node_b.adj = [node_start, node_A, node_end]
    node_end.adj = [node_A, node_b]
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "start-A",
                "start-b",
                "A-b",
                "A-end",
                "b-end",
            ],
            {"start": node_start, "A": node_A,
                "b": node_b, "end": node_end}
        )
    ]

    for test_case in test_cases:
        actual_output = create_graph(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "start-A",
                "start-b",
                "A-c",
                "A-b",
                "b-d",
                "A-end",
                "b-end",
            ],
            10
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                "start-A",
                "start-b",
                "A-c",
                "A-b",
                "b-d",
                "A-end",
                "b-end",
            ],
            36
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
