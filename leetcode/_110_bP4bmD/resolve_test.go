package _110_bP4bmD

import (
	"log"
	"testing"
)

func TestRoute1(t *testing.T) {
	graph := [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}
	res := allPathsSourceTarget(graph)
	log.Printf("res: %v", res)
}

func TestRoute2(t *testing.T) {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	res := allPathsSourceTarget(graph)
	log.Printf("res: %v", res)
}
