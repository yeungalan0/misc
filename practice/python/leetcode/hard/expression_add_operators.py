from typing import List


class Solution:
    # Things to remember:
    # Order of operations matter
    def addOperators(self, num: str, target: int) -> List[str]:
        valid_list = []

        # Using the & symbol as a no op operator to combine multiple digits
        self.add_operators_helper(num, "", valid_list, target, "+-*&")

        return valid_list

    def add_operators_helper(self, num: str, result, valid_list, target, operators):
        if len(num) == 0:
            if self.is_equal(result, target):
                valid_list.append(result)
            return

        if len(num) > 1:
            for op in operators:
                if op == "&":
                    new_result = result + num[0]
                else:
                    new_result = result + num[0] + op
                self.add_operators_helper(
                    num[1:], new_result, valid_list, target, operators)
        else:
            new_result = result + num[0]
            self.add_operators_helper(
                "", new_result, valid_list, target, operators)

    # Mistake: Forgot to do order of operations here...
    # def is_equal(self, result_string, target_int):
    #     total = int(result_string[0])
    #     pair = ""

    #     for char in result_string[1:]:
    #         pair = pair + char
    #         if len(pair) == 2:
    #             operator = pair[0]
    #             value = int(pair[1])
    #             if operator == "+":
    #                 total += value
    #             elif operator == "*":
    #                 total *= value
    #             else:
    #                 total -= value
    #             pair = ""

    #     return total == target_int

    def is_equal(self, result_string, target_int):
        expression_list = self.build_expression_list(result_string)
        return self.evaluate(expression_list) == target_int

    def evaluate(self, expression_list):
        total = int(expression_list.pop(0))

        while len(expression_list) > 0:
            current = expression_list.pop(0)

            if current == "+":
                return total + self.evaluate(expression_list)
            elif current == "-":
                return total - self.evaluate(expression_list)
            elif current == "*":
                continue
            else:
                total *= int(current)

        return total

    def build_expression_list(self, expression_string):
        expression_list = []

        digits = ""
        for char in expression_string:
            if char in "+-*":
                expression_list.append(digits)
                expression_list.append(char)
                digits = ""
            else:
                digits += char

        if digits != "":
            expression_list.append(digits)

        return expression_list


if __name__ == "__main__":
    out = Solution().addOperators("105", 5)
    print(out)
