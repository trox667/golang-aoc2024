package tools

import "testing"

func TestReader(t *testing.T) {
	data := ReadInput("../inputs/test")
	if len(data) != 4 {
		t.Error("Not all lines parsed")
	}
	if data[0] != "A" || data[1] != "B" || data[2] != "C" || data[3] != "D" {
		t.Error("File not correctly parsed")
	}
}
