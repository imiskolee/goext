package goext

import "testing"

func Test_ArrayChunk(t *testing.T) {

	example1 := []interface{}{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	}

	chunk1, _ := ArrayChunk(example1, 1)

	chunk2, _ := ArrayChunk(example1, 2)

	chunk3, _ := ArrayChunk(example1, 7)

	chunk4, _ := ArrayChunk(example1, 8)

	if len(chunk1) != 7 {

		t.Error("chunk1 testing failed.")
	}

	if len(chunk2) != 4 {
		t.Error("chunk2 testing failed")
	}

	if len(chunk3) != 1 {
		t.Error("chunk3 testing failed")
	}

	if len(chunk4) != 1 {
		t.Error("chunk4 testing failed")
	}

}
