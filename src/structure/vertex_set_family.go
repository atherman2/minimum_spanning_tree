package structure

type Vertex_set_family interface {
	Find_root(v Vertex) Vertex
	Unite_sets(v, w Vertex)
}
