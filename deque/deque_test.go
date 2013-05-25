package deque

import (
	"math/rand"
	"time"
	"testing"
	"reflect"
)

func TestDequeCreation(t *testing.T) {
	d := New()
	if d.Top() != nil {
		t.Errorf("Empty deque had a non-nil Top()")
	}
	if d.Bottom() != nil {
		t.Errorf("Empty deque had a non-nil Bottom()")
	}
}

func TestPush(t *testing.T) {
	d := New()
	top := d.PushTop(42)
	bottom := d.PushBottom(1337)
	if n := d.Top(); n.Val != 42 || top != n {
		t.Errorf("Top was not correct after PushTop, PushBottom")
	}
	if n := d.Bottom(); n.Val != 1337 || bottom != n {
		t.Errorf("Bottom was not correct after PushTop, PushBottom")
	}
}

func TestShuffleTinyDeques(t *testing.T) {
	d1, d2 := New(), New()
	if ! reflect.DeepEqual(d1, d2.Shuffle())  {
		t.Errorf("Shuffle was not id on the empty deque")
	}
	d1.PushTop(25)
	d2.PushTop(25)
	d2.Shuffle()
	if ! reflect.DeepEqual(d1, d2.Shuffle())  {
		t.Errorf("Shuffle was not id on deque of size 1")
	}
}

func TestShuffleLargerDeque(t *testing.T) {
	// How do you test a random "function"?
	// Statistics on number of inversions in repeatedly shuffled
	// ranges? Let's just print out a shuffled deque for now.
	d := New()
	for i := 1; i<=20; i++ {
		d.PushBottom(i)
	}
	
	// `go test -test.v [...]` to get printed output even if 
	// tests are successful
	t.Logf("Before Shuffle(): %s", d.String())
	rand.Seed(time.Now().UnixNano())
	d.Shuffle()
	t.Logf("After Shuffle():  %s", d.String())
}
