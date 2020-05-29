package sss

import (
	"testing"
)

var data = [][]string{
	{"", ""}, // controversial case. But I decided to make it countable (not Infinite)
	{"]", ""},
	{"[", ""},
	{"[{(", ""},
	{"[{(}", ""},
	{"[]{]]]]]]]]}[]", "[][]"},
	{"a[{(}b", "ba"},
	{"a[{(}b]", "a"},

	{"[]", "Infinite"},
	{"a", "Infinite"},
	{"]abc[d", "Infinite"},
	{"abc", "Infinite"},
	{"a][][b", "Infinite"},
	{"[{(a)}]", "Infinite"},
	{"a)}][{(", "Infinite"},
	{"a)}][{(a)}][{(", "Infinite"},
	{"}][{(a)", "Infinite"},
	{"][{(a)}", "Infinite"},

	{"a]cccc{ccccc[b", "ccccc[ba]cccc"},
	{"a]cccc{{{{ccccc[b", "ccccc[ba]cccc"},
	{"a]cccc{{}ccccc[b", "{}ccccc[ba]cccc"},
	{"a]cccc}ccccc[b", "ccccc[ba]cccc"},
	{"a]cccc}}ccccc[b", "ccccc[ba]cccc"},
	{"a]cccc{}}ccccc[b", "ccccc[ba]cccc{}"},
	{"]}a})abc[d]}}]])", "abc[d]"},
	{"]aW[c[", "c[]aW"},
	{"ab[Z[", "ab"},
	{"[abc[de]", "abc[de]"},

	{"a[bc", "bca"},
	{"a[bc[", "bc"},
	{")a[bcd}(", "bcd"},
	{"a]b]c", "ca"},
}

var wrongData = []string{
	"]abc4[d",
	"Inf{.}inite",
	"abc[d]_",
	"abc [d]",
	"@",
}

var longestData = [][]string{
	// latest value is the correct result
	{"a", "bb", "ccc", "dddd", "eeee", "fff", "gg", "h", "dddd"},
}

func Test_GetSubsequence(t *testing.T) {
	for _, rec := range data {
		longest, err := GetSubsequence(rec[0])
		if err != nil {
			t.Errorf("can't get subsequence: %s", err)
		}

		if longest != rec[1] {
			t.Errorf("wrong subsequence (%s != %s)", longest, rec[1])
		}
	}

	for _, str := range wrongData {
		_, err := GetSubsequence(str)
		if err == nil {
			t.Errorf("wrong input string treated as valid: %s", str)
		}
	}

}

func Test_findLongestSubstring(t *testing.T) {
	for _, buf := range longestData {
		longest := findLongestSubstring(buf[:len(buf)-2])
		if longest != buf[len(buf)-1] {
			t.Errorf("wrong longest string. await %s, got %s", buf[len(buf)-1], longest)
		}
	}
}
