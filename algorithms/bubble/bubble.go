package algorithms

// var xi = []int{53, 12, 1024, 35} // ,  101, 11

// func presort() {
// 	fmt.Printf("Presort\t:%v\n", xi)
// }

// func postsort() {
// 	fmt.Printf("Sorted\t:%v\n", xi)
// }

/*
 *
 */
// func main() {
// 	presort()
// 	bubble(xi)
// 	postsort()
// }

//Bubble ...
func Bubble(xi []int) []int {
	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi)-i-1; j++ {
			// fmt.Printf("outer loop index [%v] inner loop comparison: %v > %v?\n", i, xi[j], xi[j+1])
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}

	return xi
}
