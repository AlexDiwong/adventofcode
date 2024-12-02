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
	fmt.Println("Calculating safe reports...")
	i := readInput()
	safe := calcSafe(i)
	safeWithDamp := calcSafeWithDamp(i)
	fmt.Println(safe)
	fmt.Println(safeWithDamp)
}

func readInput() [][]int {
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		input = append(input, trimCast(spl))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func trimCast(s []string) []int {
	inp := make([]int, 0)
	for _, v := range s {
		s2 := strings.Trim(v, " ")
		if i, err := strconv.Atoi(s2); err == nil {
			inp = append(inp, i)
		}
	}
	return inp
}

func calcSafe(l [][]int) int {
	cnt := 0
	for _, i := range l {
		if isSafe(i) {
			cnt++
		}
	}
	return cnt
}

func calcSafeWithDamp(l [][]int) int {
	cnt := 0
	for _, i := range l {
		if isSafeWithDamp(i) {
			cnt++
		}
	}
	return cnt
}

func isSafe(i []int) bool {
	if len(i) < 2 {
		return false
	}
	inc := false
	if i[0] < i[1] {
		inc = true
	}
	old := i[0]
	for j := 1; j < len(i); j++ {
		dst := i[j] - old
		if inc {
			if dst < 1 || dst > 3 {
                return false
			}
		} else {
			if dst > -1 || dst < -3 {
                return false
			}
		}
		old = i[j]
	}
	return true
}

func isSafeWithDamp(i []int) bool {
	damp := true
	if len(i) < 2 {
		return false
	}
	inc := false
	if i[0] < i[1] {
		inc = true
	}
	old := i[0]
	for j := 1; j < len(i); j++ {
		dst := i[j] - old
		if inc {
			if dst < 1 || dst > 3 {
				if damp {
					damp = false
					continue
				} else {
					return false
				}
			}
		} else {
			if dst > -1 || dst < -3 {
				if damp {
					damp = false
					continue
				} else {
					return false
				}
			}
		}
		old = i[j]
	}
	return true
}
