package main

import "fmt"

func findMoves_i(a []int) int {

	if len(a) == 1 {
		return 0
	}

	t := 0

	for len(a) > 1 {
		l := 0
		lSum := 0
		r := len(a) - 1
		rSum := 0

		for l <= r {
			if lSum <= rSum {
				lSum += a[l]
				l++
			} else if rSum < lSum {
				rSum += a[r]
				r--
			}
		}

		//fmt.Printf("lSum %d [%d] rSum %d [%d]\n", lSum, l, rSum, r)
		if lSum != rSum {
			break
		} else {
			t += 1
		}

		if lSum == 0 && rSum == 0 {
			// an array of zeros take tip and tail
			//fmt.Println("zero array")
			a = a[0 : len(a)-1]
		} else {
			// calc new array
			lLen := l
			rLen := (len(a) - 1) - r

			if lLen > rLen {
				a = a[:l]
			} else if lLen < rLen {
				a = a[r+1:]
			} else {
				// They are the same take right or left.
				// could consolidate this into an or with the above
				//a = a[r + 1:]
				a = a[:l]
			}
		}

		/*
		   fmt.Println("-- dump array for next call --")
		   for x := range a {
		       fmt.Printf("%d ", a[x])
		   }
		   fmt.Println()
		   fmt.Println()
		*/
	}

	return t
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var n int
	var x int

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		fmt.Scanf("%d\n", &x)
		a := make([]int, x)

		for y := 0; y < x; y++ {
			if y == (x - 1) {
				fmt.Scanf("%d\n", &a[y])
			} else {
				fmt.Scanf("%d", &a[y])
			}
		}

		fmt.Printf("%d\n", findMoves_i(a))
	}
}
