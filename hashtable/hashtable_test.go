package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertGetDelete(t *testing.T, actual, expected interface{}, okActual, okExpected bool) {
	assert.Equal(t, expected, actual, "nil is an expected value")
	assert.Equal(t, okExpected, okActual, "false is an expected value")
}

func TestHashTable(t *testing.T) {
	hashTable := New()

	data, ok := hashTable.Get(1)
	assertGetDelete(t, data, nil, ok, false)

	data, ok = hashTable.Delete(1)
	assertGetDelete(t, data, nil, ok, false)

	hashTable.Set(1, "1")
	hashTable.Set("2", "2")
	hashTable.Set("4", "4")

	data, ok = hashTable.Get(1)
	assertGetDelete(t, data, "1", ok, true)
	data, ok = hashTable.Delete(1)
	assertGetDelete(t, data, "1", ok, true)
	data, ok = hashTable.Delete(1)
	assertGetDelete(t, data, nil, ok, false)

	data, ok = hashTable.Get("4")
	assertGetDelete(t, data, "4", ok, true)
	data, ok = hashTable.Delete("4")
	assertGetDelete(t, data, "4", ok, true)
	data, ok = hashTable.Delete("4")
	assertGetDelete(t, data, nil, ok, false)

	data, ok = hashTable.Get("2")
	assertGetDelete(t, data, "2", ok, true)
	hashTable.Set("2", "5")
	data, ok = hashTable.Get("2")
	assertGetDelete(t, data, "5", ok, true)
}
