package helpers

import (
	"errors"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
)

func ParseArgs(args []string) ([]int, error) {
	if len(args) == 0 {
		return []int{}, nil
	}

	var fields []string
	if len(args) == 1 {
		fields = strings.Fields(args[0])
	} else {
		fields = make([]string, 0, len(args))
		for _, a := range args {
			for _, f := range strings.Fields(a) {
				fields = append(fields, f)
			}
		}
	}

	joined := strings.TrimSpace(strings.Join(fields, " "))
	if strings.HasPrefix(joined, "<") && strings.HasSuffix(joined, ">") {
		inner := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(joined, "<"), ">"))
		parts := strings.Fields(inner)
		if len(parts) >= 3 && strings.ToLower(parts[1]) == "random" && strings.ToLower(parts[2]) == "numbers" {
			n, err := strconv.Atoi(parts[0])
			if err == nil && n > 0 {
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
