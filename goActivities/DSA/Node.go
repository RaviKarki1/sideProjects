package main

type Node struct {
	value      string
	linkedNode *Node
}

func NewNode(value string, linkedNode *Node) *Node {
	return &Node{
		value:      value,
		linkedNode: linkedNode,
	}
}

func (n *Node) GetValue() string {
	return n.value
}

func (n *Node) GetLinkedNode() *Node {
	return n.linkedNode
}

func (n *Node) SetLinkedNode(linkedNode *Node) {
	n.linkedNode = linkedNode
}
