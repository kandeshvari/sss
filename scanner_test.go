package main

import (
	"testing"
)

type DataSuiteScanner struct {
	inputString string
	result      []Token
}

var dataSuite1 = []DataSuiteScanner{
	{"abc[de]f", []Token{
		{Type: LITERAL, Value: "abc"},
		{Type: LSBRACKET, Value: "["},
		{Type: LITERAL, Value: "de"},
		{Type: RSBRACKET, Value: "]"},
		{Type: LITERAL, Value: "f"},
		{Type: EOF, Value: ""},
	}},
	{"[{(ab)}]", []Token{
		{Type: LSBRACKET, Value: "["},
		{Type: LBRACE, Value: "{"},
		{Type: LBRACKET, Value: "("},
		{Type: LITERAL, Value: "ab"},
		{Type: RBRACKET, Value: ")"},
		{Type: RBRACE, Value: "}"},
		{Type: RSBRACKET, Value: "]"},
		{Type: EOF, Value: ""},
	}},
	{"[{(ab5)}]", []Token{
		{Type: LSBRACKET, Value: "["},
		{Type: LBRACE, Value: "{"},
		{Type: LBRACKET, Value: "("},
		{Type: LITERAL, Value: "ab"},
		{Type: BAD, Value: "5"},
	}},
}

func TestScanner_Read(ts *testing.T) {
	for _, data := range dataSuite1 {
		s := NewScanner(&data.inputString)
		var t = Token{}
		idx := 0
		for t.Type != EOF && t.Type != BAD {
			t = s.Read()
			if t.Type != data.result[idx].Type || t.Value != data.result[idx].Value {
				ts.Errorf("await: t: %s, v: %s; got t: %s, v: %s",
					tokenMap[data.result[idx].Type], data.result[idx].Value,
					tokenMap[t.Type], t.Value)
			}
			idx++
		}
	}
}
