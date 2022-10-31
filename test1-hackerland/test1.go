package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var d int

	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n, &d)
	expenditure := make([]int, n)
	createArr(expenditure, 0, n)

	var result int = activityNotifications(expenditure, d)
	fmt.Println(result)
}

func activityNotifications(expenditure []int, d int) int {

	var count [200]int

	var result int = 0

	for i := 0; i < d; i++ {
		count[expenditure[i]]++
	}

	for i := d; i < len(expenditure); i++ {

		var median int = getMedian(count[:], d)

		if median <= expenditure[i] {
			result++
		}

		count[expenditure[i-d]]--
		count[expenditure[i]]++

	}

	return result

}

func getMedian(count []int, d int) int {

	var sum int = 0

	for i := 0; i < len(count); i++ {

		sum += count[i]

		if (2 * sum) == d {
			return (2*i + 1)
		} else if (2 * sum) > d {
			return (i * 2)
		}

	}

	return 1

}

func createArr(all []int, i, n int) {

	if n == 0 {
		return
	}

	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}

	createArr(all, i+1, n-1)

}

func Scan(a *int) (int, error) {
	return fmt.Scan(a)
}
