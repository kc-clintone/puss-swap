package helpers

import (
	"errors"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
)

// ParseArgs takes a string of space-separated integers, validates them, and returns a slice of integers. It checks for empty input, non-integer values, and duplicates, returning appropriate errors for each case. If the input is valid, it returns the slice of integers.
//
// Special feature: if the single token is of the form "<N random numbers>" (e.g. "<10 random numbers>"),
// ParseArgs will generate N unique random integers and parse those instead.
func ParseArgs(args []string) ([]int, error) {
	// no arguments -> empty slice (caller may decide to exit/print nothing)
	if len(args) == 0 {
		return []int{}, nil
	}

	var fields []string
	if len(args) == 1 {
		// single quoted string: split into fields
		fields = strings.Fields(args[0])
	} else {
		// multiple arguments provided separately
		fields = make([]string, 0, len(args))
		for _, a := range args {
			// allow tokens that may contain extra whitespace/newlines
			for _, f := range strings.Fields(a) {
				fields = append(fields, f)
			}
		}
	}

	// special placeholder handling: "<N random numbers>"
	// support placeholder whether the quoted argument was split into multiple fields
	joined := strings.TrimSpace(strings.Join(fields, " "))
	if strings.HasPrefix(joined, "<") && strings.HasSuffix(joined, ">") {
		inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(joined, "<"), ">"))
		parts := strings.Fields(inner)
		if len(parts) >= 3 && strings.ToLower(parts[1]) == "random" && strings.ToLower(parts[2]) == "numbers" {
			n, err := strconv.Atoi(parts[0])
			if err == nil && n > 0 {
				// generate n unique random numbers
				// deterministic seed derived from the placeholder inner text
				h := fnv.New32a()
				h.Write([]byte(inner))
				seed := int64(h.Sum32())
				r := rand.New(rand.NewSource(seed))
				maxRange := n * 100
				if maxRange < 1000 {
					maxRange = 1000
				}
				perm := r.Perm(maxRange)
				gen := make([]string, 0, n)
				for i := 0; i < n; i++ {
					gen = append(gen, strconv.Itoa(perm[i]))
				}
				fields = gen
			}
		}
	}

	if len(fields) == 0 {
		return []int{}, nil
	}

	nums := make([]int, 0, len(fields))
	seen := make(map[int]bool)

	for _, f := range fields {
		t := strings.TrimSpace(f)
		if t == "" {
			continue
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			return nil, errors.New("invalid integer")
		}
		if seen[n] {
			return nil, errors.New("duplicate")
		}
		seen[n] = true
		nums = append(nums, n)
	}

	return nums, nil
}
