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

		//  If the board is not valid, continue
		if len(board) == size {
			results = append(results, board)
		}
		//  If the board is incomplete, add next possible states to frontier
	}

	return results
}
