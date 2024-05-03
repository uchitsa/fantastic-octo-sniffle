package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RequestType int

const (
	Send RequestType = iota + 1
	ShowLastest
)

var reader *bufio.Reader
var writer *bufio.Writer

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	input, _ := reader.ReadString('\n')
	values := strings.Fields(input)

	n, _ := strconv.Atoi(values[0])
	q, _ := strconv.Atoi(values[1])

	notifs := make([]int, n+1)
	num := 1

	for i := 0; i < q; i++ {
		input, _ = reader.ReadString('\n')
		values = strings.Fields(input)

		typeVal, _ := strconv.Atoi(values[0])
		typeReq := RequestType(typeVal)
		id, _ := strconv.Atoi(values[1])

		if typeReq == Send {
			notifs[id] = num
			num++
		} else if typeReq == ShowLastest {
			max := notifs[0]
			if notifs[id] > max {
				max = notifs[id]
			}
			fmt.Fprintln(writer, max)
		}
	}
}
