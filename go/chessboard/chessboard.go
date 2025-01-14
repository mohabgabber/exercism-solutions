package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var num int
	for _, c := range cb[file] {
		if c {
			num++
		}
	}
	return num
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var num int
	if rank <= 8 && rank >= 1 {
		for _, c := range cb {
			if c[rank-1] {
				num++
			}
		}
	}
	return num
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var num int
	for _, c := range cb {
		for range c {
			num++
		}
	}
	return num
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var num int
	for _, c := range cb {
		for _, t := range c {
			if t {
				num++
			}
		}
	}
	return num
}
