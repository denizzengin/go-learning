package stack_imp 


type Stack struct {
	length int
	arr []int
}

func (s *Stack) Push(e int) {

	newArr := make([]int, len(s.arr)+1)
	newArr[0] = e

	for i := 1; i < len(newArr); i++ {
		newArr[i] = s.arr[i-1]
	}
	s.arr = newArr
	s.length = len(newArr)
}

func (s *Stack) Pop() int {

	if len(s.arr) == 0 {
		return -1
	}

	lastAdded := s.arr[0];
	newArr := make([]int, len(s.arr)-1)
	for i := 0; i < len(s.arr)-1; i++ {
		newArr[i] = s.arr[i+1]
	}
	s.length = len(newArr)
	s.arr = newArr
	return lastAdded 	
}

type StackFixedSize struct {
	size int
	data []int
}

func (s *StackFixedSize) Push(e int) {
	if s.size == 0 {
		s.data = make([]int, 10)
	}
	s.data[s.size] = e;
	s.size++
}
func (s *StackFixedSize) Pop() int {
	if s.size <= 0 {
		return -1
	}
	last := s.data[s.size-1]	
	s.data[s.size-1] = 0
	s.size--
	return last
}
