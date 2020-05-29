package sss

// Find _first_ longest string from slice
func findLongestSubstring(buf []string) string {
	s := ""
	for _, str := range buf {
		if len(str) > len(s) {
			s = str
		}
	}
	return s
}

// Get longest subsequence from infinite string ssss..., or return `Infinite` otherwise
func GetSubsequence(s string) (string, error) {
	// we have to concatenate (double) input string to build correct trees
	// two consequent strings are enough to realize do we have infinite substring or not
	str := s + s

	// controversial case. But I decided to make it countable (not Infinite) with length 0 (how it is)
	if len(str) == 0 {
		return "", nil
	}

	// build all possible trees from doubled input string
	trees, err := ParseStringToTrees(&str)
	if err != nil {
		return "", err
	}

	var buf []string // buffer for found substrings

	// iterate over all trees and traverse them to get all valid substrings
	for _, t := range trees {
		s := ""
		// get all substrings from current tree
		GetSubstrings(t.Root, &s, &buf)

		// add last string to all others in buffer
		// NOTE: buffer are being filled inside `GetSubstrings` also, see comments in function
		if len(s) > 0 {
			buf = append(buf, s)
		}
	}
	// find the _first_ longest substring
	longest := findLongestSubstring(buf)

	// if length of the longest string is greater or equal of the length of input string, so substring is Infinite
	if len(longest) >= len(s) {
		return "Infinite", nil
	}

	return longest, nil
}
