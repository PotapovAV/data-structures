package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyList(t *testing.T) {
	list := New()
	assert.Equal(t, nil, list.GetFirst(), "nil should be returned by list.GetFirst()")
	assert.Equal(t, nil, list.GetLast(), "nil should be returned by list.GetLast()")
}

func TestInsertSearchRemove(t *testing.T) {
	list := New()

	for i := 0; i < 5; i++ {
		list.Insert(i)
	}
	list.InsertFirst(5)
	list.InsertLast(6)

	unitPtr := list.Search(7)
	assert.Equal(t, nil, list.Remove(unitPtr), "nil should be returned by list.Remove(unitPtr)")
	assert.Equal(t, 5, list.GetFirst(), "5 should be returned by list.GetFirst()")
	assert.Equal(t, 6, list.GetLast(), "6 should be returned by list.GetLast()")

	unitPtr = list.Search(2)
	assert.Equal(t, 2, list.Remove(unitPtr), "2 should be returned by list.Remove(unitPtr)")
	assert.Equal(t, 0, list.GetLast(), "0 should be returned by list.GetLast()")
	assert.Equal(t, 1, list.GetLast(), "1 should be returned by list.GetLast()")
	assert.Equal(t, 3, list.GetLast(), "3 should be returned by list.GetLast()")
	assert.Equal(t, 4, list.GetFirst(), "4 should be returned by list.GetFirst()")

	assert.Equal(t, nil, list.GetLast(), "nil should be returned by list.GetLast()")
	assert.Equal(t, nil, list.GetFirst(), "nil should be returned by list.GetFirst()")
}
