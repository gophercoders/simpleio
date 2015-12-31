// Copyright 2015 Kulawe Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package simpleio provides a simplified interface for keyboard input.The package
provides a simple way to read strings, integers and floating point numbers from
the terminal keyboard. The package is intended to be used by young porgrammers
who are beginning to learn to program using Go.

Notes

The package is not idomatic Go. The package is intended to be used by very
young porgrammers with little experience of programming or Go. The package
is intended to be used as a teaching aid and so the function signatures
have been deliberatly simplified, compared to an idomatic Go version.

The package is not go routine safe. Internally the package relies on a single,
ungarded, global instance of bufio.Scanner, scanning from os.Stdin.
*/
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

// ReadStringFromKeyboard reads a string that a user types in at the keyboard.
// If there is no error the string is returned.
// If the user types just a return, or one or more spaces, or one or or more tabs
// then an empty string, "", is returned and no error message will be printed.
// If there is an error, then the error message is printed to the console and
// and empty string, "", is returned.
func ReadStringFromKeyboard() string {
	s, errorStr := readStringFromKeyboard(scanner)
	for errorStr != "" {
		fmt.Println(errorStr)
		s, errorStr = readStringFromKeyboard(scanner)
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
		s = ""
	}
	// in the case of an error we return the empty string.
	return s, errorStr
}

// ReadNumberFromKeyboard reads an single decimal (base 10) integer number,
// positive or negative, that a
// user types in at the keyboard. Any leading or training white space will be
// stripped before conversion to a number is attempted.
// If there are no errors the number is returned as an int.
//
// If the input cannot be converted to an int, because the input is
// badly formed e.g. contains letters, zero is returned and the message
// 		Sorry I don't think that was a number. Try again...
// is printed to the console.
//
// If the input cannot be converted to an int, because the input is to large
// to be stored in an int type zero, is returned and the message
// 		Sorry that number was too big. Try again...
// is printed to the console.
//
// If the input cannot be converted to an int, because the input is to small
// to be stored in an int type, zero is returned and the message
// 		Sorry that number was too small. Try again...
// is printed to the console.
//
// If the user types just a return, or one or more spaces, or one or more tabs
// then this is considered an error condition and zero is returned and the error
// message
// 		Sorry I don't think that was a number. Try again...
// will be printed.
//
// If the user attempts to type more than one number, separated by one ore more spaces
// or tabs then the function returns zero and the error message
// 		Sorry I don't think that was a number. Try again...
// will be printed.
//
// Internally the function uses a bufio.Scanner (http://golang.org/pkg/bufio/#Scanner). If the scanner aborts early
// for any reason (apart from EOF (End Of File) then the an error string of the form
//		Sorry I could not scan the line. Error: <error-msg>. Try again...
// where <error-msg> will be replaced with the reason why the scanner aborted.
//
// Notes
//
// The function will panic with the string
//		The error type returned by strconv.ParseInt is NOT an *strconv.NumError!. This contradics the ParseInt docs.
// if the error condition returned by strconv.ParseInt is not of type *strconv.NumError.
// The panic will immediatly stop the programs execution if it occurs.
func ReadNumberFromKeyboard() int {
	i, s := readNumberFromKeyboard(scanner)
	for s != "" {
		fmt.Println(s)
		i, s = readNumberFromKeyboard(scanner)
	}
	return i
}

