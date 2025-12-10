package structure

type TreeSetFamily struct {
	Parent []Vertex
	Height []Vertex
}

func (f TreeSetFamily) Find_root(v Vertex) Vertex {
	for v != f.Parent[v-1] {
		v = f.Parent[v-1]
	}
	return v
}

func (f *TreeSetFamily) Unite_sets(v, w Vertex) {
	root_v := f.Find_root(v)
	root_w := f.Find_root(w)
	if f.Height[root_v-1] > f.Height[root_w-1] {
		f.Parent[root_w-1] = root_v
	} else if f.Height[root_w-1] > f.Height[root_v-1] {
		f.Parent[root_v-1] = root_w
	} else {
		f.Parent[root_w-1] = root_v
		f.Height[root_v-1]++
	}
}
