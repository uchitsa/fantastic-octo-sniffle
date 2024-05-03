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

	var s string
	fmt.Fscan(in, &s)

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var start, end int
		var ri string
		fmt.Fscan(in, &start, &end, &ri)

		s = s[:start-1] + ri + s[end:]
	}

	fmt.Fprint(out, s)
}
