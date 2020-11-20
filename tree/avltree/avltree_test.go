package avltree

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
		if d.key == other.(int) {
			return true
		}
	case string:
		if d.payload == other.(string) {
			return true
		}
	case testData:
		if d == other.(testData) {
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

func TestAVLTree(t *testing.T) {
	tree := New()

	tree.Insert(createTestData(1, "one"))
	tree.Insert(createTestData(2, "two"))
	tree.Insert(createTestData(3, "three"))
	tree.Insert(createTestData(4, "four"))
	tree.Insert(createTestData(5, "five"))
	tree.Insert(createTestData(6, "six"))

	sorted := tree.InorderTreeWalk()
	assert.Equal(t, 6, len(sorted), "6 should be the length of sorted slice")

	node := tree.Search(7)
	assert.Equal(t, (*Node)(nil), node, "nil should be returned by tree.Search(7)")

	node = tree.Search(4)
	assert.Equal(t, 4, node.Data().(testData).key, "4 should be returned by tree.Search(4)")

	node = tree.Minimum()
	assert.Equal(t, 1, node.Data().(testData).key, "1 should be returned by tree.Minimum()")

	node = tree.Maximum()
	assert.Equal(t, 6, node.Data().(testData).key, "6 should be returned by tree.Maximum()")

	tree.Delete(tree.Search(5))
	node = tree.Search(5)
	assert.Equal(t, (*Node)(nil), node, "nil should be returned by tree.Search(5)")

	node = tree.Search(4)
	node = node.Successor()
	assert.Equal(t, 6, node.Data().(testData).key, "6 should be returned by node.Successor()")

	node = tree.Search(4)
	node = node.Predecessor()
	assert.Equal(t, 3, node.Data().(testData).key, "3 should be returned by node.Predecessor()")

	tree.Delete(tree.Search(1))
	tree.Delete(tree.Search(2))
	tree.Delete(tree.Search(3))
	tree.Delete(tree.Search(4))
	tree.Delete(tree.Search(6))

	sorted = tree.InorderTreeWalk()
	assert.Equal(t, 0, len(sorted), "0 should be the length of sorted slice")
}
