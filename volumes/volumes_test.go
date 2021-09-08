package main

import (
	"testing"
	"unicode/utf16"
)

func TestUTF16BufferToStrings(t *testing.T) {
	tests := []struct {
		name     string
		buffer   []uint16
		expected []string
	}{
		{"nil", nil, nil},
		{"zero", []uint16{0}, []string{}},
		{"one", makeUTF16Buffer("foo"), []string{"foo"}},
		{"two", makeUTF16Buffer("foo", "bar"), []string{"foo", "bar"}},
		{"three", makeUTF16Buffer("foo", "bar", "baz"), []string{"foo", "bar", "baz"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := utf16BufferToStrings(test.buffer)
			if a, b := len(result), len(test.expected); a != b {
				t.Fatalf("lengths do not match. got=%d, expected=%d", a, b)
			}
			for i, a := range result {
				b := test.expected[i]
				if a != b {
					t.Fatalf("string[%d] does not match. got=%s, expected=%s", i, a, b)
				}
			}
		})
	}
}

func makeUTF16Buffer(strs ...string) []uint16 {
	var buffer []uint16
	for _, str := range strs {
		enc := utf16.Encode([]rune(str))
		enc = append(enc, 0)
		buffer = append(buffer, enc...)
	}
	return append(buffer, 0)
}
