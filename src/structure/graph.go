package structure

import "minimum_spanning_tree/mst_util"

type Graph interface {
	Get_adjacent_list() *Vertex_list
	Are_adjcent() bool
}

type Vertex uint32

type Vertex_list struct {
	V    Vertex
	Next *Vertex_list
}

func (v_list Vertex_list) Value() Vertex {
	return v_list.V
}

func (v_list Vertex_list) Next_list() mst_util.Linked_list[Vertex] {
	if v_list.Next == nil {
		return nil
	} else {
		return v_list.Next
	}
}

type Adjacency_structure []*Vertex_list

func (s *Adjacency_structure) Get_adjacent_list(v Vertex) *Vertex_list {
	if int(v) > len(*s) {
		return nil
	} else {
		return (*s)[int(v)]
	}
}

func (s *Adjacency_structure) Are_adjacent(v, w Vertex) bool {
	return mst_util.Search(*s.Get_adjacent_list(v), w)
}
