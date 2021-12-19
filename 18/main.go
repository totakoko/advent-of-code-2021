package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	fmt.Println("# Part 1", part1(input))
	fmt.Println("# Part 2", part2(input))
}

func part1(input []string) int {
	nodes := parseInput(input)
	return AddAll(nodes).GetMagnitude()
}

func part2(input []string) int {
	nodes := parseInput(input)
	maxMagnitude := 0
	for _, nodeA := range nodes {
		for _, nodeB := range nodes {
			if nodeA != nodeB {
				magnitude := nodeA.Copy().Add(nodeB.Copy()).GetMagnitude()
				if magnitude > maxMagnitude {
					maxMagnitude = magnitude
				}

			}
		}
	}
	return maxMagnitude
}

func AddAll(nodes []*Node) *Node {
	finalNode := nodes[0]
	for _, node := range nodes[1:] {
		finalNode = finalNode.Add(node)
	}
	return finalNode
}

type Node struct {
	Parent *Node

	// if node
	Left  *Node
	Right *Node

	// if leaf
	Value int
}

func (node *Node) GetMagnitude() int {
	if node.Left != nil {
		return 3*node.Left.GetMagnitude() + 2*node.Right.GetMagnitude()
	}
	return node.Value
}

func (node *Node) Add(other *Node) *Node {
	parentNode := &Node{
		Left:  node,
		Right: other,
	}
	node.Parent = parentNode
	other.Parent = parentNode
	parentNode.Reduce()
	return parentNode
}

func (node *Node) Reduce() *Node {
	for node.explodes(0) || node.splits() {
	}
	return node
}

func (node *Node) explodes(depth int) bool {
	if depth >= 4 && node.Left != nil && node.Left.Left == nil && node.Right.Left == nil { // node is a pair of two values
		if leftSibling := node.Left.FindLeftSibling(); leftSibling != nil {
			leftSibling.Value += node.Left.Value
		}
		if rightSibling := node.Right.FindRightSibling(); rightSibling != nil {
			rightSibling.Value += node.Right.Value
		}
		node.Left = nil
		node.Right = nil
		node.Value = 0
		return true
	}
	if node.Left != nil {
		return node.Left.explodes(depth+1) || node.Right.explodes(depth+1)
	}
	return false
}

func (node *Node) splits() bool {
	if node.Value >= 10 {
		node.Left = &Node{
			Parent: node,
			Value:  int(math.Floor(float64(node.Value) / 2)),
		}
		node.Right = &Node{
			Parent: node,
			Value:  int(math.Ceil(float64(node.Value) / 2)),
		}
		node.Value = 0
		return true
	}
	if node.Left != nil {
		return node.Left.splits() || node.Right.splits()
	}
	return false
}

func (node *Node) FindLeftSibling() *Node {
	currentNode := node
	for currentNode.Parent.Left == currentNode {
		currentNode = currentNode.Parent
		if currentNode.Parent == nil {
			return nil
		}
	}

	// descend to the rightmost value of the ancestor left branch if we were on a left branch
	targetNode := currentNode.Parent.Left
	for targetNode.Right != nil {
		targetNode = targetNode.Right
	}
	return targetNode
}

func (node *Node) FindRightSibling() *Node {
	currentNode := node
	for currentNode.Parent.Right == currentNode {
		currentNode = currentNode.Parent
		if currentNode.Parent == nil {
			return nil
		}
	}

	// descend to the leftmost value of the ancestor right branch if we were on a right branch
	targetNode := currentNode.Parent.Right
	for targetNode.Left != nil {
		targetNode = targetNode.Left
	}
	return targetNode
}

func (node *Node) Copy() *Node {
	return node.copyInternal(nil)
}

func (node *Node) copyInternal(parent *Node) *Node {
	nodeCopy := &Node{
		Parent: parent,
		Value:  node.Value,
	}
	if node.Left != nil {
		nodeCopy.Left = node.Left.copyInternal(nodeCopy)
		nodeCopy.Right = node.Right.copyInternal(nodeCopy)
	}
	return nodeCopy
}

// Add colors to the brackets
func (node *Node) PrettyPrint(indent int) {
	if node.Left != nil {
		fmt.Print("\u001b[" + strconv.Itoa(31+indent%9) + "m[\u001b[0m")
		node.Left.PrettyPrint(indent + 1)
		fmt.Print("\u001b[" + strconv.Itoa(31+indent%9) + "m,\u001b[0m")
		node.Right.PrettyPrint(indent + 1)
		fmt.Print("\u001b[" + strconv.Itoa(31+indent%9) + "m]\u001b[0m")
	} else {
		fmt.Print(node.Value)
	}
	if indent == 0 {
		fmt.Println()
	}
}

func (node *Node) String() string {
	if node.Left != nil {
		return "[" + node.Left.String() + "," + node.Right.String() + "]"
	}
	return strconv.Itoa(node.Value)
}

func parseInput(input []string) []*Node {
	nodes := []*Node{}
	for _, lineStr := range input {
		nodes = append(nodes, ReadNode(lineStr))
	}
	return nodes
}

func ReadNode(line string) *Node {
	node, _ := readNodeInternal(line)
	return node
}

func readNodeInternal(line string) (*Node, int) {
	if line[0] == '[' {
		leftChild, bytesRead := readNodeInternal(line[1:])
		rightChild, rightChildBytesRead := readNodeInternal(line[2+bytesRead:])
		node := &Node{
			Left:  leftChild,
			Right: rightChild,
			Value: -1,
		}
		leftChild.Parent = node
		rightChild.Parent = node
		return node, bytesRead + rightChildBytesRead + 3 // [ , ]
	} else {
		// only one number
		value, _ := strconv.Atoi(string(line[0]))
		return &Node{
			Value: value,
		}, 1
	}
}
