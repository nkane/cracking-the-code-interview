package rotate_matrix

/*
	1.7: Rotate Matrix: Given an image represented by an N x N matrix, where each pixel in the image is represented
	by an integer, write a method to rotate the image by 90 degrees. Can you do this in place?

	required help on this one:
	- https://www.youtube.com/watch?v=fMSJSS7eO1w
*/

func RotateMatrix(matrix [][]int) {
	left := 0
	right := len(matrix) - 1
	for left < right {
		for i := 0; i < right-left; i++ {
			top := left
			bottom := right
			// save top left value
			topLeft := matrix[top][left+i]
			// move bottom left into top left
			matrix[top][left+i] = matrix[bottom-i][left]
			// move bottom right into bottom left
			matrix[bottom-i][left] = matrix[bottom][right-i]
			// move top rigth into bottom right
			matrix[bottom][right-i] = matrix[top+i][right]
			// move top left into top right
			matrix[top+i][right] = topLeft
		}
		right--
		left++
	}
}
