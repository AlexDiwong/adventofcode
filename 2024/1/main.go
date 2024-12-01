package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
    fmt.Println("Calculating distance...")
    l, r := readInput()
    sim := similarity(l, r)
    dist := listDistance(l, r)
    fmt.Println(sim)
    fmt.Println(dist)
}

func readInput() ([]int, []int) {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    l := make([]int, 0)
    r := make([]int, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        spl := strings.Split(line, "   ")
        l = append(l, trimCast(spl[0]))
        r = append(r, trimCast(spl[1]))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return l, r
}

func trimCast(s string) int {
    s = strings.Trim(s, " ")
    if i, err := strconv.Atoi(s); err == nil {
        return i
    }
    return 0
}


func listDistance(l, r []int) int {
    sort.Ints(l)
    sort.Ints(r)
    dist := 0
    for i := 0; i < len(l); i++ {
        dist += int(math.Abs(float64(r[i] - l[i])))
    }
    return dist
}


func similarity(l, r []int) int {
    sum := 0
    occ := agg(r)
    for _, v := range l {
        sum += v*occ[v]
    }
    return sum
}

func agg(r []int) map[int]int {
    res := make(map[int]int,0)
    for _, v := range r {
        res[v] += 1
    }
    return res
}
