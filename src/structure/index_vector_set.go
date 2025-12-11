package structure

type Index_vector_set_family []Vertex

func (f Index_vector_set_family) Find_root(v Vertex) Vertex {
	for v != f[v-1] {
		v = f[v-1]
	}
	return v
}

func (f *Index_vector_set_family) Unite_sets(v, w Vertex) {
	root_v := f.Find_root(v)
	root_w := f.Find_root(w)
	(*f)[root_w-1] = root_v
}

func New_index_vector_set_family(n int) Index_vector_set_family {
	result := make(Index_vector_set_family, n)
	for i := range n {
		result[i] = Vertex(i + 1)
	}
	return result
}
