package simpleio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

func ReadStringFromKeyboard() string {
	s, errorStr := readStringFromKeyboard(scanner)
	if errorStr != "" {
		fmt.Println(errorStr)
	}
	return s
}

func readStringFromKeyboard(scanner *bufio.Scanner) (string, string) {
	var s string
	var errorStr string
	for scanner.Scan() {
		s = strings.TrimSpace(scanner.Text())
		// there was no error sowe can return the string here. This
		// limits the function to the first line only.
		return s, errorStr
	}
	// The scanner stoped. Why?
	if err := scanner.Err(); err != nil {
		errorStr = fmt.Sprintf("Sorry I could not scan the line. Error: %v. Try again...", err)
	}
	// in the case of an error we return the empty string.
	return s, errorStr
}

func ReadNumberFromKeyboard() int {
	i, s := readNumberFromKeyboard(scanner)
	if s != "" {
		fmt.Println(s)
	}
	return i
}

func readNumberFromKeyboard(scanner *bufio.Scanner) (int, string) {
	var errorStr string
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			if err, ok := err.(*strconv.NumError); ok {
				switch err.Err {
				case strconv.ErrSyntax:
					errorStr = "Sorry I don't think that was a number. Try again..."
				case strconv.ErrRange:
					if i > 0 {
						errorStr = "Sorry that number was too big. Try again..."
						i = 0
					} else {
						errorStr = "Sorry that number was too small. Try again..."
						i = 0
					}
				}
			} else {
				panic("The error type returned by strconv.ParseInt is NOT an *strconv.NumError!. This contradics the ParseInt docs.")
			}
		}
		// return the first number we find.
		// i always has a value at this point See the docs for ParseInt
		// http://godoc.org/strconv#ParseInt
		return int(i), errorStr
	}
	// The scanner stopped. Why?
	if err := scanner.Err(); err != nil {
		errorStr = fmt.Sprintf("Sorry I could not scan the line. Error: %v. Try again...", err)
	}
	// in the case of an error we return zero
	return 0, errorStr
}

func ReadDecimalFractionFromKeyboard() float64 {
	i, s := readDecimalFractionFromKeyboard(scanner)
	if s != "" {
		fmt.Println(s)
	}
	return i
}

func readDecimalFractionFromKeyboard(scanner *bufio.Scanner) (float64, string) {
	var errorStr string
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			if err, ok := err.(*strconv.NumError); ok {
				switch err.Err {
				case strconv.ErrSyntax:
					errorStr = "Sorry I don't think that was a number. Try again..."
				case strconv.ErrRange:
					if f > 0 {
						errorStr = "Sorry that number was too big. Try again..."
					} else {
						errorStr = "Sorry that number was too small. Try again..."
					}
					f = 0.0
				}
			} else {
				panic("The error type returned by strconv.ParseParse is NOT an *strconv.NumError!. This contradics the ParseFloat docs.")
			}
		}
		// return the first number we find.
		// f always has a value at this point See the docs for ParseFloat
		// http://godoc.org/strconv#ParseFloat
		return f, errorStr
	}
	// The scanner stoped. Why?
	if err := scanner.Err(); err != nil {
		errorStr = fmt.Sprintf("Sorry I could not scan the line. Error: %v. Try again...", err)
	}
	// in the case of an error we return 0.0
	return 0.0, errorStr
}
