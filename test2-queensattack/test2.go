package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(s []int, value int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func main() {

	in := bufio.NewReader(os.Stdin)

	var n, k, r_q, c_q int

	fmt.Fscan(in, &n, &k)
	fmt.Fscan(in, &r_q, &c_q)

	// Declare array list
	var ll, ll2 [][]int
	// Loop array list
	for i := 0; i <= n; i++ {
		ll = append(ll)
		ll2 = append(ll2)
	}

	for a0 := 0; a0 < k; a0++ {
		var r, c int

		fmt.Fscan(in, &r, &c)

		ll[r] = append(ll2[c])
		ll2[c] = append(ll[r])
	}

	var result int64 = 0
	for i := c_q - 1; i >= 1; i-- {
		if contains(ll[r_q], i) {
			break
		}
		result++
	}

	for i := c_q + 1; i <= n; i++ {
		if contains(ll[r_q], i) {
			break
		}
		result++
	}

	for i := r_q - 1; i >= 1; i-- {
		if contains(ll2[c_q], i) {
			break
		}
		result++
	}

	for i := r_q + 1; i <= n; i++ {
		if contains(ll2[c_q], i) {
			break
		}
		result++
	}

	var cc int = c_q - 1
	for i := r_q - 1; i >= 1; i-- {
		if cc == 0 || contains(ll[i], cc) {
			break
		}
		cc--
		result++
	}

	cc = c_q - 1
	for i := r_q + 1; i <= n; i++ {
		if cc == 0 || contains(ll[i], cc) {
			break
		}
		cc--
		result++
	}

	cc = c_q + 1
	for i := r_q + 1; i <= n; i++ {
		if cc == n+1 || contains(ll[i], cc) {
			break
		}
		cc++
		result++
	}

	cc = c_q + 1
	for i := r_q - 1; i >= 1; i-- {
		if cc == n+1 || contains(ll[i], cc) {
			break
		}
		cc++
		result++
	}

	fmt.Println(result)

}
