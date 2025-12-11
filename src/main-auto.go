package main

import (
	"fmt"
	alg "minimum_spanning_tree/algorithms"
	mst_io "minimum_spanning_tree/io"
	st "minimum_spanning_tree/structure"
	"os"
	"runtime"
	"time"
)

func main_auto() {
	fmt.Println("this is minimum_spanning_tree source code main file")

	/*
		Tratamento de flags de entrada

		go run main.go <algo-flag> [graph-id]

		<algo-flag>
		-kt : para o algoritmo de Kruskal com Tree Set Familly
		-kv : para o algoritmo de Kruskal com Index Vector
		-pm : para o algoritmo de Prim com Heap mínima
	*/
	algo_flag := os.Args[1]
	graph_id := os.Args[2]

	nodesFile := fmt.Sprintf("testes/Grafo%v/Nodes%v.csv", graph_id, graph_id)
	edgesFile := fmt.Sprintf("testes/Grafo%v/Edges%v.csv", graph_id, graph_id)

	var g st.Graph = mst_io.Adjacency_structure_graph_from_csv(nodesFile, edgesFile)

	// Medidas de memória e tempo
	var mStart, mEnd runtime.MemStats
	var total_weight st.Weight
	var duration time.Duration
	var algo string

	tamanho := g.Get_n()

	switch algo_flag {
	case "-kt":
		algo = "Kruskal Tree Set Familly"
		runtime.ReadMemStats(&mStart)
		start := time.Now()

		_, total_weight = alg.Tree_set_family_kruskal(g)

		duration = time.Since(start)
		runtime.ReadMemStats(&mEnd)

	case "-kv":
		algo = "Kruskal Vector Set Familly"
		runtime.ReadMemStats(&mStart)
		start := time.Now()

		_, total_weight = alg.Index_vector_set_family_kruskal(g)

		duration = time.Since(start)
		runtime.ReadMemStats(&mEnd)
	case "-pm":
		algo = "Prim com Heap Minima"
		runtime.ReadMemStats(&mStart)
		start := time.Now()

		_, total_weight = alg.Core_prim(g, st.Vertex(1))

		duration = time.Since(start)
		runtime.ReadMemStats(&mEnd)

	}

	// exibimos os resultados no terminal
	fmt.Printf("\nTestes para o Grafo %v - %v\n", graph_id, algo)
	fmt.Println("----------------------------------------------")
	fmt.Printf("Numero de Vertices : %v\n", tamanho)
	fmt.Println("----------------------------------------------")
	fmt.Printf("\nPeso total MST        | %.7f\n", total_weight)
	fmt.Printf("Tempo (s)             | %.7f s\n", duration.Seconds())
	fmt.Printf("Memória usada         | %.7f MB\n", float64(mEnd.Alloc-mStart.Alloc)/(1024*1024))
	fmt.Printf("Memória total alocada | %.7f MB\n", float64(mEnd.TotalAlloc-mStart.TotalAlloc)/(1024*1024))
	fmt.Printf("Pico de memória (Sys) | %.7f MB\n", float64(mEnd.Sys)/(1024*1024))
}
