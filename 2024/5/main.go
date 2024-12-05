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
	fmt.Println("Calculating muls...")
	rules := readInput("5a.txt", "|")
	pages := readInput("5b.txt", ",")
	fmt.Println(rules)
	fmt.Println(pages)
}

func readInput(f, div string) [][]string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, div)
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func countValid(rules, pages [][]string) int {
    sum := 0 
    for _, p := range pages {
        if seqValid(rules, p) {
            sum += castInt(p[len(p)/2])
        }
    }
	return 0
}

func castInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func seqValid(rules [][]string, page []string) bool {
    return true
}
