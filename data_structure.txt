package mst_util

type Linked_list[T comparable, U any] interface {
	Key() T
	Next_list() Linked_list[T, U]
	Data() U
	Set_next(Linked_list[T, U])
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

func LL_insert[T comparable, U any](old, new Linked_list[T, U]) Linked_list[T, U] {
	new.Set_next(old)
	return new
}
