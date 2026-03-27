package fuse

import "testing"

func TestReadResultFdWithDoneInvokesCallbackOnce(t *testing.T) {
	calls := 0
	res := ReadResultFdWithDone(123, 456, 789, func() {
		calls++
	})

	res.Done()
	res.Done()

	if calls != 1 {
		t.Fatalf("Done callback calls = %d, want 1", calls)
	}
}
