class Node:
    def __init__(self, operator, value, left, right):
        self.operator = operator
        self.value = value
        self.left = left
        self.right = right

    def result(self):
        if self.operator == "x":
            return self.left.result() * self.right.result()
        elif self.operator == "÷":
            return self.left.result() / self.right.result()
        elif self.operator == "+":
            return self.left.result() + self.right.result()
        elif self.operator == "-":
            return self.left.result() - self.right.result()
        else:
            return float(self.value)

    def __str__(self):
        if self.operator == "x":
            return f"({self.left} x {self.right})"
        elif self.operator == "÷":
            return f"({self.left} ÷ {self.right})"
        elif self.operator == "+":
            return f"({self.left} + {self.right})"
        elif self.operator == "-":
            return f"({self.left} - {self.right})"
        else:
            return str(self.value)

import unittest

class TestNode(unittest.TestCase):
    def test_node(self):
        tree = Node(
            "÷",
            None,
            Node(
                "+",
                None,
                Node("", 7, None, None),
                Node(
                    "x",
                    None,
                    Node("-", None, Node("", 3, None, None), Node("", 2, None, None)),
                    Node("", 5, None, None)
                )
            ),
            Node("", 6, None, None)
        )
        self.assertEqual(str(tree), "((7 + ((3 - 2) x 5)) ÷ 6)")
        self.assertEqual(tree.result(), 2)

if __name__ == "__main__":
    unittest.main()

