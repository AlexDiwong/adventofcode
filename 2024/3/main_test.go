package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMul(t *testing.T) {
    s:= "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	want := 161
	assert.Equal(t, want, parseMul(s))
}

func TestMulDo(t *testing.T) {
    s:= "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	want := 48
	assert.Equal(t, want, parseMulDo2(s))
}

