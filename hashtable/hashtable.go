package hashtable

import (
	"fmt"
	"math"

	"github.com/PotapovAV/data-structures/linkedlist"
)

const (
	hashTableSize = 997

	a = 3529
	b = 3539
	p = 3571
)

type HashTable struct {
	hashTable []*linkedlist.List
}

type Node struct {
	key  interface{}
	data interface{}
}

func New() *HashTable {
	return &HashTable{
		hashTable: make([]*linkedlist.List, hashTableSize),
	}
}

func index(key interface{}) int {
	return hash(key) % hashTableSize
}

func hash(key interface{}) int {
	switch key.(type) {
	case string:
		return hashString(key.(string))
	case int:
		return hashInt(key.(int))
	default:
		fmt.Println("Provided type of key is not acceptable")
		return -1
	}
}

func hashString(key string) int {
	h := 0
	for idx, char := range key {
		h += int(char) * int(math.Pow(31, float64(len(key)-idx+1)))
	}
	return h % p
}

func hashInt(key int) int {
	return (a*key + b) % p
}

func nodeCreate(key interface{}, data interface{}) Node {
	return Node{
		key:  key,
		data: data,
	}
}

func (h *HashTable) Set(key interface{}, data interface{}) {
	idx := index(key)
	if idx == -1 {
		return
	}
	if h.hashTable[idx] == nil {
		h.hashTable[idx] = linkedlist.New()
		h.hashTable[idx].Insert(nodeCreate(key, data))
	} else {
		linkedList := h.hashTable[idx]
		nextNodePtr := linkedList.Head()
		for nextNodePtr != nil {
			nodeData := linkedList.Data(nextNodePtr).(Node)
			if nodeData.key == key {
				linkedList.Update(nextNodePtr, nodeCreate(key, data))
				return
			}
			nextNodePtr = linkedList.Next(nextNodePtr)
		}
		linkedList.Insert(nodeCreate(key, data))
	}
}

func (h *HashTable) Get(key interface{}) (interface{}, bool) {
	idx := index(key)
	if idx == -1 {
		return nil, false
	}
	linkedList := h.hashTable[idx]
	if linkedList == nil {
		return nil, false
	}
	nextNodePtr := linkedList.Head()
	for nextNodePtr != nil {
		nodeData := linkedList.Data(nextNodePtr).(Node)
		if nodeData.key == key {
			return nodeData.data, true
		}
		nextNodePtr = linkedList.Next(nextNodePtr)
	}
	return nil, false
}

func (h *HashTable) Delete(key interface{}) (interface{}, bool) {
	idx := index(key)
	if idx == -1 {
		return nil, false
	}
	linkedList := h.hashTable[idx]
	if linkedList == nil {
		return nil, false
	}
	nextNodePtr := linkedList.Head()
	for nextNodePtr != nil {
		nodeData := linkedList.Data(nextNodePtr).(Node)
		if nodeData.key == key {
			linkedList.Remove(nextNodePtr)
			return nodeData.data, true
		}
		nextNodePtr = linkedList.Next(nextNodePtr)
	}
	return nil, false
}
