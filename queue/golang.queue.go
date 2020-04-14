package queue

import "errors"

type Queue []interface{}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x)
}

func (q *Queue) Pop() (interface{}, error) {
	if len(*q) == 0 {
		return nil, errors.New("Can't pop an empty queue")
	}
	top := (*q)[0]
	*q = (*q)[1:]
	return top, nil
}
func (q Queue) Top() (interface{}, error) {
	if len(q) == 0 {
		return nil, errors.New("Can't top an empty queue")
	}
	return q[0], nil
}
func (q Queue) IsEmpty() bool {
	return (len(q) == 0)
}

func (q Queue) Len() int {
	return len(q)
}
