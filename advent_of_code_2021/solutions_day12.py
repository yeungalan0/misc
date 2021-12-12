from typing import List, Dict, Optional
from utils import parse_input


class GraphNode:
    def __init__(self, _value: str):
        self.value: str = _value
        self.adj: List[GraphNode] = []

    def __eq__(self, o) -> bool:
        return self.value == o.value and [gn.value for gn in self.adj] == [gn.value for gn in o.adj]


def create_graph(connections: List[str]) -> Dict[str, GraphNode]:
    graph_dict: Dict[str, GraphNode] = {}

    for connection in connections:
        first_str, sec_str = connection.split("-")

        if first_str not in graph_dict:
            graph_dict[first_str] = GraphNode(first_str)

        if sec_str not in graph_dict:
            graph_dict[sec_str] = GraphNode(sec_str)

        first_node = graph_dict[first_str]
        sec_node = graph_dict[sec_str]
        first_node.adj.append(sec_node)
        sec_node.adj.append(first_node)

    return graph_dict


def dfs_paths(gn: GraphNode, path: List[str], paths: List[str]) -> List[str]:
    path.append(gn.value)

    if gn.value == "end":
        paths.append(",".join(path))

    for adj_node in gn.adj:
        if (adj_node.value.islower() and adj_node.value not in path) or adj_node.value.isupper():
            new_path = path.copy()
            dfs_paths(adj_node, new_path, paths)

    return paths


def dfs_paths2(gn: GraphNode, path: List[str], paths: List[str], twice_visited: bool = False) -> List[str]:
    path.append(gn.value)

    if gn.value == "end":
        paths.append(",".join(path))
        return paths

    for adj_node in gn.adj:
        if twice_visited:
            is_valid_lower = adj_node.value != "start" and adj_node.value.islower(
            ) and adj_node.value not in path
        else:
            is_valid_lower = adj_node.value != "start" and adj_node.value.islower()

        if adj_node.value.isupper() or is_valid_lower:
            new_path = path.copy()
            # This stumped me for a while, gotta make sure twice visited isn't set with uppercase chars
            new_twice_visted = twice_visited or (
                adj_node.value.islower() and adj_node.value in path)
            dfs_paths2(adj_node, new_path, paths, new_twice_visted)

    return paths


def solve1(problem_input: List[str]) -> int:
    graph_dict = create_graph(problem_input)

    return len(dfs_paths(graph_dict["start"], [], []))


def solve2(problem_input: List[str]) -> int:
    graph_dict = create_graph(problem_input)
    paths = dfs_paths2(graph_dict["start"], [], [])
    # for path in paths:
    #     print(path)

    return len(paths)


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day12.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
