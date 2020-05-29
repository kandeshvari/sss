package sss

import (
	"testing"
)

var tr *Tree

// this setup indirectly covers all common tree routines
func Setup() {
	// setup for `abc[de]f`
	tr = NewTree()
	// abc
	tr.AddChild(LITERAL, "abc")
	// [
	tr.AddChildAndMoveOnto(LSBRACKET, "[")
	// de
	tr.AddChild(LITERAL, "de")
	// ]
	tr.ChangeNode(SBRACKETS, "[]")
	tr.MoveUp()
	// f
	tr.AddChild(LITERAL, "f")
}

func Test_CommonCreateTreeRoutines(t *testing.T) {
	Setup()

	if tr.Root.Children[0].Type != LITERAL && tr.Root.Children[0].Value != "abc" {
		t.Errorf("node must be (LITERAL, 'abc'), got (%#v)", tr.Root.Children[0])
	}

	if tr.Root.Children[1].Type != SBRACKETS && tr.Root.Children[0].Value != "[]" {
		t.Errorf("node must be (SBRACKETS, '[]'), got (%#v)", tr.Root.Children[1])
	}

	if tr.Root.Children[1].Children[0].Type != LITERAL && tr.Root.Children[1].Children[0].Value != "de" {
		t.Errorf("node must be (LITERAL, 'de'), got (%#v)", tr.Root.Children[1].Children[0])
	}

	if tr.Root.Children[2].Type != LITERAL && tr.Root.Children[2].Value != "f" {
		t.Errorf("node must be (LITERAL, 'f'), got (%#v)", tr.Root.Children[2])
	}

	if tr.CurrentNode != tr.Root {
		t.Errorf("current node must reference root node, got (%#v)", tr.CurrentNode)
	}

}

func Test_GetString(t *testing.T) {
	Setup()

	var s string
	var buf []string
	GetSubstrings(tr.Root, &s, &buf)

	if s != "abc[de]f" {
		t.Errorf("string must be abc[de]f, got %s", s)
	}
}
