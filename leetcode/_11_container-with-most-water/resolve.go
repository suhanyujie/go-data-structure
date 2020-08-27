package main

import "fmt"

/// 题目地址：https://leetcode-cn.com/problems/container-with-most-water/
func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(arr)
	fmt.Println(res)
}

// 找到盛最多水的容器 range
func maxArea(height []int) int {
	var (
		l       int = 0
		r       int = len(height) - 1
		maxArea int = 0
	)
	for {
		if l > r {
			break
		}
		min1 := min(height[l], height[r])
		area := min1 * (r - l)
		if maxArea < area {
			maxArea = area
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return maxArea
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}
