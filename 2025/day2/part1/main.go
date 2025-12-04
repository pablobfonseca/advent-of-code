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
				mid := len(str) / 2
				left, right := str[:mid], str[mid:]

				if left == right {
					sum += i
				}
			}
		}
	}

	fmt.Println("Sum", sum)
}
