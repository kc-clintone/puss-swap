package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		switch line {
		case "pa":
			helpers.Pa(a, b)
		case "pb":
			helpers.Pb(a, b)
		case "sa":
			helpers.Sa(a)
		case "sb":
			helpers.Sb(b)
		case "ss":
			helpers.Ss(a, b)
		case "ra":
			helpers.Ra(a)
		case "rb":
			helpers.Rb(b)
		case "rr":
			helpers.Rr(a, b)
		case "rra":
			helpers.Rra(a)
		case "rrb":
			helpers.Rrb(b)
		case "rrr":
			helpers.Rrr(a, b)
		default:
			// invalid instruction -> Error
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if helpers.IsSorted(a) && b.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
