package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type client struct {
	wnd int
	id  int
	act string
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for ti := 0; ti < t; ti++ {
		var n, m int
		fmt.Fscanln(in, &n, &m)

		cleints := make([]client, m)
		for mi := 0; mi < m; mi++ {
			var w int
			fmt.Fscan(in, &w)
			cleints[mi] = client{
				wnd: w,
				id:  mi,
			}
		}
		fmt.Fscanln(in)

		res, err := updateClients(cleints, n)
		if err != nil {
			fmt.Fprintln(out, "x")
			continue
		}

		fmt.Fprintln(out, res)
	}
}

func updateClients(clients []client, n int) (string, error) {
	if len(clients) > n {
		return "", fmt.Errorf("too many cleints")
	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].wnd < clients[j].wnd
	})

	if clients[0].wnd > 1 {
		clients[0].wnd--
		clients[0].act = "-"
	} else {
		clients[0].act = "0"
	}

	for i := 1; i < len(clients); i++ {
		if (clients[i].wnd - clients[i-1].wnd) == 1 {
			clients[i].act = "0"
		}

		if (clients[i].wnd - clients[i-1].wnd) > 1 {
			clients[i].wnd--
			clients[i].act = "-"
		}

		if (clients[i].wnd - clients[i-1].wnd) == 0 {
			clients[i].wnd++
			clients[i].act = "+"
		}

		if (clients[i].wnd - clients[i-1].wnd) < 0 {
			return "", fmt.Errorf("update is not possible")
		}

		if clients[i].wnd > n {
			return "", fmt.Errorf("update is not possible")
		}

	}

	sort.Slice(clients, func(i, j int) bool {
		return clients[i].id < clients[j].id
	})

	sb := strings.Builder{}
	for _, p := range clients {
		sb.WriteString(p.act)
	}

	return sb.String(), nil
}
