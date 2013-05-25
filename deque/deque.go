// Circular sentinel Deque, culled/retyped from containers/list
// to start reading/writing golang properly. 

package deque

import (
	"math/rand"
	"bytes"
	"fmt"
)

type Value interface {}

type Node struct {
	// Liking this case=access feature.
	next, prev *Node
	deque      *Deque // Since we deal with nodes in user code
	
	// Public because we access the Val via the Node in user code:
	Val Value 
}

func (n *Node) Next() *Node {
	// Not so sure how well I like these if mini-statements:
	if node := n.next; node != &n.deque.sentinel {
		return node
	}
	return nil
}

func (n *Node) Prev() *Node {
	// Actually, pretty sure I hate them; why remove ternary op
	// and include/advertise this feature? Seems a bit arbitrary.
	if node := n.prev; node != &n.deque.sentinel {
		return node
	}
	return nil
}

type Deque struct {
	// Nice side effect of the sentinel as a member not a pointer
	// to one: the access syntax/semantics differ enough to remind
	// us of the fact.
	sentinel   Node
	size       int
}

func (d *Deque) Init() *Deque {
	// Nice that member/deref chaining doesn't look like ascii art
	// like in C/C++.
	d.sentinel.next = &d.sentinel
	d.sentinel.prev = &d.sentinel
	d.size = 0
	return d
}

func New() *Deque {
	return new(Deque).Init()
}

func (d *Deque) lazyInit() {
	if d.sentinel.next == nil {
		d.Init()
	}
}

func (d *Deque) insert(n, at *Node) *Node {
	// Bind the other soon-to-be-neighbor
	next      := at.next

	// If you like it then you'd better put some links on it
	at.next   = n
	n.prev    = at
	n.next    = next
	next.prev = n

	// Welcome to your new home, little node.
	n.deque   = d
	d.size++

	// Do be sure to tell us how you're getting on.
	return n
}

func (d *Deque) insertValue(v Value, at *Node) *Node {
	return d.insert(&Node{Val: v}, at)
}

func (d *Deque) remove(n *Node) *Node {
	// Magic Harry Potter spell to hide node
	n.prev.next = n.next
	n.next.prev = n.prev

        // Help a GC out
	n.next      = nil
	n.prev      = nil
	n.deque     = nil

	// http://www.youtube.com/watch?v=i2fkzL8pauM&t=2m10s
	d.size--
	return n
}

func (d *Deque) Remove(n *Node) Value {
	if n.deque == d {
		d.remove(n)
	}
	return n.Val
}
	
	
	
func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) Top() *Node {
	// inverting error filter in original code to error fallthrough
	if d.size > 0 {
		return d.sentinel.next
	}
	return nil
}

func (d *Deque) Bottom() *Node {
	if d.size > 0 {
		return d.sentinel.prev
	}
	return nil
}

func (d *Deque) PushTop(v Value) *Node {
	d.lazyInit()
	return d.insertValue(v, &d.sentinel)
}
	
func (d *Deque) PushBottom(v Value) *Node {
	d.lazyInit()
	return d.insertValue(v, d.sentinel.prev)
}

func (d *Deque) Shuffle() *Deque {
	if d.size < 2 {
		return d
	}

	// Read values into a slice
	var values []Value
	for node := d.sentinel.next; node != &d.sentinel; node = node.next {
		values = append(values, node.Val)
	}
	
	// Shuffle the slice
	// TODO: learn about randoms and seeds in Go; the test
	// currently generates the same data each run.
	for i := range values {
		j := rand.Intn(i+1)
		values[i], values[j] = values[j], values[i]
	}

	// Write them back
	node := d.sentinel.next
        for i := range values {
		node.Val = values[i]
		node = node.next
	}
	return d
}

func (d *Deque) String() string {
	var values bytes.Buffer
	for n, i := d.sentinel.next, 0; i<d.size; n, i = n.next, i+1 {
		values.WriteString(fmt.Sprintf("%v ", n.Val))
	}
	return values.String()
}
	
