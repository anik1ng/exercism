package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cells, exists := cb[file]
    if !exists {
        return 0
    }

    result := 0
    for _, c := range cells {
        if c {
            result++
        }
    }
    return result
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }

    result := 0
    for _, column := range cb {
        if column[rank-1] {
            result++
        }
    }
    return result
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	result := 0
    for _, column := range cb {
        for range column {
            result++
        }
    }
    return result
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	result := 0
    for _, column := range cb {
        for _, s := range column {
            if s {
                result++
            }
        }
    }
    return result
}
