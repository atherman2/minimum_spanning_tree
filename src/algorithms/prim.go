package algorithms

import (
	"container/heap"
	st "minimum_spanning_tree/structure"
)

// Algoritmo de Prim
func Core_prim(g st.Graph, start st.Vertex) (st.Graph, st.Weight) {
	n := g.Get_n()
	visited := make([]bool, n+1) // marca vértices já incluídos
	totalWeight := st.Weight(0.0)

	// grafo resultante (MST)
	t := st.Adjacency_structure{
		N:         n,
		Edges:     []st.Edge{},
		Structure: make([]st.Vertex_list, n),
	}

	pq := &st.PriorityQueue{}
	heap.Init(pq)

	// começa pelo vértice inicial
	visited[start] = true
	adj := g.Get_adjacent_list(start)
	for adj != nil {
		heap.Push(pq, st.EdgeItem{Edge: st.Edge{V: start, W: adj.Key(), Weight: adj.Data()}})
		adj = adj.Next_list()
	}

	// enquanto houver arestas na fila
	for pq.Len() > 0 {
		item := heap.Pop(pq).(st.EdgeItem)
		e := item.Edge

		if visited[e.W] {
			continue
		}

		// adiciona aresta à MST
		t.Add_edge(e)
		t.Edges = append(t.Edges, e)
		totalWeight += e.Weight
		visited[e.W] = true

		// adiciona novas arestas do vértice recém incluído
		adj := g.Get_adjacent_list(e.W)
		for adj != nil {
			if !visited[adj.Key()] {
				heap.Push(pq, st.EdgeItem{Edge: st.Edge{V: e.W, W: adj.Key(), Weight: adj.Data()}})
			}
			adj = adj.Next_list()
		}
	}

	return &t, totalWeight
}
