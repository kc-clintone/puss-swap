package main

import (
	"fmt"
	"os"

	"pushswap/helpers"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	nums, err := helpers.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	a := helpers.NewStack(nums)
	b := helpers.NewStack([]int{})

	ops := helpers.Sort(a, b)

	for _, op := range ops {
		fmt.Println(op)
	}
}
