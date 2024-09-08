package main

import (
	"fmt"
	"os"
)

type Node struct {
	left     *Node
	right    *Node
	operator string
	value    float64
}

func (n *Node) Result() float64 {
	switch n.operator {
	case "x":
		return n.left.Result() * n.right.Result()
	case "÷":
		return n.left.Result() / n.right.Result()
	case "+":
		return n.left.Result() + n.right.Result()
	case "-":
		return n.left.Result() - n.right.Result()
	default:
		return n.value
	}
}

func (n *Node) String() string {
	switch n.operator {
	case "x":
		return fmt.Sprintf("(%s x %s)", n.left.String(), n.right.String())
	case "÷":
		return fmt.Sprintf("(%s ÷ %s)", n.left.String(), n.right.String())
	case "+":
		return fmt.Sprintf("(%s + %s)", n.left.String(), n.right.String())
	case "-":
		return fmt.Sprintf("(%s - %s)", n.left.String(), n.right.String())
	default:
		return fmt.Sprintf("%v", n.value)
	}
}

func main() {
	tree := &Node{
		operator: "÷",
		left: &Node{
			operator: "+",
			left:     &Node{value: 7},
			right: &Node{
				operator: "x",
				left: &Node{
					operator: "-",
					left:     &Node{value: 3},
					right:    &Node{value: 2},
				},
				right: &Node{value: 5},
			},
		},
		right: &Node{value: 6},
	}

	expectedStr := "((7 + ((3 - 2) x 5)) ÷ 6)"
	if tree.String() != expectedStr {
		fmt.Printf("Expected %s, got %s\n", expectedStr, tree.String())
		os.Exit(1)
	}

	expectedResult := 2.0
	if tree.Result() != expectedResult {
		fmt.Printf("Expected %f, got %f\n", expectedResult, tree.Result())
		os.Exit(1)
	}

	fmt.Println(tree.String())
	fmt.Println(tree.Result())
}
