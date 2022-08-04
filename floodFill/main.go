package main

import "fmt"

func main() {
	image := make([][]int,3,3)
	image[0] = []int{1,1,1}
	image[1] = []int{1, 1, 1}
	image[2] = []int{1, 0, 1}
	fmt.Print(floodFill(image,1,1,2))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color { return image }
	fill(&image, sr, sc, image[sr][sc], color)
	return image
}

func fill(image *[][]int, i int, j int, oldColor int, newColor int) {
	if i < 0 || i >= len(*image) || j < 0 || j >= len((*image)[0]) {return}
	if (*image)[i][j] != oldColor { return }
	(*image)[i][j] = newColor
	fill(image, i+1, j, oldColor, newColor)
	fill(image, i, j+1, oldColor, newColor)
	fill(image, i-1, j, oldColor, newColor)
	fill(image, i, j-1, oldColor, newColor)
}