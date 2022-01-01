from solutions_day18 import SnailNumberNode, explode, snail_number_to_tree, snail_number_tree_to_list, solve1, solve2, split
import utils


def test_snail_number_to_tree():
    first_root = SnailNumberNode()
    first_root.left = SnailNumberNode(first_root)
    first_root.left.left = SnailNumberNode(first_root.left, 1)
    first_root.left.right = SnailNumberNode(first_root.left, 2)
    first_root.right = SnailNumberNode(first_root, 3)

    second_root = SnailNumberNode()
    second_root.left = SnailNumberNode(second_root)
    second_root.left.left = SnailNumberNode(second_root.left, 1)
    second_root.left.right = SnailNumberNode(second_root.left, 2)
    second_root.right = SnailNumberNode(second_root)
    second_root.right.left = SnailNumberNode(second_root.right)
    second_root.right.right = SnailNumberNode(second_root.right, 5)
    second_root.right.left.left = SnailNumberNode(second_root.right.left, 3)
    second_root.right.left.right = SnailNumberNode(second_root.right.left, 4)

    test_cases = [
        utils.TestCase(
            [[1, 2], 3],
            first_root
        ),
        utils.TestCase(
            [[1, 2], [[3, 4], 5]],
            second_root
        )
    ]

    for test_case in test_cases:
        actual_output = snail_number_to_tree(test_case.case)

        assert test_case.expected_output == actual_output


def test_snail_number_tree_to_list():
    test_cases = [
        utils.TestCase(
            [[1, 2], 3],
            [[1, 2], 3]
        ),
        utils.TestCase(
            [[1, 2], [[3, 4], 5]],
            [[1, 2], [[3, 4], 5]]
        ),
        utils.TestCase(
            [[[[8, 7], [7, 7]], [[8, 6], [7, 7]]], [[[0, 7], [6, 6]], [8, 7]]],
            [[[[8, 7], [7, 7]], [[8, 6], [7, 7]]], [[[0, 7], [6, 6]], [8, 7]]]
        ),
        utils.TestCase(
            [[[[[9, 8], 1], 2], 3], 4],
            [[[[[9, 8], 1], 2], 3], 4]
        )
    ]

    for test_case in test_cases:
        actual_output = snail_number_tree_to_list(
            snail_number_to_tree(test_case.case))

        assert test_case.expected_output == actual_output


def test_explode():
    test_cases = [
        utils.TestCase(
            [[[[[9, 8], 1], 2], 3], 4],
            [[[[0, 9], 2], 3], 4]
        ),
        utils.TestCase(
            [7, [6, [5, [4, [3, 2]]]]],
            [7, [6, [5, [7, 0]]]]
        ),
        utils.TestCase(
            [[6, [5, [4, [3, 2]]]], 1],
            [[6, [5, [7, 0]]], 3]
        ),
        utils.TestCase(
            [[3, [2, [8, 0]]], [9, [5, [4, [3, 2]]]]],
            [[3, [2, [8, 0]]], [9, [5, [7, 0]]]]
        ),
        utils.TestCase(
            [[3, [2, [1, [7, 3]]]], [6, [5, [4, [3, 2]]]]],
            [[3, [2, [8, 0]]], [9, [5, [7, 0]]]]
        )
    ]

    for test_case in test_cases:
        root = snail_number_to_tree(test_case.case)
        explode(root)
        actual_output = snail_number_tree_to_list(root)

        assert test_case.expected_output == actual_output


def test_split():
    test_cases = [
        utils.TestCase(
            [[[[0, 7], 4], [15, [0, 13]]], [1, 1]],
            [[[[0, 7], 4], [[7, 8], [0, 13]]], [1, 1]]
        ),
        utils.TestCase(
            [[[[0, 7], 4], [[7, 8], [0, 13]]], [1, 1]],
            [[[[0, 7], 4], [[7, 8], [0, [6, 7]]]], [1, 1]]
        )
    ]

    for test_case in test_cases:
        root = snail_number_to_tree(test_case.case)
        split(root)
        actual_output = snail_number_tree_to_list(root)

        assert test_case.expected_output == actual_output


def test_solve1():
    test_cases = [
        utils.TestCase(  # Given test case
            [
                "[[[[4,3],4],4],[7,[[8,4],9]]]",
                "[1,1]"
            ],
            1384
        ),
        utils.TestCase(
            [
                "[1,1]",
                "[2,2]",
                "[3,3]",
                "[4,4]",
            ],
            445
        ),
        utils.TestCase(
            [
                "[1,1]",
                "[2,2]",
                "[3,3]",
                "[4,4]",
                "[5,5]",
            ],
            # [[[[3, 0], [5, 3]], [4, 4]], [5, 5]]
            791
        ),
        utils.TestCase(
            [
                "[1,1]",
                "[2,2]",
                "[3,3]",
                "[4,4]",
                "[5,5]",
                "[6,6]",
            ],
            # [[[[5, 0], [7, 4]], [5, 5]], [6, 6]]
            1137
        ),
        utils.TestCase(
            [
                "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
                "[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
                "[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
                "[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
                "[7,[5,[[3,8],[1,4]]]]",
                "[[2,[2,2]],[8,[8,1]]]",
                "[2,9]",
                "[1,[[[9,3],9],[[9,0],[0,7]]]]",
                "[[[5,[7,4]],7],1]",
                "[[[[4,2],2],6],[8,7]]",
            ],
            # [[[[8, 7], [7, 7]], [[8, 6], [7, 7]]], [[[0, 7], [6, 6]], [8, 7]]]
            3488
        )
    ]

    for test_case in test_cases:
        actual_output = solve1(test_case.case)

        assert test_case.expected_output == actual_output


def test_solve2():
    test_cases = [
        utils.TestCase(  # Given
            [
                "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
                "[[[5,[2,8]],4],[5,[[9,9],0]]]",
                "[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
                "[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
                "[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
                "[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
                "[[[[5,4],[7,7]],8],[[8,3],8]]",
                "[[9,3],[[9,9],[6,[4,9]]]]",
                "[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
                "[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
            ],
            3993
        ),
    ]

    for test_case in test_cases:
        actual_output = solve2(test_case.case)

        assert test_case.expected_output == actual_output
