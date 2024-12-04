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
	c := walk(i)
	fmt.Println(c)
	c2 := walkMas(i)
	fmt.Println(c2)
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
	sum := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == "X" {
				sum += xmas(i, j, s)
			}
		}
	}
	return sum
}

func xmas(i, j int, s [][]string) int {
	x := "MAS"
	cnt := 0
	if i+3 < len(s) {
		if s[i+1][j]+s[i+2][j]+s[i+3][j] == x {
			cnt++
		}
		if j+3 < len(s[i]) && s[i+1][j+1]+s[i+2][j+2]+s[i+3][j+3] == x {
			cnt++
		}
		if j-3 >= 0 && s[i+1][j-1]+s[i+2][j-2]+s[i+3][j-3] == x {
			cnt++
		}
	}
	if i-3 >= 0 {
		if s[i-1][j]+s[i-2][j]+s[i-3][j] == x {
			cnt++
		}
		if j+3 < len(s[i]) && s[i-1][j+1]+s[i-2][j+2]+s[i-3][j+3] == x {
			cnt++
		}
		if j-3 >= 0 && s[i-1][j-1]+s[i-2][j-2]+s[i-3][j-3] == x {
			cnt++
		}
	}
	if j+3 < len(s[i]) && s[i][j+1]+s[i][j+2]+s[i][j+3] == x {
		cnt++
	}
	if j-3 >= 0 && s[i][j-1]+s[i][j-2]+s[i][j-3] == x {
		cnt++
	}

	return cnt
}

func walkMas(s [][]string) int {
	sum := 0
	for i := 1; i < len(s)-1; i++ {
		for j := 1; j < len(s[i])-1; j++ {
			if s[i][j] == "A" && masX(i, j, s) {
				sum++
			}
		}
	}
	return sum
}

func masX(i, j int, s [][]string) bool {
	a1 := s[i+1][j+1]
	a2 := s[i-1][j-1]
	b1 := s[i-1][j+1]
	b2 := s[i+1][j-1]
    c1 := a1 == "M" && a2 == "S"
    c2 := a1 == "S" && a2 == "M"
    c3 := b1 == "M" && b2 == "S"
    c4 := b1 == "S" && b2 == "M"
	return (c1 || c2) && (c3 || c4)
}
