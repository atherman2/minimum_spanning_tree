package structure

type Tree_set_family struct {
	Parent []Vertex
	Height []int
}

func (f Tree_set_family) Find_root(v Vertex) Vertex {
	for v != f.Parent[v-1] {
		v = f.Parent[v-1]
	}
	return v
}

func (f *Tree_set_family) Unite_sets(v, w Vertex) {
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

func New_tree_set_family(n int) Tree_set_family {
	result := Tree_set_family{
		Parent: make([]Vertex, n),
		Height: make([]int, n),
	}
	for i := range n {
		result.Parent[i] = Vertex(i + 1)
		result.Height[i] = 0
	}
	return result
}
