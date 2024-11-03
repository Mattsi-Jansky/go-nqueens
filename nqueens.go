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

		if invalidState(board) {
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

func invalidState(board []int) bool {
	existing_row := make(map[int]int)
	for _, row := range board {
		_, ok := existing_row[row]
		if ok {
			return true
		} else {
			existing_row[row] = row
		}

	}
	return false
}
