package stack

type Stack struct {
	Slice []int
}

func CreateStack(values ...int) *Stack {
	stck := &Stack{}

	for _, i := range values {
		stck.Push(i)
	}

	return stck
}

func (stck *Stack) Push(newElem int) {
	stck.Slice = append(stck.Slice, newElem)
}

func (stck *Stack) Pop() (int, bool) {
	if len(stck.Slice) == 0 {
		return 0, false
	} else {
		lastIndex := len(stck.Slice) - 1
		top := stck.Slice[lastIndex]
		stck.Slice = stck.Slice[:lastIndex]
		return top, true
	}
}

func (stck *Stack) IsEmpty() bool {
	return len(stck.Slice) == 0
}

func (stck *Stack) Size() int {
	return len(stck.Slice)
}

func (stck *Stack) Clear() {
	stck.Slice = stck.Slice[:0]
}
