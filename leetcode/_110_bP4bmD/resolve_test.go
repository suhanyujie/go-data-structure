package _110_bP4bmD

import (
	"log"
	"reflect"
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
	expectRes := [][]int{
		{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4},
	}
	if !reflect.DeepEqual(res, expectRes) {
		t.Errorf("res error")
	}
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
	expectRes := [][]int{
		{0, 1, 3},
		{0, 2, 3},
	}
	if !reflect.DeepEqual(res, expectRes) {
		t.Errorf("res error")
	}
}
