package stack

type Stack struct {
	datas []interface{}
	top   int
	n     int
}

func New(capacity int) *Stack {
	s := &Stack{}
	s.datas = make([]interface{}, capacity)
	return s
}

func (st *Stack) Push(val interface{}) {
	if st.full() {
		st.grow()
	}

	st.datas[st.top] = val
	st.top++
	st.n++
}

func (st *Stack) Pop() (val interface{}, ok bool) {
	if st.top == 0 {
		return
	}

	val = st.datas[st.top-1]
	st.top--
	st.n--
	ok = true
	return
}

func (st *Stack) Size() int {
	return st.n
}

func (st *Stack) Clear() {
	st.top = 0
	st.n = 0
}

func (st *Stack) full() bool {
	return st.top == len(st.datas)
}

func (st *Stack) grow() {
	newDatas := make([]interface{}, 2*len(st.datas))
	copy(newDatas, st.datas)
	st.datas = newDatas
}
