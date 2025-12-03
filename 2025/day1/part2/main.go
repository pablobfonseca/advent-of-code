package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	StartPosition = 50
	DialSize      = 100
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dialPosition := StartPosition
	zeroCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		direction, distance, err := parseDirection(line)
		if err != nil {
			log.Fatalf("Failed to parse input: %v", err)
		}

		step := 1
		if direction == "L" {
			step = -1
		}

		for i := 1; i <= distance; i++ {
			dialPosition = ((dialPosition+step)%DialSize + DialSize) % DialSize
			if dialPosition == 0 {
				zeroCount++
			}
		}
	}

	fmt.Printf("Zero Count: %d\n", zeroCount)
}

func parseDirection(line string) (direction string, value int, err error) {
	direction = strings.ToUpper(string(line[0]))
	valString := line[1:]

	value, err = strconv.Atoi(valString)
	if err != nil {
		return "", 0, err
	}

	return
}
