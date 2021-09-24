class GraphNode:
    def __init__(self, init_value) -> None:
        self.value = init_value
        self.children = []
        self.adjacent = self.children


class TreeNode(GraphNode):
    def get_left_child(self):
        return self.get_child(0)

    def get_right_child(self):
        return self.get_child(1)

    def get_child(self, index):
        try:
            return self.children[index]
        except IndexError:
            return None

# Assumes root node is 0


def tree_builder(tree_dict, root_val):
    root = tree_builder_helper(tree_dict, root_val)

    return root


def tree_builder_helper(tree_dict, node_val):
    node = TreeNode(node_val)
    for child_val in tree_dict.get(node_val, []):
        child_node = tree_builder_helper(tree_dict, child_val)
        node.children.append(child_node)

    return node


def print_node(node: GraphNode):
    print(node.value)

# Assuming is a binary tree
def in_order_traversal(tree_node: TreeNode, visit=print_node):
    if not tree_node or not tree_node.value:
        return

    left_child = tree_node.get_left_child()
    right_child = tree_node.get_right_child()

    in_order_traversal(left_child)
    visit(tree_node)
    in_order_traversal(right_child)


def pre_order_traversal(tree_node: TreeNode, visit=print_node):
    if not tree_node or not tree_node.value:
        return


    left_child = tree_node.get_left_child()
    right_child = tree_node.get_right_child()

    visit(tree_node)
    pre_order_traversal(left_child)
    pre_order_traversal(right_child)

def post_order_traversal(tree_node: TreeNode, visit=print_node):
    if not tree_node or not tree_node.value:
        return


    left_child = tree_node.get_left_child()
    right_child = tree_node.get_right_child()

    post_order_traversal(left_child)
    post_order_traversal(right_child)
    visit(tree_node)

if __name__ == "__main__":
    tree_dict = {
        10: [5, 20],
        5: [3, 7],
        20: [None, 30]
    }

    tree_root = tree_builder(tree_dict, 10)

    print("In order:")
    in_order_traversal(tree_root)

    print("Pre order:")
    pre_order_traversal(tree_root)

    print("Post order:")
    post_order_traversal(tree_root)