package queue

type Queue struct {
	value string
	next  *Queue
}

var head *Queue
var tail *Queue
var length int

func New() *Queue {
	return &Queue{}
}

func Peek(q *Queue) string {
	return ""
}

func (q *Queue) Enqueue(value string) {
}

func (q *Queue) Dequeue() interface{} {
	return nil
}
