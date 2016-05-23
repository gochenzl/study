package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New(10)

	for k := 1; k <= 10; k++ {
		for i := 0; i < k; i++ {
			q.Push(i)
		}

		if q.Size() != k {
			t.Errorf("Size")
		}

		for i := 0; i < k; i++ {
			if val, ok := q.Pop(); !ok || val.(int) != i {
				t.Errorf("Pop")
			}
		}

		if q.Size() != 0 {
			t.Errorf("Size")
		}

		if _, ok := q.Pop(); ok {
			t.Errorf("Pop")
		}
	}

	for k := 1; k <= 20; k++ {
		for i := 0; i < k*2; i++ {
			q.Push(i)
		}

		if q.Size() != k*2 {
			t.Errorf("Size")
		}

		for i := 0; i < k*2; i++ {
			if val, ok := q.Pop(); !ok || val.(int) != i {
				t.Errorf("Pop")
			}
		}

		if q.Size() != 0 {
			t.Errorf("Size")
		}

		if _, ok := q.Pop(); ok {
			t.Errorf("Pop")
		}
	}

}
