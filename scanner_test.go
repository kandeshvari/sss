package main

import (
	"testing"
)

type Result struct {
	token Token
}

type DataSuite struct {
	inputString string
	result      []Result
}

var dataSuite1 = []DataSuite{
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
	{"[{(ab5)}]", []Result{
		{Token{Type: LSBRACKET, Value: "["}},
		{Token{Type: LBRACE, Value: "{"}},
		{Token{Type: LBRACKET, Value: "("}},
		{Token{Type: LITERAL, Value: "ab"}},
		{Token{Type: BAD, Value: ""}},
	}},
}

func TestScanner_Read(ts *testing.T) {
	for _, data := range dataSuite1 {
		s := NewScanner(&data.inputString)
		var t = Token{}
		idx := 0
		for t.Type != EOF && t.Type != BAD {
			t = s.Read()
			if t.Type != data.result[idx].token.Type || t.Value != data.result[idx].token.Value {
				ts.Errorf("await: t: %s, v: %s; got t: %s, v: %s",
					tokenMap[data.result[idx].token.Type], data.result[idx].token.Value,
					tokenMap[t.Type], t.Value)
			}
			idx++
		}
	}
}
