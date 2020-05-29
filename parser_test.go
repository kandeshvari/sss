package main

type DataSuiteParser struct {
	inputString string
	result      []string
}

var dataSuiteParser1 = []DataSuiteParser{
	{"abcd[ef]g", []string{"abcd[ef]g"}},
	{"abcd[ef]]g", []string{"abcd[ef]", "g"}},
	{"abcd[ef)]g", []string{"abcd", "g"}},
}
