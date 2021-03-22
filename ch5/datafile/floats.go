package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads float64s from lines in a file
func GetFloats(fileName string) ([]float64, error) {
	// got into a mess here
	// should have just returned nil, err when the errors occur, 
	// with this implementation have to account for one error overwriting another
	// and always checking error to see if it should bypass and go to the end
	// and what about closing the file and scanner - errors bypasing that now too-
	// just a mess

	numbers := []float64{} // creates an empty slice- same as make([]float64, 0) or var numbers []float64
	file, err := os.Open(fileName)
	if err == nil {
		i := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			num, err := strconv.ParseFloat((scanner.Text()), 64)
			if err != nil {
				break 
			}
			numbers = append(numbers, num)
			i++ 
		}
		
		if err == nil {
			err = file.Close()
			if err == nil {
				err = scanner.Err()
			}
		}
	}

	// in case of error, set numbers slice to nil
	if err != nil {
		numbers = nil
	}
	
	return numbers, err
}
