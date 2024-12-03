package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculating safe reports...")
	i := readInput()
    fmt.Println(i[0])
    cnt := sumMul(i)
    fmt.Println(cnt)
}

func readInput() []string {
	file, err := os.Open("3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
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

func sumMul(s []string) int {
    sum :=0
    for i :=0; i < len(s); i++ {
        sum += parseMul(s[i])
    }
    return sum
}

func parseMul(s string) int {
    re := regexp.MustCompile(`mul\(\d+,\d+\)`)
    parsed := re.FindAll([]byte(s),-1)
    sum := 0
    for _, p := range parsed {
        sum += calcMul(string(p))
    }
    return sum
}

func calcMul(s string) int {
    re := regexp.MustCompile(`\d+`)
    nums := re.FindAll([]byte(s), -1)
    if len(nums) > 2 {
        fmt.Println("too many numbers!", nums)
        return 0
    }
    a := atoi(string(nums[0]))
    b := atoi(string(nums[1]))
    return a*b
}

func atoi(s string) int {
    if i, err := strconv.Atoi(s); err == nil {
		return i
	}
    return 0
}

