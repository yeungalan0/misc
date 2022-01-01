from typing import Any, List, Optional, Set, Tuple, Union
import typing
from utils import parse_input
import math
import json


class SnailNumberNode:
    def __init__(self, _parent: "SnailNumberNode" = None, _value: Optional[int] = None) -> None:
        self.value = _value
        self.parent: Optional[SnailNumberNode] = _parent
        self.left: Optional[SnailNumberNode] = None
        self.right: Optional[SnailNumberNode] = None

    def __eq__(self, __o: object) -> bool:
        if not isinstance(__o, SnailNumberNode):
            return NotImplemented

        return self.value == __o.value and (self.parent.value == __o.parent.value if self.parent and __o.parent else self.parent == __o.parent) and self.left == __o.left and self.right == __o.right


def snail_number_to_tree(snail_number: List[Any], parent: SnailNumberNode = None) -> SnailNumberNode:
    if parent is None:
        parent = SnailNumberNode()  # Create root

    left_number: Union[List[Any], int] = snail_number[0]
    right_number: Union[List[Any], int] = snail_number[1]

    if isinstance(left_number, int):
        parent.left = SnailNumberNode(parent, left_number)
    else:
        parent.left = snail_number_to_tree(
            left_number, SnailNumberNode(parent))
    if isinstance(right_number, int):
        parent.right = SnailNumberNode(parent, right_number)
    else:
        parent.right = snail_number_to_tree(
            right_number, SnailNumberNode(parent))

    return parent


@typing.no_type_check
def snail_number_tree_to_list(node: SnailNumberNode, output: Optional[List] = None) -> List[Any]:
    if output is None:
        output = []
    if node.left is None or node.right is None:
        print(
            f"Error, left/right node is None, shouldn't be in this state, L: {node.left}, R: {node.right}")

    if isinstance(node.left.value, int):
        output.append(node.left.value)
    else:
        left = snail_number_tree_to_list(node.left)
        output.append(left)
    if isinstance(node.right.value, int):
        output.append(node.right.value)
    else:
        right = snail_number_tree_to_list(node.right)
        output.append(right)

    return output


@typing.no_type_check
def get_leftmost_leaf(node: SnailNumberNode) -> Optional[SnailNumberNode]:
    while True:
        if not node.parent:
            return None
        elif node.parent.right is node:
            node = node.parent.left
            break
        else:
            node = node.parent

    while node.value is None:
        node = node.right

    return node


@typing.no_type_check
def get_rightmost_leaf(node: SnailNumberNode) -> Optional[SnailNumberNode]:
    while True:
        if not node.parent:
            return None
        elif node.parent.left is node:
            node = node.parent.right
            break
        else:
            node = node.parent

    while node.value is None:
        node = node.left

    return node


@typing.no_type_check
def explode(node: SnailNumberNode, depth: int = 0) -> bool:
    if node.value is not None:
        return False
    elif isinstance(node.left.value, int) and isinstance(node.right.value, int) and depth >= 4:
        leftmost_leaf = get_leftmost_leaf(node)
        rightmost_leaf = get_rightmost_leaf(node)
        if leftmost_leaf is not None:
            leftmost_leaf.value += node.left.value
        if rightmost_leaf is not None:
            rightmost_leaf.value += node.right.value
        node.value = 0
        node.left = None
        node.right = None
        return True

    left_explosion = explode(node.left, depth + 1)
    right_explosion = explode(node.right, depth + 1)
    return left_explosion or right_explosion


def split(node: SnailNumberNode) -> bool:
    if node.value is not None:
        if node.value < 10:
            return False
        else:
            node.left = SnailNumberNode(node, math.floor(node.value/2.0))
            node.right = SnailNumberNode(node, math.ceil(node.value/2.0))
            node.value = None
            return True

    return split(node.left) or split(node.right)  # type: ignore


def reduce(node: SnailNumberNode, depth: int = 0) -> SnailNumberNode:
    while True:
        did_explode = explode(node)
        if not did_explode:
            did_split = split(node)
            if not did_split:
                break

    return node


def add_snail_numbers(node1: SnailNumberNode, node2: SnailNumberNode) -> SnailNumberNode:
    new_number = SnailNumberNode()
    new_number.left = node1
    node1.parent = new_number
    new_number.right = node2
    node2.parent = new_number

    return reduce(new_number)


@typing.no_type_check
def get_magnitude(node: SnailNumberNode) -> int:
    left_value = node.left.value if isinstance(
        node.left.value, int) else get_magnitude(node.left)
    right_value = node.right.value if isinstance(
        node.right.value, int) else get_magnitude(node.right)

    return 3 * left_value + 2 * right_value


def solve1(problem_input: List[str]) -> int:
    root_number: SnailNumberNode = snail_number_to_tree(
        json.loads(problem_input[0]))

    for snail_number_string in problem_input[1:]:
        new_number = snail_number_to_tree(json.loads(snail_number_string))
        root_number = add_snail_numbers(root_number, new_number)

    return get_magnitude(root_number)


def solve2(problem_input: List[str]) -> int:
    all_numbers = set(problem_input)
    max_magnitude = 0

    for number_string in all_numbers:
        subset = all_numbers - set([number_string])
        for number_string_2 in subset:
            magnitude = solve1([number_string, number_string_2])
            max_magnitude = max(max_magnitude, magnitude)

    return max_magnitude


if __name__ == "__main__":
    problem_input = parse_input("inputs/input_day18.txt")

    # problem_input_ints = [int(num) for num in problem_input]

    # problem_input_2d_ints = [[int(num) for num in line]
    #                         for line in problem_input]

    print(solve1(problem_input))
    print(solve2(problem_input))
