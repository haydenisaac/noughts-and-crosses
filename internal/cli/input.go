package cli

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("Input invalid unable to parse numbers.")
var ErrUnpack = errors.New("Too many numbers parsed")

func GetPlayerMove(reader *bufio.Reader) (int, int, error) {

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Something went wrong %v", err)
	}

	num1, num2, err := parseInput(strings.TrimSpace(input))

	return num1, num2, err
}

func parseInput(input string) (int, int, error) {

	in_numbers := strings.Split(input, ",")

	if len(in_numbers) != 2 {
		return 0, 0, ErrUnpack
	}

	numbers := [2]int{}

	for i, num := range in_numbers {
		conv, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			return 0, 0, ErrInvalidInput
		}

		numbers[i] = conv
	}

	return numbers[0], numbers[1], nil

}
