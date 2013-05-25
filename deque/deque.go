// Encapsulated Deque, based on container/list

package deque

import (
	"bytes"
	"container/list"
	"fmt"
	"math/rand"
)

type Deque struct {
	l *list.List
	last *list.Element // future access caching hook
}

func (d *Deque) Init() *Deque {
	d.l = list.New()
	return d
}

func New() *Deque {
	return new(Deque).Init()
}

func (d *Deque) Size() int {
	return d.l.Len()
}

func (d *Deque) Top() interface{} {
	d.last = d.l.Front()
	if d.last == nil {
		return nil
	}
	return d.last.Value
}

func (d *Deque) Bottom() interface{} {
	d.last = d.l.Back()
	if d.last == nil {
		return nil
	}
	return d.last.Value
}

func (d *Deque) PushTop(v interface {}) *Deque {
	d.last = d.l.PushFront(v)
	return d 
}
	
func (d *Deque) PushBottom(v interface {}) *Deque {
	d.last = d.l.PushBack(v)
	return d 
}

func (d *Deque) PopTop() interface{} {
	if d.Size() < 1 {
		return nil
	}
	return d.l.Remove(d.l.Front())
}

func (d *Deque) PopBottom() interface{} { // he he
	if d.Size() < 1 {
		return nil
	}
	return d.l.Remove(d.l.Back())
}

func (d *Deque) Shuffle() *Deque {
	if d.Size() < 2 {
		return d
	}

	// TODO: Should lock deque during this, or figure out a more
	// in-place shuffle.

	// Read values into a slice
	var values []interface {}
	for n, i := d.l.Front(), 0; i<d.Size(); n, i = n.Next(), i+1 {
		values = append(values, n.Value)
	}
	
	// Shuffle the slice
	for i := range values {
		j := rand.Intn(i+1)
		values[i], values[j] = values[j], values[i]
	}

	// Write them back
	n := d.l.Front()
        for i := range values {
		n.Value = values[i]
		n = n.Next()
	}
	return d
}

func (d *Deque) String() string {
	var values bytes.Buffer
	for n, i := d.l.Front(), 0; i<d.Size(); n, i = n.Next(), i+1 {
		values.WriteString(fmt.Sprintf("%v ", n.Value))
	}
	return values.String()
}
	
