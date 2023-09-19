package main

// Import required packages
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// setScanner repositions the file scanner to the appropriate position based on the ASCII character passed.
func setScanner(file *os.File, offsetMap map[int]int64, sign rune) {
	// Calculate the starting line for the given ASCII character in the file.
	calc_sign_start := (int(sign)-31)*9 - 8
	// Seek to the calculated position in the file.
	file.Seek(offsetMap[calc_sign_start+1], 0)
}

// printAscii prints the ASCII representation of the given text.
func printAscii(file *os.File, offsetMap map[int]int64, text string) {
	// Create a map to hold each line of the ASCII art.
	line := make(map[int]string)

	// Loop over each rune in the text string.
	for _, run := range text {
		// Initialize a new scanner.
		scanner := bufio.NewScanner(file)
		// Set the scanner's position based on the current rune.
		setScanner(file, offsetMap, run)
		// Scan and add the corresponding ASCII lines to the map.
		for i := 0; i < 9; i++ {
			scanner.Scan()
			line[i] += scanner.Text()
		}
	}
	// Print the ASCII lines.
	for i := 0; i < 8; i++ {
		fmt.Println(line[i])
	}
}

// check is a simple utility function to check for errors.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// lol
// main is the entry point of the program.
func main() {
	// Retrieve command-line arguments and join them into a single string.
	text := os.Args[1]

	// text := strings.Join(args, " ")

	// Check if the input text is just a newline or empty.
	if text == "\\n" {
		fmt.Println()
		return
	} else if text == "" {
		return
	}

	// Open the standard.txt file.
	flag := "standard.txt"
	if len(os.Args) > 2 {
		flag = os.Args[2]

		if !strings.Contains(flag, ".txt") {
			flag += ".txt"
		}
	}
	file, err := os.Open(flag)
	check(err)
	defer file.Close()

	// Initialize a scanner to read the file.
	scanner := bufio.NewScanner(file)

	// Initialize byte offset and line number variables.
	byteOffset := int64(0)
	lineNumber := 1
	// Create a map to hold line numbers and their corresponding byte offsets.
	offsetMap := make(map[int]int64)

	// Read the entire file line by line, populating the offsetMap.
	for scanner.Scan() {
		line := scanner.Text()
		offsetMap[lineNumber] = byteOffset
		byteOffset += int64(len(line) + 1)
		lineNumber++
	}

	// Split the input text by newline characters.
	result := regexp.MustCompile(`\\n`).Split(text, -1)
	for _, s := range result {
		if s == "" {
			fmt.Println()
		} else {
			// Print the ASCII representation of each segment.
			printAscii(file, offsetMap, s)
		}
	}

	// Check for scanning errors.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
