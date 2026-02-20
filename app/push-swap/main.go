package main

import (
	"fmt"
	"os"

	"pushswap/helpers"
)

// The main function initializes stack a with the integers parsed from the command-line arguments and an empty stack b. It then calls the Sort function from the helpers package to get the list of operations needed to sort stack a using stack b as auxiliary. Finally, it prints each operation to standard output, one per line. If there is an error during parsing, it prints "Error" to standard error and exits with a non-zero status code.

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
