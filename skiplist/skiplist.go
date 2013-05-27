// Package skiplist: a skiplist (slavishly) modeled on container/list.
// See Pugh, W. "Skip Lists: A Probabilistic Alternative to Balanced
// Trees", W. Pugh, CACM, 33:6 (June 1990).
package skiplist

import (
	//"container/list"
	//"math/rand"
)

// Datum _cum_ forward skip iterator
type Element struct {
	next  []*Element
	sl    *Skiplist
	Value interface{}
}

// Iterate forward at a given skip level
func (e *Element) Next(lev int) *Element {
	if len(e.next) < lev || lev < 0 || e.next[lev] == &e.sl.root {
		return nil
	}
	return e.next[lev]
}

// Finds the next non-root element at the highest possible level
func (e *Element) SkipForward(max int) *Element {
	for lev := len(e.next)-1; lev >= 0; lev-- {
		if n := e.Next(lev); n != &e.sl.root {
			return n
		}
	}
	return nil 
}

type Skiplist struct {
	root Element
	len int
}
	
func New() *Skiplist {
	return new(Skiplist).Init()
}

func (sl *Skiplist) Init() *Skiplist {
	sl.root.next = []*Element{&sl.root}
	sl.len = 0
	return sl
}

func (sl *Skiplist) lazyInit() {
	if len(sl.root.next) == 0 {
		sl.Init()
	}
}
	
func (sl *Skiplist) Len() int {
	return sl.len
}

func (sl *Skiplist) Min() *Element {
	n := sl.root.Next(0)
	if n == &sl.root {
		return nil
	}
	return n;
}
// TODO: func (sl *Skiplist) Max() // bit more involved, or tail
//       ptr/cache?

// STUB
func (sl *Skiplist) insert(e, at *Element) *Element {
        //height := randomLevel(sl.len)
	//if len(sl.next)
	return e 
}

// STUB
func randomLevel(len int) int {
	return 4 
}

// TODO: factor max out into some kind of util package. I really miss
// generics/ternary ops here; what's the non-casting alternative? c.f.
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/auLWUqourlA%5B1-25-false%5D
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
