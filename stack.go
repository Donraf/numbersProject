package main

type Stack struct {
	list []string
}

func (st *Stack) Append(s string) {
	(*st).list = append((*st).list[:], s)
}

func (st *Stack) Pop() string {
	length := len((*st).list)
	last := (*st).list[length-1]
	(*st).list = (*st).list[:length-1]
	return last
}

func (st *Stack) Get(index int) string {
	return (*st).list[index]
}

func (st *Stack) Len() int {
	return len((*st).list)
}
