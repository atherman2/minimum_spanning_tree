package mst_io

import (
	"bufio"
	"fmt"
	mu "minimum_spanning_tree/mst_util"
	st "minimum_spanning_tree/structure"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	X, Y float64
}

func Adjacency_structure_graph_from_csv(nodes_file_path, edges_file_path string) st.Graph {
	nodes_file, err := os.Open(nodes_file_path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer nodes_file.Close()
	nodes_scanner := bufio.NewScanner(nodes_file)
	nodes := Read_nodes(nodes_scanner)

	edges_file, err := os.Open(edges_file_path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer edges_file.Close()
	edges_scanner := bufio.NewScanner(edges_file)
	edges := Read_edges(edges_scanner, nodes)
	return st.Adjacency_structure_graph_from_edges(len(nodes), edges)
}

func Read_nodes(scanner *bufio.Scanner) []Vertex {
	result := []Vertex{}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		str_x, str_y := line_split[1], line_split[2]
		x, err := strconv.ParseFloat(str_x, 64)
		if err != nil {
			fmt.Println("Error parsing nodes file: ", err)
		}
		y, err := strconv.ParseFloat(str_y, 64)
		if err != nil {
			fmt.Println("Error parsing nodes file: ", err)
		}
		result = append(result, Vertex{X: x, Y: y})
	}
	return result
}

func Read_edges(scanner *bufio.Scanner, nodes []Vertex) []st.Edge {
	result := []st.Edge{}
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		str_v, str_w := line_split[0], line_split[1]
		v, err := strconv.Atoi(str_v)
		if err != nil {
			fmt.Println("Error parsing edges file: ", err)
		}
		w, err := strconv.Atoi(str_w)
		if err != nil {
			fmt.Println("Error parsing edges file: ", err)
		}
		xv, yv, xw, yw := nodes[v-1].X, nodes[v-1].Y, nodes[w-1].X, nodes[w-1].Y
		e := st.Edge{V: st.Vertex(v), W: st.Vertex(w), Weight: st.Weight(mu.Distance_2D_float64(xv, yv, xw, yw))}
		result = append(result, e)
	}
	return result
}
