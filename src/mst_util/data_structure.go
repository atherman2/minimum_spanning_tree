package mst_util

type Linked_list[T comparable, U any] interface {
	Key() T
	Next_list() Linked_list[T, U]
	Data() U
}

func Search[T comparable, U any](ll Linked_list[T, U], e T) Linked_list[T, U] {
	if ll == nil {
		return nil
	} else if ll.Key() == e {
		return ll
	} else {
		return Search(ll.Next_list(), e)
	}
}
