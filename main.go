package main

import (
	"fmt"
	"os"
)

type Len string

func (l Len) String() string {
	return fmt.Sprintf("%3d %s", len(l), string(l))
}

func main() {
	for _, a := range os.Args[1:] {
		fmt.Println(Len(a))
	}
}
