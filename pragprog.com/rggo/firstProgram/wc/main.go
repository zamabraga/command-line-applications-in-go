package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Calling the count function to count the number of words received from the
	// standard Input and print it out

	lines := flag.Bool("l", false, "Count lines")
	b := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *b))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is user to read text from a REader (such as files)

	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words (default is split by lines)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Define a counter
	wc := 0

	//For ever word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	//Return the total
	return wc

}
