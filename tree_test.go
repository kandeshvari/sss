package main

import (
	"testing"
)

var tr *Tree

func Setup() {
	tr = NewTree()
	// abc[de]f
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

func TestGetString(t *testing.T) {
	Setup()

	var s string
	GetString(tr.Root, &s)

	if s != "abc[de]f" {
		t.Errorf("string must be abc[de]f, got %s", s)
	}
}
