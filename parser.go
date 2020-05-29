package main

import (
	"errors"
	"fmt"
	"log"
)

func BuildTrees(str *string) ([]*Tree, error) {
	var trees []*Tree
	scanner := NewScanner(str)

	t := NewTree()
	var tok = Token{}
	for tok.Type != EOF && tok.Type != BAD {
		tok = scanner.Read()
		log.Printf("T: %s, V: %s", tokenMap[tok.Type], tok.Value)

		switch tok.Type {
		case BAD:
			return nil, errors.New(fmt.Sprintf("bad input '%s' at position %d", tok.Value, scanner.pos))
		case LITERAL:
			t.AddChild(LITERAL, tok.Value)
			log.Printf("added child")
		case LSBRACKET, LBRACKET, LBRACE:
			t.AddChildAndMoveOnto(NodeType(tok.Type), tok.Value)
			log.Printf("added child and move onto")
		case RSBRACKET:
			log.Printf("move up")
			typ, val := t.GetCurrentNodeValues()
			if typ != LSBRACKET {
				log.Printf("abnormal string at pos: %d, rune: '%s' (type: %d)", scanner.pos, val, typ)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			t.ChangeNode(SBRACKETS, "[]")
			log.Printf("node changed")
			_, err := t.MoveUp()
			if err != nil {
				log.Printf("broken tree at pos: %d, rune: '%s' (type: %d): can't move up: %s", scanner.pos, val, typ, err)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			log.Printf("move up")
		case RBRACKET:
			log.Printf("move up")
			typ, val := t.GetCurrentNodeValues()
			if typ != LBRACKET {
				log.Printf("abnormal string at pos: %d, rune: '%s' (type: %d)", scanner.pos, val, typ)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			t.ChangeNode(BRACKETS, "()")
			log.Printf("node changed")
			_, err := t.MoveUp()
			if err != nil {
				log.Printf("broken tree at pos: %d, rune: '%s' (type: %d): can't move up: %s", scanner.pos, val, typ, err)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			log.Printf("move up")
		case RBRACE:
			log.Printf("move up")
			typ, val := t.GetCurrentNodeValues()
			if typ != LBRACE {
				log.Printf("abnormal string at pos: %d, rune: '%s' (type: %d)", scanner.pos, val, typ)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			t.ChangeNode(BRACES, "{}")
			log.Printf("node changed")
			_, err := t.MoveUp()
			if err != nil {
				log.Printf("broken tree at pos: %d, rune: '%s' (type: %d): can't move up: %s", scanner.pos, val, typ, err)
				trees = append(trees, t)
				t = NewTree()
				continue
			}
			log.Printf("move up")
		case EOF:
			log.Printf("EOF reached")
			trees = append(trees, t)
			break
		}
	}
	return trees, nil
}
