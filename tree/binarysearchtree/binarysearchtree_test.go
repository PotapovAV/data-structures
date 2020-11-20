package binarysearchtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	key     int
	payload string
}

func (d testData) LessThan(other interface{}) bool {
	switch other.(type) {
	case int:
		if d.key < other.(int) {
			return true
		}
	case testData:
		if d.key < other.(testData).key {
			return true
		}
	default:
		fmt.Println("Provided type is not supported")
	}
	return false
}

func (d testData) EqualsTo(other interface{}) bool {
	switch other.(type) {
	case int:
		key := other.(int)
		if d.key == key {
			return true
		}
	case string:
		payload := other.(string)
		if d.payload == payload {
			return true
		}
	case testData:
		data := other.(testData)
		if d == data {
			return true
		}
	default:
		fmt.Println("Provided type is not supported")
	}
	return false
}

func createTestData(key int, payload string) testData {
	return testData{
		key:     key,
		payload: payload,
	}
}

func TestBinarySearchTree(t *testing.T) {
	tree := New()
	node := tree.Search(1)
	assert.Equal(t, (*Node)(nil), node, "nil should be returned by tree.Search(1)")

	a := []int{2, 1, 3, 4, 6, 9, 0, 5}
	for _, value := range a {
		tree.Insert(createTestData(value, "test"))
	}

	node = tree.Search(2)
	assert.Equal(t, 2, node.Data().(testData).key, "2 should be returned by tree.Search(2)")
	node = tree.Search(7)
	assert.Equal(t, (*Node)(nil), node, "nil should be returned by tree.Search(7)")

	node = tree.Search(3)
	tree.Delete(node)
	node = tree.Search(3)
	assert.Equal(t, (*Node)(nil), node, "nil should be returned by tree.Search(3)")

	node = tree.Minimum()
	assert.Equal(t, 0, node.Data().(testData).key, "0 should be returned by tree.Minimum()")

	node = tree.Maximum()
	assert.Equal(t, 9, node.Data().(testData).key, "9 should be returned by tree.Maximum()")

	node = tree.Search(4)
	node = node.Successor()
	assert.Equal(t, 5, node.Data().(testData).key, "5 should be returned by node.Successor()")

	node = tree.Search(4)
	node = node.Predecessor()
	assert.Equal(t, 2, node.Data().(testData).key, "2 should be returned by node.Predecessor()")
}
