package main

import (
	"fmt"
	mu "minimum_spanning_tree/mst_util"
	st "minimum_spanning_tree/structure"
)

func main() {
	fmt.Println("this is minimum_spanning_tree source code main file")
	var v_list st.Vertex_list = &st.Vertex_linked_list{V: st.Vertex(1), Next: nil, Weight: st.Weight(1)}
	fmt.Println(mu.Search(v_list, st.Vertex(1)))
	fmt.Println(mu.Search(v_list, st.Vertex(2)))
}
