package nqueens

func Nqueens(size int) [][]int {
	frontier := make([][]int, 0)
	results := make([][]int, 0)
	for i := 0; i < size; i++ {
		frontier = append(frontier, []int{i})
	}

	for len(frontier) > 0 {
		board := frontier[0]
		frontier = frontier[1:]

		if InvalidState(board) {
			continue
		}
		if len(board) == size {
			results = append(results, board)
		} else {
			for j := 0; j < size; j++ {
				newBoard := make([]int, len(board))
				copy(newBoard, board)
				newBoard = append(newBoard, j)
				frontier = append(frontier, newBoard)
			}
		}
	}

	return results
}

func InvalidState(board []int) bool {
	existing_row := make(map[int]int)
	for _, row := range board {
		_, ok := existing_row[row]
		if ok {
			return true
		} else {
			existing_row[row] = row
		}

	}

	last_column_added := board[len(board)-1]
	current := []int{(len(board) - 1), last_column_added}
	for columToCheck := 0; columToCheck < len(board)-1; columToCheck++ {
		toCompare := []int{columToCheck, board[columToCheck]}

		diffColumn := current[0] - toCompare[0]
		diffRow := current[1] - toCompare[1]

		if diffColumn < 0 {
			diffColumn = -diffColumn
		}

		if diffRow < 0 {
			diffRow = -diffRow
		}

		if diffColumn == diffRow {
			return true
		}
	}
	return false
}
