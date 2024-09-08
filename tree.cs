using System;
using NUnit.Framework;

public class Node
{
    private string operatorSymbol;
    private double? value;
    private Node left;
    private Node right;

    public Node(string operatorSymbol, double? value, Node left, Node right)
    {
        this.operatorSymbol = operatorSymbol;
        this.value = value;
        this.left = left;
        this.right = right;
    }

    public double Result()
    {
        switch (operatorSymbol)
        {
            case "x":
                return left.Result() * right.Result();
            case "÷":
                return left.Result() / right.Result();
            case "+":
                return left.Result() + right.Result();
            case "-":
                return left.Result() - right.Result();
            default:
                return value.HasValue ? value.Value : 0;
        }
    }

    public override string ToString()
    {
        switch (operatorSymbol)
        {
            case "x":
                return $"({left.ToString()} x {right.ToString()})";
            case "÷":
                return $"({left.ToString()} ÷ {right.ToString()})";
            case "+":
                return $"({left.ToString()} + {right.ToString()})";
            case "-":
                return $"({left.ToString()} - {right.ToString()})";
            default:
                return value.HasValue ? value.Value.ToString() : "0";
        }
    }
}

[TestFixture]
public class NodeTests
{
    [Test]
    public void TestNode()
    {
        // Create the tree structure
        Node tree = new Node(
            "÷",
            null,
            new Node(
                "+",
                null,
                new Node("", 7, null, null),
                new Node(
                    "x",
                    null,
                    new Node("-", null, new Node("", 3, null, null), new Node("", 2, null, null)),
                    new Node("", 5, null, null)
                )
            ),
            new Node("", 6, null, null)
        );

        // Assert that the string representation matches
        Assert.AreEqual("((7 + ((3 - 2) x 5)) ÷ 6)", tree.ToString());

        // Assert that the result is correct
        Assert.AreEqual(2, tree.Result());
    }
}

