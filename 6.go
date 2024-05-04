package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	line, _ := in.ReadString('\n')
	parts := strings.Split(strings.TrimSpace(line), " ")

	friendsAmount, err := strconv.Atoi(parts[0])
	if err != nil {
		friendsAmount = -1
	}
	cardsAmount, err := strconv.Atoi(parts[1])
	if err != nil {
		cardsAmount = -1
	}

	line, _ = in.ReadString('\n')
	inp := strings.Split(strings.TrimSpace(line), " ")

	targetRedirect := make(map[int]int)
	var dirtyNodes []int
	var answer []string

	for i := 0; i < friendsAmount; i++ {
		cardIndex, _ := strconv.Atoi(inp[i])
		tempIndex := cardIndex + 1
		for {
			if nextIndex, ok := targetRedirect[tempIndex]; ok {
				tempIndex = nextIndex
				dirtyNodes = append(dirtyNodes, tempIndex)
			} else {
				dirtyNodes = append(dirtyNodes, tempIndex)
				break
			}
			if _, ok := targetRedirect[tempIndex]; !ok {
				break
			}
		}

		if tempIndex > cardsAmount {
			fmt.Fprintln(out, "-1")
			return
		}

		for _, node := range dirtyNodes {
			if _, ok := targetRedirect[node]; ok {
				targetRedirect[node] = tempIndex
			} else {
				targetRedirect[tempIndex] = tempIndex + 1
			}
		}

		answer = append(answer, strconv.Itoa(tempIndex))
		dirtyNodes = nil
	}

	fmt.Fprintln(out, strings.Join(answer, " "))
}