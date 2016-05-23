package stack

import "testing"

func TestStack(t *testing.T) {
	st := New(10)

	for k := 1; k <= 10; k++ {
		for i := 0; i < k; i++ {
			st.Push(i)
		}

		if st.Size() != k {
			t.Errorf("Size")
		}

		for i := k - 1; i >= 0; i-- {
			if val, ok := st.Pop(); !ok || val.(int) != i {
				t.Errorf("Pop")
			}
		}

		if st.Size() != 0 {
			t.Errorf("Size")
		}

		if _, ok := st.Pop(); ok {
			t.Errorf("Pop")
		}
	}

	for k := 1; k <= 20; k++ {
		for i := 0; i < 2*k; i++ {
			st.Push(i)
		}

		if st.Size() != 2*k {
			t.Errorf("Size")
		}

		for i := 2*k - 1; i >= 0; i-- {
			if val, ok := st.Pop(); !ok || val.(int) != i {
				t.Errorf("Pop")
			}
		}

		if st.Size() != 0 {
			t.Errorf("Size")
		}

		if _, ok := st.Pop(); ok {
			t.Errorf("Pop")
		}
	}
}
