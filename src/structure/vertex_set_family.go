package structure

type VertexSetFamily interface {
	Find_root(v Vertex) Vertex
	Unite_sets(v, w Vertex)
}
