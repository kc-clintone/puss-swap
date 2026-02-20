package helpers

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
