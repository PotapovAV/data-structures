package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyStackPop(t *testing.T) {
	stack := New(2)
	assert.Equal(t, nil, stack.Pop(), "nil should be returned by stack.Pop()")
}

func TestPush(t *testing.T) {
	stack := New(2)
	assert.Equal(t, 1, stack.Push(1), "1 value should be returned by stack.Push(1)")
	assert.Equal(t, 2, stack.Push(2), "2 value should be returned by stack.Push(2)")
}

func TestFullStackPush(t *testing.T) {
	stack := New(2)
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, nil, stack.Push(3), "nil should be returned by stack.Push(3)")
}

func TestPop(t *testing.T) {
	stack := New(2)
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Pop(), "2 should returned by stack.Pop()")
	assert.Equal(t, 1, stack.Pop(), "1 should returned by stack.Pop()")
}
