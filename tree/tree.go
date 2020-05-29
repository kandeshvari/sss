package tree

import (
	"errors"
)

const (
	ROOT = iota
	LITERAL
	LBRACKET
	RBRACKET
	BRACKETS
	LBRACE
	RBRACE
	BRACES
	LSBRACKET
	RSBRACKET
	SBRACKETS
)

type NodeType int

type Node struct {
	Children []*Node
	Parent   *Node
	Type     NodeType
	Value    string
}

func (n *Node) GetValue() (string, NodeType) {
	return n.Value, n.Type
}

func (n *Node) SetValue(value string, typ NodeType) {
	n.Value = value
	n.Type = typ
}

func (n *Node) GetParent() *Node {
	return n.Parent
}

type Tree struct {
	Root        *Node
	S           string
	Len         string
	CurrentNode *Node
}

func NewTree() *Tree {
	node := &Node{Type: ROOT}
	return &Tree{Root: node, CurrentNode: node}
}

func (t *Tree) AddChild(value string, typ NodeType) *Node {
	node := &Node{Parent: t.CurrentNode, Type: typ, Value: value}
	t.CurrentNode.Children = append(t.CurrentNode.Children, node)
	return node
}

func (t *Tree) AddChildAndMoveDown(value string, typ NodeType) *Node {
	node := t.AddChild(value, typ)
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

func (t *Tree) ChangeNode(value string, typ NodeType) *Node {
	t.CurrentNode.SetValue(value, typ)
	return t.CurrentNode
}

func GetString(n *Node, str *string) {

	if n == nil {
		return
	}
	for _, node := range n.Children {
		switch node.Type {
		case LITERAL:
			{
				*str = *str + node.Value
			}
		case SBRACKETS:
			{
				*str = *str + "["
				GetString(node, str)
				*str = *str + "]"
			}
		case BRACKETS:
			{
				*str = *str + "("
				GetString(node, str)
				*str = *str + ")"
			}
		case BRACES:
			{
				*str = *str + "{"
				GetString(node, str)
				*str = *str + "}"
			}
		}
	}
}
