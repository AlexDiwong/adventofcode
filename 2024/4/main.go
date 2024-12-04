package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Calculating muls...")
	i := readInput()
	fmt.Println(i)
}

func readInput() [][]string {
	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l := make([]string, 0)
		for _, v := range line {
			l = append(l, string(v))
		}
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func walk(s [][]string) int {
    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s[i]); j++ {
            if s[i][j] == "X" {
            }
        }
    }
    return 0
}

func xmas(i, j int, s [][]string) {
    if i+3 < len(s) && {
    }
}
