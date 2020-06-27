package implementDS

import (
	"fmt"
	"maze/models"
)

type MinHeap struct {
	HeapArray []models.GameQueue
	Size      int
	MaxSize   int
}

func NewMinHeap(maxsize int) *MinHeap {
	MinHeap := &MinHeap{
		HeapArray: []models.GameQueue{},
		Size:      0,
		MaxSize:   maxsize,
	}
	return MinHeap
}

func (m *MinHeap) leaf(index int) bool {
	if index >= (m.Size/2) && index <= m.Size {
		return true
	}
	return false
}

func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MinHeap) leftchild(index int) int {
	return 2*index + 1
}

func (m *MinHeap) rightchild(index int) int {
	return 2*index + 2
}

func (m *MinHeap) Insert(item models.GameQueue) error {
	if m.Size >= m.MaxSize {
		return fmt.Errorf("Heal is ful")
	}
	m.HeapArray = append(m.HeapArray, item)
	m.Size++
	m.upHeapify(m.Size - 1)
	return nil
}

func (m *MinHeap) swap(first, second int) {
	temp := m.HeapArray[first]
	m.HeapArray[first] = m.HeapArray[second]
	m.HeapArray[second] = temp
}

func (m *MinHeap) upHeapify(index int) {
	for m.HeapArray[index].Age < m.HeapArray[m.parent(index)].Age {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *MinHeap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}
	smallest := current
	leftChildIndex := m.leftchild(current)
	rightRightIndex := m.rightchild(current)
	//If current is smallest then return
	if leftChildIndex < m.Size && m.HeapArray[leftChildIndex].Age < m.HeapArray[smallest].Age {
		smallest = leftChildIndex
	}
	if rightRightIndex < m.Size && m.HeapArray[rightRightIndex].Age < m.HeapArray[smallest].Age {
		smallest = rightRightIndex
	}
	if smallest != current {
		m.swap(current, smallest)
		m.downHeapify(smallest)
	}
	return
}
func (m *MinHeap) BuildMinHeap() {
	for index := ((m.Size / 2) - 1); index >= 0; index-- {
		m.downHeapify(index)
	}
}

func (m *MinHeap) Remove() models.GameQueue {
	top := m.HeapArray[0]
	m.HeapArray[0] = m.HeapArray[m.Size-1]
	m.HeapArray = m.HeapArray[:(m.Size)-1]
	m.Size--
	m.downHeapify(0)
	return top
}
