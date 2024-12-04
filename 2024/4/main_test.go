package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalk(t *testing.T) {
	s := "MMMSXXMASM\n"+
        "MSAMXMSMSA\n"+
        "AMXSXMAAMM\n"+
        "MSAMASMSMX\n"+
        "XMASAMXAMM\n"+
        "XXAMMXXAMA\n"+
        "SMSMSASXSS\n"+
        "SAXAMASAAA\n"+
        "MAMMMXMMMM\n"+
        "MXMXAXMASX"
    inp := parse(s)
	want := 18
	assert.Equal(t, want, walk(inp))
}

func TestWalkMas(t *testing.T) {
	s := "MMMSXXMASM\n"+
        "MSAMXMSMSA\n"+
        "AMXSXMAAMM\n"+
        "MSAMASMSMX\n"+
        "XMASAMXAMM\n"+
        "XXAMMXXAMA\n"+
        "SMSMSASXSS\n"+
        "SAXAMASAAA\n"+
        "MAMMMXMMMM\n"+
        "MXMXAXMASX"
    inp := parse(s)
	want := 9
	assert.Equal(t, want, walkMas(inp))
}

func parse(s string) [][]string {
	spl := strings.Split(s, "\n")
	ret := make([][]string, 0)
    for i := 0; i < len(spl); i++ {
        tmp := make([]string, 0)
        for _, v := range spl[i] {
            tmp = append(tmp, string(v))
        }
        ret = append(ret, tmp)
    }
	return ret
}
