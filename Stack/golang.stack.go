package Stack

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	list := list.New()
	return &Stack{
		list: list,
	}
}

//入栈
func (p *Stack) Push(v interface{}) {
	p.list.PushBack(v)
}

//从栈顶出栈
func (p *Stack) Pop() (interface{}, bool) {
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		return e.Value, true
	}
	return nil, false
}

//获取栈顶元素
func (p *Stack) Top() (interface{}, bool) {
	e := p.list.Back()
	if e != nil {
		return e.Value, true
	}
	return nil, false
}

func (p *Stack) Len() int {
	return p.list.Len()
}

func (p *Stack) Empty() bool {
	return p.list.Len() == 0
}
