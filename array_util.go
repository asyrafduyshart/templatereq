package templatereq

import "sort"

// a function that will remove duplicate string array
// example I : a, b, b, c, d, e, f, f, g -> a, b, c, d, e, f, g
// example II : b, c, f, d, d, a, b, e -> b, c, f, d, e

type StringsWithPos struct {
	// slice of strings
	strs []string
	// slice of integers that represent their original positions
	pos []int
}

// implement sort.Interface interface for StringsWithPos type

// Len returns the length of strs slice
func (s StringsWithPos) Len() int {
	return len(s.strs)
}

// Less compares two elements by their positions in pos slice
func (s StringsWithPos) Less(i, j int) bool {
	return s.pos[i] < s.pos[j]
}

// Swap swaps two elements in both strs and pos slices
func (s StringsWithPos) Swap(i, j int) {
	s.strs[i], s.strs[j] = s.strs[j], s.strs[i]
	s.pos[i], s.pos[j] = s.pos[j], s.pos[i]
}

// RemoveDuplicates removes duplicate strings from a slice and keeps their original index
func RemoveDuplicateStrInArray(strSlice []string) []string {
	// create a map to store the seen strings and their positions
	seen := make(map[string]int)
	// create an instance of StringsWithPos type with empty slices
	result := StringsWithPos{[]string{}, []int{}}
	for i, s := range strSlice {
		// if the string is not in the map, add it with its position to result and seen
		if _, ok := seen[s]; !ok {
			result.strs = append(result.strs, s)
			result.pos = append(result.pos, i)
			seen[s] = i
		}
	}
	// sort the result by the positions
	sort.Sort(result)
	// return the result slice of strings
	return result.strs
}
