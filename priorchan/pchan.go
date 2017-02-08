package priorchan

import (
	"sync"
)

type Element struct {
	Priority int
	Content  string
}

type Elements struct {
	lock     *sync.Mutex
	elements []*Element
}

func NewElements() *Elements {
	return &Elements{&sync.Mutex{}, make([]*Element, 0, 30)}
}

func (e *Elements) Push(ele *Element) {
	e.lock.Lock()
	defer e.lock.Unlock()
	length := len(e.elements)
	if length <= 0 {
		e.elements = append(e.elements, ele)
		return
	}
	pos := e.search(0, length-1, ele)
	if pos >= length {
		e.elements = append(e.elements, ele)
	} else {
		e.elements = append(e.elements[:pos+1], e.elements[pos:]...)
		e.elements[pos] = ele
	}
}

func (e *Elements) search(start, end int, ele *Element) int {
	middle := (start + end) / 2
	if end < start {
		return start
	}
	if ele.Priority > e.elements[middle].Priority {
		return e.search(start, middle-1, ele)
	}
	return e.search(middle+1, end, ele)
}

func (e *Elements) Pop() (*Element, bool) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if len(e.elements) > 0 {
		msg := e.elements[0]
		e.elements = e.elements[1:]
		return msg, true
	} else {
		return &Element{999, "empty"}, false
	}
}
