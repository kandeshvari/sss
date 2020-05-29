package sss

import (
	"errors"
)

const (
	EMPTY = iota // pure token type

	// pure node types
	ROOT = iota
	BRACKETS
	BRACES
	SBRACKETS

	// shared btw node types and token types
	LITERAL
	LBRACKET
	RBRACKET
	LBRACE
	RBRACE
	LSBRACKET
	RSBRACKET

	// pure token types
	BAD
	EOF
)

type NodeType int

type Node struct {
	Children []*Node
	Parent   *Node
	Type     NodeType
	Value    string
}

func (n *Node) GetValue() (NodeType, string) {
	return n.Type, n.Value
}

func (n *Node) SetValue(typ NodeType, value string) {
	n.Value = value
	n.Type = typ
}

func (n *Node) GetParent() *Node {
	return n.Parent
}

type Tree struct {
	Root        *Node
	CurrentNode *Node
}

func NewTree() *Tree {
	node := &Node{Type: ROOT}
	return &Tree{Root: node, CurrentNode: node}
}

func (t *Tree) AddChild(typ NodeType, value string) *Node {
	node := &Node{Parent: t.CurrentNode, Type: typ, Value: value}
	t.CurrentNode.Children = append(t.CurrentNode.Children, node)
	return node
}

func (t *Tree) AddChildAndMoveOnto(typ NodeType, value string) *Node {
	node := t.AddChild(typ, value)
	t.CurrentNode = node
	return node
}

func (t *Tree) MoveUp() (*Node, error) {
	if t.CurrentNode.Parent == nil {
		return nil, errors.New("already on top")
	}
	t.CurrentNode = t.CurrentNode.Parent
	return t.CurrentNode, nil
}

func (t *Tree) ChangeNode(typ NodeType, value string) *Node {
	t.CurrentNode.SetValue(typ, value)
	return t.CurrentNode
}

func (t *Tree) GetCurrentNodeValues() (NodeType, string) {
	return t.CurrentNode.GetValue()
}

// Get substrings from tree by traversal
func GetSubstrings(n *Node, str *string, buf *[]string) {

	if n == nil {
		return
	}
	for _, node := range n.Children {
		switch node.Type {
		case LITERAL:
			*str = *str + node.Value
		case LSBRACKET, LBRACE, LBRACKET:
			// here we catch alone open brackets and split substring to correct pieces
			*buf = append(*buf, *str)
			*str = ""
			GetSubstrings(node, str, buf)
		case SBRACKETS:
			*str = *str + "["
			GetSubstrings(node, str, buf)
			*str = *str + "]"
		case BRACKETS:
			*str = *str + "("
			GetSubstrings(node, str, buf)
			*str = *str + ")"
		case BRACES:
			*str = *str + "{"
			GetSubstrings(node, str, buf)
			*str = *str + "}"
		}
	}
}
