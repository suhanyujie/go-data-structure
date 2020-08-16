package main

import "fmt"

/// LeetCode 地址：https://leetcode-cn.com/problems/course-schedule-ii/solution/ke-cheng-biao-ii-by-leetcode-solution/
/// 有向图

func main() {
	s1 := [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	}
	res := findOrderByDfs(4, s1)
	fmt.Println(res)
	res = findOrderByBreadth(4, s1)
	fmt.Println(res)
}

/// 深度优先算法:
/// 通过一些简单的例子，我们可以画出依赖关系图
/// n 个课程，也就意味着 0,1...n-1 恰好分别是对应课程的代号，因而可以"课程代号"作为数组的索引
/// 因为一门课程，可能会依赖多个课程（多条边），因此，以课程代号为 key，对应的数组值为其依赖项。形如 [1=>[2,4],]
/// 因此，我们先通过二维数组，装载好依赖关系
/// 遍历每一项，解决依赖，并递归调用依赖的每个课程
/// 如果出现了 a 依赖 b，b 又依赖 a 这种循环依赖，只需在访问每个课程时，设定好状态 0：未访问；1：访问中；2：访问完成
/// 当访问到无依赖的课程时，将其设为 2，并将此课程加入到"结果栈"中
func findOrderByDfs(numCourses int, prerequisites [][]int) []int {
	var (
		// 边
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
		dfs     func(u int)
	)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		// 处理完后，赋值2
		visited[u] = 2
		result = append(result, u)
	}
	// 初始化，装载关系数据
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

func ShowGraph() {
	fmt.Println("this is digraph")
}

/// 广度优先
func findOrderByBreadth(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		// 统计入度：入度表示某个课程的依赖。
		// 入度为1，表示依赖一个课程，入度为2，表示依赖2个课程
		indeg[info[0]]++
	}
	q := []int{}
	for i := 0; i < numCourses; i++ {
		// 寻找入度为 0 的课程
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		// 此时的 u，入度为0
		// 查询，是否有课程依赖 u
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}
	return result
}
