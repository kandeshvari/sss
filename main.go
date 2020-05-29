package main

import (
	"fmt"
)

func main() {
	sym := 'a'
	if isLetter(sym) {
		fmt.Printf("ok\n")
	} else {
		fmt.Printf("bad: %s\n", string('a'))
	}

	str := "ab[de]}}5(ad){f}"
	scanner := NewScanner(&str)

	var t = Token{}
	for t.Type != EOF && t.Type != BAD {
		t = scanner.Read()
		fmt.Printf("T: %s, V: %s\n", tokenMap[t.Type], t.Value)
	}

	//
	//
	//
	//t := tree.NewTree()
	//// abc[de]f
	//// abc
	//t.AddChild("abc", tree.LITERAL)
	//// [
	//t.AddChildAndMoveDown("[", tree.LSBRACKET)
	//// de
	//t.AddChild("de", tree.LITERAL)
	//// ]
	////t.MoveUp()
	//t.ChangeNode("[]", tree.SBRACKETS)
	//t.MoveUp()
	//// f
	//t.AddChild("f", tree.LITERAL)
	//
	//var s string
	//tree.GetString(t.Root, &s)
	//
	//fmt.Printf(">>> %s", s)

}
