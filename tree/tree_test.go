package tree

import "testing"

var tr *Tree

func Setup() {
	tr = NewTree()
	// abc[de]f
	// abc
	tr.AddChild("abc", LITERAL)
	// [
	tr.AddChildAndMoveDown("[", LSBRACKET)
	// de
	tr.AddChild("de", LITERAL)
	// ]
	//t.MoveUp()
	tr.ChangeNode("[]", SBRACKETS)
	tr.MoveUp()
	// f
	tr.AddChild("f", LITERAL)
}

func TestGetString(t *testing.T) {
	Setup()

	var s string
	GetString(tr.Root, &s)

	if s != "abc[de]f" {
		t.Errorf("string must be abc[de]f, got %s", s)
	}
}
