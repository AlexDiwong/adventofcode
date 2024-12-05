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
	fmt.Println("Calculating valid...")
	rules := readInput("5a.txt", "|")
	pages := readInput("5b.txt", ",")
	cnt := countValid(rules, pages)
	fmt.Println(cnt)
	cntFixed := countFixInvalid(rules, pages)
	fmt.Println(cntFixed)
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
	return sum
}

func countFixInvalid(rules, pages [][]string) int {
	sum := 0
    length := len(pages)
	for i, p := range pages {
        fmt.Println("Working on page", i, length)
		if a := seqFixInvalid(rules, p); len(a) > 0 {
			sum += castInt(a[len(a)/2])
		}
	}
	return sum
}

func seqFixInvalid(rules [][]string, page []string) []string {
	for _, r := range rules {
		a := indexOf(page, r[0])
		b := indexOf(page, r[1])
		if a != -1 && b != -1 && a > b {
			c := fixInvalid(rules, page)
			return c
		}
	}
	return make([]string, 0)
}

func fixInvalid(rules [][]string, page []string) []string {
    for i := 0; i < len(page); i++ {
        if seqValid(rules, p) {
            return p
        }
    }
	return make([]string, 0)
}

func seqValid(rules [][]string, page []string) bool {
	for _, r := range rules {
		a := indexOf(page, r[0])
		b := indexOf(page, r[1])
		if a != -1 && b != -1 && a > b {
			return false
		}
	}
	return true
}

func castInt(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func removeIndex(s []string, i int) (string, []string) {
	r := make([]string, 0)
	r = append(r, s[:i]...)
	if i < len(s)-1 {
		r = append(r, s[i+1:]...)
	}
	return s[i], r
}

func indexOf(s []string, t string) int {
	for p, v := range s {
		if v == t {
			return p
		}
	}
	return -1
}
