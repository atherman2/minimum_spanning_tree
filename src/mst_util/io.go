package mst_util

import (
	"bufio"
	"fmt"
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
