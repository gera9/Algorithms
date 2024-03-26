package main

import (
	"fmt"
	"io"
)

// Tower of Hanoi
func TowerOfHanoi(w io.Writer, n int, from, aux, to string) {
	if n == 0 {
		return
	}

	TowerOfHanoi(w, n-1, from, to, aux)

	fmt.Fprintf(w, "Disk %d moved from %s to %s.\n", n, from, to)

	TowerOfHanoi(w, n-1, aux, from, to)
}
