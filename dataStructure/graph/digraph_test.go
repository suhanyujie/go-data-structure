package main

import "testing"

func TestShowGraph(t *testing.T) {
	ShowGraph()
	t.Fatal("this is test log")
}

func TestRun(t *testing.T) {
	t.Fatal("this is test log for Run")
}

// /usr/local/go/bin/go test -c -o /private/var/folders/fw/gtsqf2z15wg1xg5bfm26l60w0000gn/T/___TestRun_in_digraph_test_go -gcflags "all=-N -l" /Volumes/su2/userSa/www2/tech/repo/golang/go-data-structure/dataStructure/graph/digraph_test.go /Volumes/su2/userSa/www2/tech/repo/golang/go-data-structure/dataStructure/graph/digraph.go
