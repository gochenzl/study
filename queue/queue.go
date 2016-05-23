package queue

type Queue struct {
	datas      []interface{}
	head, tail int
	n          int
}

func New(capacity int) *Queue {
	q := &Queue{}
	q.datas = make([]interface{}, capacity)
	return q
}

func (q *Queue) Push(val interface{}) {
	if q.n >= len(q.datas) { // full
		q.grow()
	}

	q.datas[q.tail] = val
	q.n++
	q.tail++
	if q.tail >= len(q.datas) {
		q.tail = 0
	}
}

func (q *Queue) Pop() (val interface{}, ok bool) {
	if q.n == 0 {
		return
	}

	val = q.datas[q.head]
	q.n--
	q.head++
	if q.head >= len(q.datas) {
		q.head = 0
	}

	ok = true
	return
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Clear() {
	q.head = 0
	q.tail = 0
	q.n = 0
}

// called when full
func (q *Queue) grow() {
	newDatas := make([]interface{}, len(q.datas)*2)

	j := 0
	for i := q.head; i < len(q.datas); i++ {
		newDatas[j] = q.datas[i]
		j++
	}

	for i := 0; i < q.head; i++ {
		newDatas[j] = q.datas[i]
		j++
	}

	q.head = 0
	q.tail = len(q.datas)
	q.datas = newDatas
}
