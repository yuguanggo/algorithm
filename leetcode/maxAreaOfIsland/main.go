package main

import (
	"fmt"
	"math"
)

func dfs(grid [][]int, x int, y int) int {
	var (
		sum int = 1
		n   int
		m   int
		dir [][]int
	)
	dir = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	n = len(grid)
	m = len(grid[0])
	grid[x][y] = 0 //当前元素设置成0避免再次搜索到
	for i := 0; i < 4; i++ {
		xn := x + dir[i][0]
		yn := y + dir[i][1]
		if xn >= 0 && xn < n && yn >= 0 && yn < m && grid[xn][yn] == 1 {
			sum += dfs(grid, xn, yn)
		}
	}
	return sum
}

func maxAreaOfIsland(grid [][]int) int {
	var (
		mx int
		n  int
		m  int
	)
	mx = 0
	n = len(grid)
	m = len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				mx = int(math.Max(float64(dfs(grid, i, j)), float64(mx)))
			}
		}
	}
	return mx
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}