func readNumberFromKeyboard(scanner *bufio.Scanner) (int, string) {
	var errorStr string
	var scannerEmpty = true // assume the scanner is empty until proved otherwise

	// scanner.Scan will return false when the scanner reads EOF. If the reader contains
	// nothing i.e. an empty string was passed to the scanner. This is a problem,
	// because we cannot distinguish between and empty string and the true EOF error.
	// Remember scanner.Err() only returns errors that are not EOF.
	// We need record is scanner.Scan worked at all. We use scannerEmpty for this.
	for scanner.Scan() {
		// scanner.Scan worked so the scanner is not empty. The string we read from the
		// scanner was not "", even if TrimSpace is about to turns it into a ""
		scannerEmpty = false
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
	// This is the normal case - the scanned encountered an error - apart from EOF.
	if err := scanner.Err(); err != nil {
		errorStr = fmt.Sprintf("Sorry I could not scan the line. Error: %v. Try again...", err)
	} else if !scannerEmpty {
		// this is the abnormal case. The scanner read nothing. I nthis case scanner.Scan returned
		// false, and scanner.Err() returns nil because EOF errors are ignored by scanner.Err().
		// The only way to pick up this case is to use our scannerEmpty variable.
		// If this is still true, the for loop never executed. We treat this as
		//bad input and return the same error message.
		errorStr = "Sorry I don't think that was a number. Try again..."
	}
	// in the case of an error we return zero
	return 0, errorStr
}

// ReadDecimalFractionFromKeyboard reads an single decimal (base 10) fractional number,
// positive or negative, that a
// user types in at the keyboard. Any leading or training white space will be
// stripped before conversion to a number is attempted.
// If there are no errors the number is returned as an float64.
//
// If the input cannot be converted to an float64, because the input is
// badly formed e.g. contains letters, 0.0 is returned and the message
// 		Sorry I don't think that was a number. Try again...
// is printed to the console.
//
// If the input cannot be converted to an float64, because the input is to large
// to be stored in a float64 type 0.0, is returned and the message
// 		Sorry that number was too big. Try again...
// is printed to the console.
//
// If the input cannot be converted to an float64, because the input is to small
// to be stored in a float64 type, 0.0 is returned and the message
// 		Sorry that number was too small. Try again...
// is printed to the console.
//
// If the user types just a return, or one or more spaces, or one or more tabs
// then this is considered an error condition and 0.0 is returned and the error
// message
// 		Sorry I don't think that was a number. Try again...
// will be printed.
//
// If the user attempts to type more than one number, separated by one ore more spaces
// or tabs then the function returns 0.0 and the error message
// 		Sorry I don't think that was a number. Try again...
// will be printed.
//
// Internally the function uses a bufio.Scanner (http://golang.org/pkg/bufio/#Scanner). If the scanner aborts early
// for any reason (apart from EOF (End Of File) then the an error string of the form
//		Sorry I could not scan the line. Error: <error-msg>. Try again...
// where <error-msg> will be replaced with the reason why the scanner aborted.
//
// Notes
//
// The function will panic with the string
//		The error type returned by strconv.ParseFloat is NOT an *strconv.NumError!. This contradics the ParseInt docs.
// if the error condition returned by strconv.ParseFloat is not of type *strconv.NumError.
// The panic will immediatly stop the programs execution if it occurs.
func ReadDecimalFractionFromKeyboard() float64 {
	i, s := readDecimalFractionFromKeyboard(scanner)
	for s != "" {
		fmt.Println(s)
		i, s = readDecimalFractionFromKeyboard(scanner)
	}
	return i
}

func readDecimalFractionFromKeyboard(scanner *bufio.Scanner) (float64, string) {
	var errorStr string
	var scannerEmpty = true // assume the scanner is empty until proved otherwise

	// scanner.Scan will return false when the scanner reads EOF. If the reader contains
	// nothing i.e. an empty string was passed to the scanner. This is a problem,
	// because we cannot distinguish between and empty string and the true EOF error.
	// Remember scanner.Err() only returns errors that are not EOF.
	// We need record is scanner.Scan worked at all. We use scannerEmpty for this.
	for scanner.Scan() {
		// scanner.Scan worked so the scanner is not empty. The string we read from the
		// scanner was not "", even if TrimSpace is about to turns it into a ""
		scannerEmpty = false
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
	} else if !scannerEmpty {
		// this is the abnormal case. The scanner read nothing. I nthis case scanner.Scan returned
		// false, and scanner.Err() returns nil because EOF errors are ignored by scanner.Err().
		// The only way to pick up this case is to use our scannerEmpty variable.
		// If this is still true, the for loop never executed. We treat this as
		// bad input and return the same error message.
		errorStr = "Sorry I don't think that was a number. Try again..."
	}
	// in the case of an error we return 0.0
	return 0.0, errorStr
}
