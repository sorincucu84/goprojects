package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Escape counts length and compiled length of string
func Escape() {
	file, _ := os.Open("solutions/escape_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lenSum := 0
	memSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lenSum += len(line)
		unquoted, _ := strconv.Unquote(line)
		memSum += len(unquoted)
	}

	fmt.Println("len - mem:", lenSum-memSum)
}

func EscapeV2() {
	file, _ := os.Open("solutions/escape_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lenSum := 0
	encodedLenSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lenSum += len(line)
		quoted := strconv.Quote(line)
		encodedLenSum += len(quoted)
	}

	fmt.Println("encodedLenSum - lenSum:", encodedLenSum-lenSum)
}
