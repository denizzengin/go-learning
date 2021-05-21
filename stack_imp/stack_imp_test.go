package stack_imp

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)


func Test_Stack(t *testing.T){
	s := Stack{}
	s.Push(1)
	fmt.Println(s)
	assert.Equal(t, len(s.arr), 1)
	assert.Equal(t, s.length, 1)
	s.Push(2)
	fmt.Println(s)
	assert.Equal(t, len(s.arr), 2)
	assert.Equal(t, s.length, 2)
	last := s.Pop()
	fmt.Println(s)
	assert.Equal(t, last, 2)
	assert.Equal(t, s.length, 1)

	stackFixedSize := StackFixedSize{}
	stackFixedSize.Push(5)
	fmt.Println(stackFixedSize)
	last2 := stackFixedSize.Pop()
	fmt.Println(last2)
	fmt.Println(stackFixedSize)
}