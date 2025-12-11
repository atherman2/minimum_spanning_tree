package algorithms

import (
	st "minimum_spanning_tree/structure"
	"sort"
)

func Core_kruskal(g st.Graph, f st.Vertex_set_family) (st.Graph, st.Weight) {
	total_weight := st.Weight(0.0)
	t := st.Adjacency_structure{N: g.Get_n(), Edges: []st.Edge{}, Structure: make([]st.Vertex_list, g.Get_n())}
	e := append([]st.Edge{}, g.Get_edges()...)
	sort.Slice(
		e, func(i, j int) bool {
			return e[i].Weight < e[j].Weight
		},
	)
	for _, edge := range e {
		if f.Find_root(edge.V) != f.Find_root(edge.W) {
			total_weight += edge.Weight
			t.Add_edge(edge)
			f.Unite_sets(edge.V, edge.W)
		}
	}
	return &t, total_weight
}

func Index_vector_set_family_kruskal(g st.Graph) (st.Graph, st.Weight) {
	f := st.New_index_vector_set_family(g.Get_n())
	return Core_kruskal(g, &f)
}

func Tree_set_family_kruskal(g st.Graph) (st.Graph, st.Weight) {
	f := st.New_tree_set_family(g.Get_n())
	return Core_kruskal(g, &f)
}
