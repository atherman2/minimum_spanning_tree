package structure

import mu "minimum_spanning_tree/mst_util"

type Graph interface {
	Get_adjacent_list(v Vertex) Vertex_list
	Get_edge(v, w Vertex) Vertex_list
	Add_edge(e Edge)
}

type Vertex uint32
type Weight float64

type Vertex_linked_list struct {
	V      Vertex
	Next   *Vertex_linked_list
	Weight Weight
}

type Vertex_list mu.Linked_list[Vertex, Weight]

func (v_list Vertex_linked_list) Key() Vertex {
	return v_list.V
}

func (v_list Vertex_linked_list) Next_list() mu.Linked_list[Vertex, Weight] {
	if v_list.Next == nil {
		return nil
	} else {
		return v_list.Next
	}
}

func (v_list Vertex_linked_list) Data() Weight {
	return v_list.Weight
}

func (v_list *Vertex_linked_list) Set_next(next mu.Linked_list[Vertex, Weight]) {
	v_list.Next = next.(*Vertex_linked_list)
}

type Adjacency_structure []Vertex_list

func (s Adjacency_structure) Get_adjacent_list(v Vertex) Vertex_list {
	if int(v) > len(s) {
		return nil
	} else {
		return s[int(v)]
	}
}

func (s Adjacency_structure) Get_edge(v, w Vertex) Vertex_list {
	return mu.Search(s.Get_adjacent_list(v), w)
}

type Edge struct {
	V, W   Vertex
	Weight Weight
}

func New_edge(v, w int32, weight float64) Edge {
	return Edge{V: Vertex(v), W: Vertex(w), Weight: Weight(weight)}
}

func (s *Adjacency_structure) Add_edge(e Edge) {
	mu.LL_insert((*s)[e.V], &Vertex_linked_list{V: e.W, Weight: e.Weight})
	mu.LL_insert((*s)[e.W], &Vertex_linked_list{V: e.V, Weight: e.Weight})
}

func Adjacency_structure_graph_from_edges(n int, edges []Edge) Graph {
	result := make(Adjacency_structure, n)
	for _, e := range edges {
		result.Add_edge(e)
	}
	return &result
}
