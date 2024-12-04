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
	fmt.Println("Calculating muls...")
	i := readInput()
    cnt := sumMul(i)
    fmt.Println(cnt)
    cntDo := sumMulDo(i)
    fmt.Println(cntDo)
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

func sumMulDo(s []string) int {
    sum :=0
    active := true
    for i :=0; i < len(s); i++ {
        sum += parseMulDo2(active, s[i])
    }
    return sum
}

func parseMulDo(s string) int {
    a := strings.Split(s, "do()")
    sum := 0
    for i := 0; i < len(a); i++ {
        do := strings.Split(a[i], "don't()")[0]
        sum += parseMul(do)
    }
    return sum
}

func parseMulDo2(active bool, s string) int {
    re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
    parsed := re.FindAll([]byte(s), -1)
    // active := true
    sum := 0
    for _, p := range parsed {
        f := string(p)
        if f == "do()" {
            active = true
        } else if f == "don't()" {
            active = false
        } else {
            if active {
                sum += calcMul(f)
            }
        }
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

