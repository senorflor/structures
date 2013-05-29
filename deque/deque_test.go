package deque

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
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
	d := New().PushTop(42).PushBottom(1337)

	if v := d.Top(); v != 42 {
		t.Errorf("Top was not correct after PushTop, PushBottom")
	}

	if v := d.Bottom(); v != 1337 {
		t.Errorf("Bottom was not correct after PushTop, PushBottom")
	}
}

func TestPop(t *testing.T) {
	d := New().PushTop(2).PushBottom(3).PushTop(1) // 1 2 3
	if v := d.PopBottom(); v != 3 {
		t.Errorf("PopBottom returned %d, not 3", v)
	}
	if v := d.PopTop(); v != 1 {
		t.Errorf("PopTop returned %d, not 1", v)
	}
	if s := d.Size(); s != 1 {
		t.Errorf("After pops, deque should have size 1, not %d", s)
	}
}

func TestShuffleTinyDeques(t *testing.T) {
	d1, d2 := New(), New()
	if !reflect.DeepEqual(d1, d2.Shuffle()) {
		t.Errorf("Shuffle was not id on the empty deque")
	}
	d1.PushTop(25)
	d2.PushTop(25)
	d2.Shuffle()
	if !reflect.DeepEqual(d1, d2.Shuffle()) {
		t.Errorf("Shuffle was not id on deque of size 1")
	}
}

func TestShuffleLargerDeque(t *testing.T) {
	// How do you test a random "function"?
	// Statistics on number of inversions in repeatedly shuffled
	// ranges? Let's just print out a shuffled deque for now.
	d := New()
	for i := 1; i <= 20; i++ {
		d.PushBottom(i)
	}

	// `go test -test.v [...]` to get printed output even if
	// tests are successful
	t.Logf("Before Shuffle(): %s", d.String())
	rand.Seed(time.Now().UnixNano())
	d.Shuffle()
	t.Logf("After Shuffle():  %s", d.String())
}
