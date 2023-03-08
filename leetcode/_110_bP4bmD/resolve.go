package _110_bP4bmD

import "log"

// 剑指 Offer II 110. 所有路径
// ref: https://leetcode.cn/problems/bP4bmD/description/

func main() {
	graph := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	graph = [][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}
	res := allPathsSourceTarget(graph)
	log.Printf("res: %v", res)
}

var (
	path   = make([]int, 0)
	result = make([][]int, 0)
)

// 寻路
// ref: https://leetcode.cn/problems/bP4bmD/solutions/939367/tong-guan-jian-2-bfs-hui-su-by-muluo-2-w4gy/
func allPathsSourceTarget(graph [][]int) [][]int {
	// clear
	path = make([]int, 0)
	result = make([][]int, 0)
	// 回溯法+递归 todo
	path = append(path, 0)
	var dfs func(graph [][]int, node int)
	dfs = func(graph [][]int, node int) {
		if node == len(graph)-1 {
			oneRes := make([]int, len(path))
			copy(oneRes, path)
			result = append(result, oneRes)
			return
		}
		for _, v := range graph[node] {
			path = append(path, v)
			dfs(graph, v)
			path = path[:len(path)-1]
		}
	}
	dfs(graph, 0)
	return result
}
