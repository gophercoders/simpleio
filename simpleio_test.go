package simpleio

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

type failingReader struct{}

func (failingReader) Read(p []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

var stringTestResults = []struct {
	s             string
	expectedStr   string
	expectedError string
}{
	{"", "", ""},
	{" ", "", ""},
	{"\n", "", ""},
	{"123", "123", ""},
	{"-123", "-123", ""},
	{"    abc\n", "abc", ""},
	{"xyz    \n", "xyz", ""},
	{"    ijk    \n", "ijk", ""},
	{"Hello World!", "Hello World!", ""},
}

func TestReadStringFromKeyboard(t *testing.T) {
	for _, td := range stringTestResults {
		sr := strings.NewReader(td.s)
		scanner := bufio.NewScanner(sr)
		readString, errorStr := readStringFromKeyboard(scanner)
		if readString != td.expectedStr {
			t.Fatalf("expected \"%s\" but got \"%s\"\n", td.expectedStr, readString)
		}
		if errorStr != td.expectedError {
			t.Fatalf("expected \"%s\" but got \"%s\"\n", td.expectedStr, readString)
		}
	}
}

func TestReadStringFromKeyboardFailScanner(t *testing.T) {
	scanner := bufio.NewScanner(failingReader{})
	s, errorStr := readStringFromKeyboard(scanner)
	expectedStr := ""
	if s != expectedStr {
		t.Fatalf("Did not get expected string. Expected \"%s\" but got \"%s\"\n", expectedStr, s)
	}
	expectedError := "Sorry I could not scan the line. Error: unexpected EOF. Try again..."
	if errorStr != expectedError {
		t.Fatalf("Did not get expected error string. Expected \"%s\" but got \"%s\"\n", expectedError, errorStr)
	}
}

var numberTestResults = []struct {
	s              string
	expectedNumber int
	expectedError  string
}{
	{"123", 123, ""},
	{"123 456", 0, "Sorry I don't think that was a number. Try again..."},
	{"-123", -123, ""},
	{"    234\n", 234, ""},
	{"345    \n", 345, ""},
	{"    456    \n", 456, ""},
	{"123x", 0, "Sorry I don't think that was a number. Try again..."},
	{"abc", 0, "Sorry I don't think that was a number. Try again..."},
	{"", 0, ""},
	{" ", 0, "Sorry I don't think that was a number. Try again..."},
	{"\n", 0, "Sorry I don't think that was a number. Try again..."},
	{"1234456778909876543211234567890909876654332123434556787890", 0, "Sorry that number was too big. Try again..."},
	{"-1234456778909876543211234567890909876654332123434556787890", 0, "Sorry that number was too small. Try again..."},
}

func TestReadNumberFromKeyboard(t *testing.T) {
	for _, td := range numberTestResults {
		sr := strings.NewReader(td.s)
		scanner := bufio.NewScanner(sr)
		readNumber, errorStr := readNumberFromKeyboard(scanner)
		if readNumber != td.expectedNumber {
			t.Fatalf("Did not get expected number. Expected \"%d\" but got \"%d\"\n", td.expectedNumber, readNumber)
		}
		if errorStr != td.expectedError {
			t.Fatalf("Did not get expected error string. Expected \"%s\" but got \"%s\"\n", td.expectedError, errorStr)
		}
	}
}

func TestReadNumberFromKeyboardFailScanner(t *testing.T) {
	scanner := bufio.NewScanner(failingReader{})
	readNumber, errorStr := readNumberFromKeyboard(scanner)
	expectedNumber := 0.0
	if readNumber != 0.0 {
		t.Fatalf("Did not get expected number. Expected \"%f\" but got \"%f\"\n", expectedNumber, readNumber)
	}
	expectedError := "Sorry I could not scan the line. Error: unexpected EOF. Try again..."
	if errorStr != expectedError {
		t.Fatalf("Did not get expected error string. Expected \"%s\" but got \"%s\"\n", expectedError, errorStr)
	}
}

var decimalFractionTestResults = []struct {
	s              string
	expectedNumber float64
	expectedError  string
}{
	{"3.14", 3.14, ""},
	{"-3.14", -3.14, ""},
	{"3.14 2.71828", 0.0, "Sorry I don't think that was a number. Try again..."},
	{"    2.71828\n", 2.71828, ""},
	{"1.4142    \n", 1.4142, ""},
	{"    0.69314    \n", 0.69314, ""},
	{"123.x", 0, "Sorry I don't think that was a number. Try again..."},
	{"abc.123", 0, "Sorry I don't think that was a number. Try again..."},
	{"", 0, ""},
	{" ", 0, "Sorry I don't think that was a number. Try again..."},
	{"\n", 0, "Sorry I don't think that was a number. Try again..."},
	{"123445677890987654321123456789090987665433212343455678789012344567789098765432112345678909098766543321234345567878901234456778909876543211234567890909876654332123434556787890123445677890987654321123456789090987665433212343455678789012344567789098765432112345678909098766543321234345567878901234456778909876543211234567890909876654332123434556787890.1234353", 0.0, "Sorry that number was too big. Try again..."},
	{"-123445677890987654321123456789090987665433212343455678789012344567789098765432112345678909098766543321234345567878901234456778909876543211234567890909876654332123434556787890123445677890987654321123456789090987665433212343455678789012344567789098765432112345678909098766543321234345567878901234456778909876543211234567890909876654332123434556787890.211324325", 0.0, "Sorry that number was too small. Try again..."},
}

func TestReadDecimalFractionFromKeyboard(t *testing.T) {
	for _, td := range decimalFractionTestResults {
		sr := strings.NewReader(td.s)
		scanner := bufio.NewScanner(sr)
		readNumber, errorStr := readDecimalFractionFromKeyboard(scanner)
		if readNumber != td.expectedNumber {
			t.Fatalf("Did not get expected number. Expected \"%f\" but got \"%f\"\n", td.expectedNumber, readNumber)
		}
		if errorStr != td.expectedError {
			t.Fatalf("Did not get expected error string. Expected \"%s\" but got \"%s\"\n", td.expectedError, errorStr)
		}
	}
}

func TestReadDecimalFractionFromKeyboardFailScanner(t *testing.T) {
	scanner := bufio.NewScanner(failingReader{})
	readNumber, errorStr := readDecimalFractionFromKeyboard(scanner)
	expectedNumber := 0.0
	if readNumber != 0.0 {
		t.Fatalf("Did not get expected number. Expected \"%f\" but got \"%f\"\n", expectedNumber, readNumber)
	}
	expectedError := "Sorry I could not scan the line. Error: unexpected EOF. Try again..."
	if errorStr != expectedError {
		t.Fatalf("Did not get expected error string. Expected \"%s\" but got \"%s\"\n", expectedError, errorStr)
	}
}
