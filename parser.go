package sss

import (
	"errors"
	"fmt"
)

// Build all possible trees from input string
func ParseStringToTrees(str *string) ([]Tree, error) {
	var trees []Tree
	scanner := NewScanner(str)

	t := NewTree()
	var token = Token{}
	for token.Type != EOF && token.Type != BAD {
		token = scanner.Read()

		switch token.Type {
		case BAD:
			return nil, errors.New(fmt.Sprintf("bad input '%s' at position %d", token.Value, scanner.pos))
		case LITERAL:
			t.AddChild(LITERAL, token.Value)
		case LSBRACKET, LBRACKET, LBRACE:
			t.AddChildAndMoveOnto(NodeType(token.Type), token.Value)
		case RSBRACKET:
			typ, _ := t.GetCurrentNodeValues()
			if typ != LSBRACKET {
				trees = append(trees, *t)
				t = NewTree()
				continue
			}
			t.ChangeNode(SBRACKETS, "[]")
			t.MoveUp()
		case RBRACKET:
			typ, _ := t.GetCurrentNodeValues()
			if typ != LBRACKET {
				trees = append(trees, *t)
				t = NewTree()
				continue
			}
			t.ChangeNode(BRACKETS, "()")
			t.MoveUp()
		case RBRACE:
			typ, _ := t.GetCurrentNodeValues()
			if typ != LBRACE {
				trees = append(trees, *t)
				t = NewTree()
				continue
			}
			t.ChangeNode(BRACES, "{}")
			t.MoveUp()
		case EOF:
			trees = append(trees, *t)
			break
		}
	}
	return trees, nil
}
