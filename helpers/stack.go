package helpers

// Stack represents a stack of integers with basic operations to manipulate it. It provides methods to create a new stack, check its length, and determine if it's empty. This struct is fundamental for implementing the push_swap algorithm, as it allows us to manage the two stacks (a and b) used in the sorting process.

type Stack struct {
	Data []int
}

func NewStack(nums []int) *Stack {
	return &Stack{Data: nums}
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}
