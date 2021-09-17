from typing import Optional
from utils import convert_to_linked_list, print_linked_list

# Definition for singly-linked list.


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def addTwoNumbers_original(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        solution_node = l1
        current_node = l2
        carry = False

        while True:
            new_value = solution_node.val + current_node.val + carry
            carry = new_value >= 10
            solution_node.val = new_value % 10

            if solution_node.next is None or current_node.next is None:
                break

            solution_node = solution_node.next
            current_node = current_node.next

        if current_node.next is not None:
            solution_node.next = current_node.next

        while carry:
            if solution_node.next is None:
                solution_node.next = ListNode(1)
                carry = False
                break

            solution_node = solution_node.next
            new_value = solution_node.val + carry
            carry = new_value >= 10
            solution_node.val = new_value % 10

        return l1

    # Simplified the logic
    def addTwoNumbers_2(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head_node = None
        tail_node = None
        carry = False
        l1_current = l1
        l2_current = l2

        while l1_current != None or l2_current != None or carry:
            l1_val = l1_current.val if l1_current != None else 0
            l2_val = l2_current.val if l2_current != None else 0

            new_value = l1_val + l2_val + carry

            new_node = ListNode(new_value % 10)

            carry = new_value > 9

            if head_node == None:
                head_node = new_node

            if tail_node == None:
                tail_node = new_node
            else:
                tail_node.next = new_node
                tail_node = tail_node.next

            l1_current = l1_current.next if l1_current != None else None
            l2_current = l2_current.next if l2_current != None else None

        return head_node

    # Even more optimized logic
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        l1_current = l1
        l2_current = l2
        dummy_head_node = ListNode()
        tail_node = dummy_head_node
        carry = False

        while l1_current or l2_current or carry:
            l1_val = l1_current.val if l1_current else 0
            l2_val = l2_current.val if l2_current else 0
            carry_val = 1 if carry else 0

            new_val = l1_val + l2_val + carry_val

            tail_node.next = ListNode(new_val % 10)
            tail_node = tail_node.next

            carry = new_val > 9
            l1_current = l1_current.next if l1_current else None
            l2_current = l2_current.next if l2_current else None

        return dummy_head_node.next


if __name__ == "__main__":
    out = Solution().addTwoNumbers(convert_to_linked_list(
        [2, 4, 3]), convert_to_linked_list([5, 6, 4]))

    print_linked_list(convert_to_linked_list([2, 4, 3]))
    print_linked_list(convert_to_linked_list([5, 6, 4]))
    print_linked_list(out)
