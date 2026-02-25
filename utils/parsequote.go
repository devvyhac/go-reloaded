package utils

// func ParseQuote() {

// }

func FindPairIndex(passage string) []int {
	var foundOne bool
	var pair []int
	// var indexes [][]int

	quote := "'"

	for i, item := range passage {
		if string(item) == quote {
			if !foundOne {
				pair = append(pair, i)
				foundOne = true
			} else {
				pair = append(pair, i)
				foundOne = false
			}
		}
	}
	return pair
}
