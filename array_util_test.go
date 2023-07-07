package templatereq

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicateStringInArray(t *testing.T) {
	character := []string{"a", "b", "b", "a", "c", "d", "c", "f", "e", "f"}
	expect := []string{"a", "b", "c", "d", "f", "e"}

	init := RemoveDuplicateStrInArray(character)

	if !reflect.DeepEqual(expect, init) {
		t.Errorf("got %v, want %v", init, expect)
	}
}
