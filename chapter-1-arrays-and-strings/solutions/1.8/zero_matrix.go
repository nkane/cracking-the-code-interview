package zero_matrix

/*
	1.8: Write an algorithm such that if an element in a M x N matrix is 0, its entire row
	and column are set to 0.
*/

/*
At first glance, this problem seems easy: just iterate through the matrix and every time we see a cell
with value zero, set its row and column to 0. There's one problem with that solution though: when we
come across other cells in that row or column, we'll see the zeros and change their row and columns
to zero. Pretty soon, our entire matrix will be set to zeros.

One way around this is to keep a second matrix which flags the zero locations. We would then do a second
pass through the matrix to set the zeros. This would take O(MN) space.

Do we really need O(MN) space? No. Since we're going to set the entire row and column to zero, we don't
need to track that it was exactly cell[2][4] (row 2, column 4). We only need to know that row 2 has
a zero somewhere, and column 4 has a zero somewhere. We'll set the entire row and column to zero anyway,
so why would we care to keep track of the exact location of the zero?

The code below implements this algorithm. We use two arays to keep track of all the rows with zeros
and all the columns with zeros. We then nullify rows and columns based on the values in these arrays.
*/

type ZeroMatrix func(matrix [][]int)

func ZeroMatrixInefficient(matrix [][]int) {
	row := make([]bool, len(matrix))
	column := make([]bool, len(matrix[0]))
	// store the row and column index with value 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}
	// nullify rows
	for i := 0; i < len(row); i++ {
		if row[i] {
			NullifyRow(matrix, i)
		}
	}
	// nullify column
	for j := 0; j < len(column); j++ {
		if column[j] {
			NullifyColumn(matrix, j)
		}
	}
}

func NullifyRow(matrix [][]int, row int) {
	for j := 0; j < len(matrix[0]); j++ {
		matrix[row][j] = 0
	}
}

func NullifyColumn(matrix [][]int, column int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
}

/*
To make this somewhat more space efficient, we could use a bit vector instead of a boolean arrya. It
would still be O(N) space.

We can reduce the space to O(1) by using the first column as a replacement for the row array and the first
row as a replacement for the column array. This works as follows:
1. Check if the first row and first column have any zeros, and set variables rowHasZero and columnHasZero.
(We'll nullify the first row and first column later, if necessary).
2. Iterate through the rest of the matrix, setting matrix[i][0] and matrix[0][j] to zero whenever
there's a zero in the matrix[i][j].
3. Iterate through rest of matrix, nullifying row i if there's z zero in matrix[i][0].
4. Iterate through rest of matrix, nullifying column j if there's a zero in matrix[0][j].
5. Nullify the first row and first column, if necessary (based on values from Step 1).

This code has a lot of "do this for rows, then the equivalent action for the column". In an interview,
you could abbreviate this code by adding comments and TODOs that explain that the next chunk of code
looks the same as the earlier code, but using rows. This would allow you to focus on the most important
parts of the algorithm.
*/

func ZeroMatrixEfficient(matrix [][]int) {
	rowHasZero := false
	columnHasZero := false

	// check if first row has a zero
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			rowHasZero = true
			break
		}
	}

	// check if first column has a zero
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			columnHasZero = true
			break
		}
	}

	// check for zeros in the rest of the array
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// nullify rows based on values in the first column
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			NullifyRow(matrix, i)
		}
	}

	// nullify columns based on values in first row
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			NullifyColumn(matrix, j)
		}
	}

	// nullify first row
	if rowHasZero {
		NullifyRow(matrix, 0)
	}

	// nullify first column
	if columnHasZero {
		NullifyColumn(matrix, 0)
	}
}
