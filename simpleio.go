package simpleio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func ReadStringFromKeyboard() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func ReadNumberFromKeyboard() int {
	var s string
	ok := false
	for !ok {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSpace(s)
		for _, c := range s {
			if c >= '0' && c <= '9' {
				ok = true
			} else {
				ok = false
				break
			}
		}
		if !ok {
			fmt.Println("Sorry I don't think that was a number. Try again...")
		}
	}

	n, _ := strconv.Atoi(s)
	return n

}
