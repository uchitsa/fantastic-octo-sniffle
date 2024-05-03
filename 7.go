package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)

		windows := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &windows[j])
		}

		result := solve(n, windows)
		fmt.Fprintln(out, result)
	}
}

func solve(n int, windows []int) string {
	occupied := make([]bool, n+1)
	result := make([]byte, len(windows))

	for i, window := range windows {
		if occupied[window] {
			return "x"
		}
		occupied[window] = true
		result[i] = '0'
	}

	for i, window := range windows {
		if window > 1 && !occupied[window-1] {
			result[i] = '-'
			occupied[window-1] = true
			occupied[window] = false
		} else if window < n && !occupied[window+1] {
			result[i] = '+'
			occupied[window+1] = true
			occupied[window] = false
		}
	}

	return string(result)
}
