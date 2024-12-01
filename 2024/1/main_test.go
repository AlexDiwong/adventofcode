package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListDistance(t *testing.T) {
	l := []int{3, 4, 2, 1, 3, 3}
	r := []int{4, 3, 5, 3, 9, 3}
	want := 11
	assert.Equal(t, want, listDistance(l, r))
}

func TestSimilarity(t *testing.T) {
	l := []int{3, 4, 2, 1, 3, 3}
	r := []int{4, 3, 5, 3, 9, 3}
	want := 31
	assert.Equal(t, want, similarity(l, r))
}
