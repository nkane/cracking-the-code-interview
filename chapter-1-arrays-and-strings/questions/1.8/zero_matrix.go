package zero_matrix

/*
	1.8: Write an algorithm such that if an element in a M x N matrix is 0, its entire row
	and column are set to 0.
*/

type Cell struct {
	Row    int
	Column int
}

func ZeroMatrix(matrix [][]int) {
	rows := len(matrix)
	columns := len(matrix[0])
	cells := []Cell{}
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if matrix[r][c] == 0 {
				cells = append(cells, Cell{
					Row:    r,
					Column: c,
				})
			}
		}
	}
	// set zeros
	for _, cell := range cells {
		// set entire row to zero
		for c := 0; c < columns; c++ {
			matrix[cell.Row][c] = 0
		}
		// set entire column to zero
		for r := 0; r < rows; r++ {
			matrix[r][cell.Column] = 0
		}
	}
}
