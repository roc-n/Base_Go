package main

import (
	"algorithm.labmem.net/al"
)

func main() {
	// ds.Array()

	// b := []byte("123 456 789 aa bb cc 200100")
	// var x int
	// for i := 0; i < len(b); {
	// 	x, i = nextInt(b, i)
	// 	fmt.Println(x)
	// }

	// ds.ListDemo()

	// 算法集合
	// al.BS()
	// al.PermutationsIDemo()
	// result := common.Fibonacci(5)
	// fmt.Println(result)
	al.SortDemo()
	// al.GreedyDemo()
}

// func nextInt(b []byte, i int) (value, nextPos int) {
// 	for ; i < len(b) && !isDigit(b[i]); i++ {
// 	}
// 	x := 0
// 	for ; i < len(b) && isDigit(b[i]); i++ {
// 		x = x*10 + int(b[i]) - '0'
// 	}
// 	return x, i
// }

// func isDigit(b byte) bool {
// 	return b >= '0' && b <= '9'
// }
