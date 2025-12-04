package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for idRange := range strings.SplitSeq(line, ",") {
			ids := strings.Split(idRange, "-")
			firstId, err := strconv.Atoi(ids[0])
			if err != nil {
				log.Fatalf("Error converting number: %v", err)
			}
			secondId, err := strconv.Atoi(ids[1])
			if err != nil {
				log.Fatalf("Error converting number: %v", err)
			}

			for i := firstId; i <= secondId; i++ {
				str := strconv.Itoa(i)

				if isRepeatingPattern(str) {
					sum += i
				}
			}
		}
	}

	fmt.Println("Sum", sum)
}

func isRepeatingPattern(s string) bool {
	for patternLen := 1; patternLen <= len(s)/2; patternLen++ {
		if len(s)%patternLen == 0 {
			pattern := s[:patternLen]
			repetitions := len(s) / patternLen

			if repetitions >= 2 {
				if isPatterRepeated(s, pattern) {
					return true
				}
			}
		}
	}

	return false
}

func isPatterRepeated(s, pattern string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != pattern[i%len(pattern)] {
			return false
		}
	}
	return true
}
