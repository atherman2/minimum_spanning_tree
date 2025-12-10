package structure

type IndexVectorSetFamily []Vertex

func (f IndexVectorSetFamily) Find_root(v Vertex) Vertex {
	for v != f[v-1] {
		v = f[v-1]
	}
	return v
}

func (f *IndexVectorSetFamily) Unite_sets(v, w Vertex) {
	root_v := f.Find_root(v)
	root_w := f.Find_root(w)
	(*f)[root_w-1] = root_v
}
