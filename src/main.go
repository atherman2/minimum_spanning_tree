package main

import (
	"fmt"
	mst_io "minimum_spanning_tree/io"
	mu "minimum_spanning_tree/mst_util"
	st "minimum_spanning_tree/structure"
	"os"
)

func main() {
	fmt.Println("this is minimum_spanning_tree source code main file")
	var v_list st.Vertex_list = &st.Vertex_linked_list{V: st.Vertex(1), Next: nil, Weight: st.Weight(1)}
	fmt.Println(mu.Search(v_list, st.Vertex(1)))
	fmt.Println(mu.Search(v_list, st.Vertex(2)))
	var g st.Graph = mst_io.Adjacency_structure_graph_from_csv(os.Args[1], os.Args[2])
	fmt.Println(g.Get_edge(st.Vertex(1), st.Vertex(2)).Data())
}
