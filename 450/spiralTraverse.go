package main

func SpiralTraverse(array [][]int) []int {
	// Write your code here.

	if len(array) == 0 {
		return []int{}
	}
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	final := []int{}

	for startRow <= endRow && startCol <= endCol {
		for i := startCol; i <= endCol; i++ {
			final = append(final, array[startRow][i])
		}
		for j := startRow + 1; j <= endRow; j++ {
			final = append(final, array[j][endCol])
		}
		for k := endCol - 1; k >= startCol; k-- {
			if startRow == endRow {
				break
			}

			final = append(final, array[endRow][k])
		}
		for l := endRow - 1; l > startRow; l-- {
			if startCol == endCol {
				break
			}
			final = append(final, array[l][startCol])
		}
		startRow++
		endRow--
		startCol++
		endCol--
	}
	return final
}
