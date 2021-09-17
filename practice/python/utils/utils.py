class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def convert_to_linked_list(nodes_list):
    head = ListNode(nodes_list[0])
    current = head

    for val in nodes_list[1:]:
        current.next = ListNode(val)
        current = current.next

    return head


def print_linked_list(head):
    list_string = f"[{head.val}"

    current_node = head.next
    while current_node:
        list_string += f", {current_node.val}"
        current_node = current_node.next

    print(list_string + "]")
