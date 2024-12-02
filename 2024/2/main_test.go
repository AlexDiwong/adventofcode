package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafe(t *testing.T) {
	l := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	want := 2
	assert.Equal(t, want, calcSafe(l))
}

func TestSafeWithDamp(t *testing.T) {
	l := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	want := 4
	assert.Equal(t, want, calcSafeWithDamp(l))
}

func TestSafeWithDamp2(t *testing.T) {
	l := [][]int{
		{1, 2, 5, 4, 8},
	}
	want := 1
	assert.Equal(t, want, calcSafeWithDamp(l))
}
