package repeated_dna_sequences

import (
	"reflect"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	result := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	if !reflect.DeepEqual(result, []string{"CCCCCAAAAA", "AAAAACCCCC"}) {
		t.Errorf("Result is incorrect, %v", result)
	}

	result = findRepeatedDnaSequences("AAAAAAAAAAA")
	if !reflect.DeepEqual(result, []string{"AAAAAAAAAA"}) {
		t.Errorf("Result is incorrect, %v", result)
	}
}
