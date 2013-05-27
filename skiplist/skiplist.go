// Package skiplist: a skiplist (slavishly) modeled on container/list
// "Skip Lists: A Probabilistic Alternative to Balanced Trees", W.
// Pugh, 1990.

package skiplist

import (
	//"math/rand"
)

type Element struct {
	Value   interface{}
	next []*Element
	parent  *Skiplist
}

func (e *Element) Next(level int) *Element {
	if len(e.next) < level {
		return nil
	}
	return e.next[level]
}

type Skiplist struct {
	root Element
	len, height int
}
	
func (sl *Skiplist) Init() *Skiplist {
	sl.root.next = []*Element{&sl.root}
	sl.len = 0
	sl.height = 0
	return sl
}

func New() *Skiplist {
	return new(Skiplist).Init()
}
	
func (sl *Skiplist) Len() int {
	return sl.len
}

// TODO: func (sl *Skiplist) Min() // easy
//       func (sl *Skiplist) Max() // bit more involved, or tail
//       ptr/cache?

func (sl *Skiplist) lazyInit() {
	if len(sl.root.next) == 0 {
		sl.Init()
	}
}

// TODO: func (sl *Skiplist) insert(e, at *Element) *Element {
	
