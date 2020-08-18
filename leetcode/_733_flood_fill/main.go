package main

import "fmt"

/// 733.图像渲染
/// 力扣地址：https://leetcode-cn.com/problems/flood-fill
/// tips：看了好几遍这个题，感觉有些读不懂，还是直接看看官方题解吧

func main() {
	var res [][]int
	res = floodFill([][]int{
		[]int{1, 1, 1},
		[]int{1, 1, 0},
		[]int{1, 0, 1},
	}, 1, 1, 2)
	fmt.Println(res)
	res = floodFill([][]int{
		[]int{0, 0, 0},
		[]int{0, 0, 0},
	}, 0, 0, 2)
	fmt.Println(res)
	res = floodFill([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
	}, 1, 1, 2)
	fmt.Println(res)
	res = floodFill([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 1},
	}, 1, 1, 1)
	fmt.Println(res)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	var (
		q              [][]int
		maxY           = len(image)
		maxX           = len(image[0])
		adjacentPoints [][]int
	)
	// 利用广度优先算法，从矩阵中找出相邻的点坐标，每个点最多有4个相邻点
	// 以给出的点(row, col)作为初始点，可以分别向外扩散：(row-1, col),(row, col-1),(row+1, y),(row, col+1)
	// 需要注意的是：点的横、纵坐标不能超出指定的矩阵宽、高
	// 如果起始点的色值和新的色值一样，则无需重新渲染，直接返回 image
	initColor := image[sr][sc]
	q = append(q, []int{sr, sc})
	for len(q) > 0 {
		tmpY := q[0][0]
		tmpX := q[0][1]
		// 如果跟给定的起始点的颜色相同，则需更新为新的颜色
		// 一旦确定某个点是「相邻」且「同色」的点，则继续寻找该点的周边色块
		if image[tmpY][tmpX] == initColor {
			image[tmpY][tmpX] = newColor
			adjacentPoints = get4Point(tmpY, tmpX)
			for _, tmpPoint := range adjacentPoints {
				if tmpPoint[0] >= 0 && tmpPoint[0] < maxY && tmpPoint[1] >= 0 && tmpPoint[1] < maxX && image[tmpPoint[0]][tmpPoint[1]] == initColor {
					q = append(q, tmpPoint)
				}
			}
		}
		q = q[1:]
	}

	return image
}

func get4Point(sr int, sc int) [][]int {
	var adjacentPoints [][]int
	adjacentPoints = append(adjacentPoints, []int{sr - 1, sc})
	adjacentPoints = append(adjacentPoints, []int{sr, sc - 1})
	adjacentPoints = append(adjacentPoints, []int{sr + 1, sc})
	adjacentPoints = append(adjacentPoints, []int{sr, sc + 1})

	return adjacentPoints
}
