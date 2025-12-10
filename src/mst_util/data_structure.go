package mst_util

type Linked_list[T comparable] interface {
	Value() T
	Next_list() Linked_list[T]
}

func Search[T comparable](ll Linked_list[T], e T) bool {
	if ll == nil {
		return false
	} else {
		return Search(ll.Next_list(), e)
	}
}
