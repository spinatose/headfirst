package input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// Gets a float from user input
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // read what user types up to "Enter" press
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}
