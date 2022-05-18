package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func timeConversion(s string) string {
	var chars = []byte(s)

	var hour, status string

	for index, char := range chars {
		if index == 0 || index == 1 {
			hour += string(char)
		}

		if index == len(chars)-1 || index == len(chars)-2 {
			status += string(char)
		}
	}

	hourInt, err := strconv.Atoi(hour)

	checkError(err)

	if status == "AM" {
		if hourInt == 12 {
			hourInt = 0
		}
	} else {
		if hourInt == 12 {
			hourInt = 12
		} else {
			hourInt += 12
		}
	}

	if hourInt < 10 {
		hour = "0"
		hour += strconv.Itoa(hourInt)
	} else {
		hour = ""
		hour += strconv.Itoa(hourInt)
	}

	chars = chars[2 : len(chars)-2]

	hour += string(chars)

	return hour
}
