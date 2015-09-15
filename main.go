package main

import (
	"fmt"
	"os"
)

func main() {
	for _, a := range os.Args[1:] {
		fmt.Printf("%3d %v\n", len(a), a)
	}
}
