package main

import (
	"testing"
)

type Result struct {
	token Token
}

type Suite struct {
	inputString string
	result      []Result
}

var dataSuite1 = []Suite{
	{"abc[de]f", []Result{
		{Token{Type: LITERAL, Value: "abc"}},
		{Token{Type: LSBRACKET, Value: "["}},
		{Token{Type: LITERAL, Value: "de"}},
		{Token{Type: RSBRACKET, Value: "]"}},
		{Token{Type: LITERAL, Value: "f"}},
		{Token{Type: EOF, Value: ""}},
	}},
	{"[{(ab)}]", []Result{
		{Token{Type: LSBRACKET, Value: "["}},
		{Token{Type: LBRACE, Value: "{"}},
		{Token{Type: LBRACKET, Value: "("}},
		{Token{Type: LITERAL, Value: "ab"}},
		{Token{Type: RBRACKET, Value: ")"}},
		{Token{Type: RBRACE, Value: "}"}},
		{Token{Type: RSBRACKET, Value: "]"}},
		{Token{Type: EOF, Value: ""}},
	}},
}

func TestScanner_Read(ts *testing.T) {
	for _, suite := range dataSuite1 {
		s := NewScanner(&suite.inputString)
		var t = Token{}
		idx := 0
		for t.Type != EOF && t.Type != BAD {
			t = s.Read()
			if t.Type != suite.result[idx].token.Type || t.Value != suite.result[idx].token.Value {
				ts.Errorf("await: t: %s, v: %s; got t: %s, v: %s",
					tokenMap[suite.result[idx].token.Type], suite.result[idx].token.Value,
					tokenMap[t.Type], t.Value)
			}
			idx++
		}
	}
}
