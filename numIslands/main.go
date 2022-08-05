package main

func main() {

	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}

	println(numIslands(grid))

}

func numIslands(grid [][]byte) int {
	counter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				counter++
				dfs(&grid, i, j)
			}
		}
	}

	return counter
}

func dfs(grid *[][]byte, i, j int) {
	if i > len(*grid)-1 || j > len((*grid)[0])-1 || i < 0 || j < 0 || (*grid)[i][j] == '0' {
		return
	}

	(*grid)[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}