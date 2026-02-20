# push-swap (Go)

A small Go implementation of the **push-swap** and **checker** programs.

---

## Overview

### push-swap

`push-swap` takes a stack of integers as input and prints a sequence of push-swap instructions that sorts the stack in ascending order.

* The **first integer** is the **top** of the stack.
* The result must place the **smallest value at the top**.
* The program prints only the list of instructions.

### checker

`checker` takes:

* A stack of integers as arguments.
* A list of instructions from standard input.

It executes the instructions and prints:

* `OK` if:

  * Stack `a` is sorted in ascending order, and
  * Stack `b` is empty.
* `KO` otherwise.

### Behavior Rules

* If called with **no arguments**, both programs print nothing.
* On error, they print `Error` to **stderr** and exit with a non-zero status code.

---

## Project Structure

```
app/push-swap   Command-line entry point for push-swap
app/checker     Command-line entry point for checker
helpers         Stack implementation, parsing, operations, algorithms
README.md       Project documentation
```

---

## Build

From the repository root:

```
go build -o push-swap ./app/push-swap
go build -o checker ./app/checker
```

---

## Usage

You can provide input in two ways:

### Single Argument (quoted string)

```
./push-swap "2 1 3 6 5 8"
```

### Multiple Arguments

```
./push-swap 2 1 3 6 5 8
```

The first number is always the top of the stack.

### Validate Output

```
./push-swap "4 67 3 87 23" | ./checker "4 67 3 87 23"
```

### Empty Input

```
./push-swap
```

Produces no output.

---

## Placeholder Input

The parser supports a deterministic placeholder format:

```
ARG="<10 random numbers>"
./push-swap "$ARG"
```

The placeholder generates the same list of numbers for the same placeholder text.
If both programs receive the same placeholder string, they will generate the same numbers.

### Shell-Based Random Input

For custom random input:

```
ARG="$(shuf -i 1-1000 -n 10 | paste -sd ' ' -)"
./push-swap "$ARG" | ./checker "$ARG"
```

For reproducible testing, generate the list once and reuse the same variable.

---

## Error Handling

The programs print `Error` to **stderr** and exit with a non-zero status if:

* A token is not a valid integer.
* There are duplicate values.
* `checker` receives an invalid instruction from standard input.

If `checker` receives valid arguments but empty standard input, it evaluates the stacks and prints `OK` or `KO`.

---

## Instruction Set

The implementation supports the standard push-swap operations:

```
pa pb
sa sb ss
ra rb rr
rra rrb rrr
```

The code operates on two stacks: `a` and `b`.

---

## Implementation Details

### Stack Model

* Two stacks: `a` and `b`.
* Operations are atomic and shared by both programs.

### Normalization

Values are converted to ranks (`0` to `n-1`) before sorting.
This simplifies comparisons and improves algorithm performance.

### Small Inputs (2–6 elements)

Special handling is used to minimize the number of instructions.

### Larger Inputs

A chunk-based strategy is used:

1. Normalize values to ranks.
2. Split the range into chunks.
3. Move elements between stacks based on chunk ranges.
4. Reassemble in sorted order.

This approach reduces the number of operations compared to a simple radix-based method.

---

## Testing

### Count Operations

```
ARG="<100 random numbers>"
./push-swap "$ARG" | wc -l
```

### Verify Correctness

```
ARG="<100 random numbers>"
./push-swap "$ARG" | ./checker "$ARG"
```

For consistent results, generate the input once and reuse it.

---

## Contributing

* The `helpers` package is small and modular.
* Sorting logic is split into `helpers/algo_*` files.
* Only Go standard library packages are used.

When modifying algorithms:

* Keep stack operations atomic.
* Maintain compatibility with the defined instruction set.
* Preserve error-handling behavior.

---

## License

This project is licensed under the **MIT License**.

You may use, copy, modify, merge, publish, distribute, sublicense, and sell copies of the software, subject to the terms of the MIT License.

A `LICENSE` file containing the full MIT license text should be included in the repository.
