package structure

import "minimum_spanning_tree/mst_util"

type Graph interface {
	Get_adjacent_list() *Vertex_linked_list
	Are_adjcent() bool
}

type Vertex uint32
type Weight float64

type Vertex_linked_list struct {
	V      Vertex
	Next   *Vertex_linked_list
	Weight Weight
}

type Vertex_list mst_util.Linked_list[Vertex, Weight]

func (v_list Vertex_linked_list) Key() Vertex {
	return v_list.V
}

func (v_list Vertex_linked_list) Next_list() mst_util.Linked_list[Vertex, Weight] {
	if v_list.Next == nil {
		return nil
	} else {
		return v_list.Next
	}
}

func (v_list Vertex_linked_list) Data() Weight {
	return v_list.Weight
}

type Adjacency_structure []*Vertex_linked_list

func (s *Adjacency_structure) Get_adjacent_list(v Vertex) Vertex_list {
	if int(v) > len(*s) {
		return nil
	} else {
		return (*s)[int(v)]
	}
}

func (s *Adjacency_structure) Get_edge(v, w Vertex) Vertex_list {
	return mst_util.Search(s.Get_adjacent_list(v), w)
}
