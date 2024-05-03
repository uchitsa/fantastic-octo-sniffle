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
		var n int
		fmt.Fscan(in, &n)

		var queue string
		fmt.Fscan(in, &queue)

		// Проверяем, можно ли выбрать правильные пары
		if canFormPairs(queue) {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}
}

func canFormPairs(queue string) bool {
	// Счетчики для каждого типа сообщений
	var xCount, yCount, zCount int

	// Проходим по очереди и считаем количество каждого типа сообщений
	for _, msg := range queue {
		switch msg {
		case 'X':
			xCount++
		case 'Y':
			yCount++
		case 'Z':
			zCount++
		}
	}

	// Проверяем, можно ли составить правильные пары
	if xCount%2 == 0 && yCount%2 == 0 && zCount%2 == 0 {
		return true
	}
	if xCount%2 == 1 && yCount%2 == 1 && zCount%2 == 1 {
		return true
	}
	if xCount%2 == 1 && yCount%2 == 0 && zCount%2 == 0 {
		return true
	}
	if xCount%2 == 0 && yCount%2 == 1 && zCount%2 == 0 {
		return true
	}
	if xCount%2 == 0 && yCount%2 == 0 && zCount%2 == 1 {
		return true
	}

	return false
}
