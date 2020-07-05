package main

import "fmt"

func ExampleLen() {
	for _, v := range []string{"whatever", "random", "something"} {
		fmt.Println(Len(v))
	}
	// Output:
	//
	//   8 whatever
	//   6 random
	//   9 something
}
