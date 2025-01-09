package cli

import (
	"bufio"
	"bytes"
	"testing"
)

func TestGetPlayerMove(t *testing.T) {

	tests := []struct {
		input        string
		expected1    int
		expected2    int
		expected_err error
	}{
		{"3,4", 3, 4, nil},                 // Valid input
		{"10, 20", 10, 20, nil},            // Valid input with spaces
		{"3", 0, 0, ErrUnpack},             // Missing second number
		{"three,4", 0, 0, ErrInvalidInput}, // Invalid first number
		{"", 0, 0, ErrUnpack},              // Empty input
		{"3,4,5", 0, 0, ErrUnpack},         // Too many numbers
		{"3.5,4", 0, 0, ErrInvalidInput},   // Floating point numbers
		{"   ,   ", 0, 0, ErrInvalidInput}, // Only commas
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			add_enter := test.input + "\n"
			inputWriter := bytes.NewBufferString(add_enter)
			reader := bufio.NewReader(inputWriter)
			num1, num2, err := GetPlayerMove(reader)

			if test.expected_err != err {
				t.Errorf("error values was not as expected. wanted %s, got %s", test.expected_err, err)
			}

			if num1 != test.expected1 || num2 != test.expected2 {
				t.Errorf(
					"Failed as numbers don't match. Got %v, wanted %v",
					[2]int{num1, num2},
					[2]int{test.expected1, test.expected2},
				)
			}
		})
	}

}
