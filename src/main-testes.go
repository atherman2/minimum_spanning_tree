package main

import (
	"fmt"
	alg "minimum_spanning_tree/algorithms"
	mst_io "minimum_spanning_tree/io"
	st "minimum_spanning_tree/structure"
	"runtime"
	"time"
)

func runTest(nodesFile, edgesFile string) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Rodando teste com: %s e %s\n", nodesFile, edgesFile)

	// carrega Grafo
	var g st.Graph = mst_io.Adjacency_structure_graph_from_csv(nodesFile, edgesFile)

	// mede memória antes
	var mStart, mEnd runtime.MemStats
	runtime.ReadMemStats(&mStart)

	// mede tempo
	start := time.Now()
	_, totalWeight := alg.Core_prim(g, st.Vertex(1)) // ou alg.Tree_set_family_kruskal(g)
	duration := time.Since(start)

	// mede memória depois
	runtime.ReadMemStats(&mEnd)

	// imprime resultados
	fmt.Printf("Peso total MST: %.2f\n", totalWeight)
	fmt.Printf("Tempo (s): %.6f\n", duration.Seconds())
	fmt.Printf("Memória usada: %.2f MB\n", float64(mEnd.Alloc-mStart.Alloc)/(1024*1024))
	fmt.Printf("Memória total alocada: %.2f MB\n", float64(mEnd.TotalAlloc-mStart.TotalAlloc)/(1024*1024))
	fmt.Printf("Pico de memória (Sys): %.2f MB\n", float64(mEnd.Sys)/(1024*1024))
}

func main_deprec() {
	fmt.Println("this is minimum_spanning_tree source code main file")

	// lista de Grafos para testar
	graphs := [][2]string{
		{"testes/Grafo1/Nodes1.csv", "testes/Grafo1/Edges1.csv"},
		{"testes/Grafo2/Nodes2.csv", "testes/Grafo2/Edges2.csv"},
		{"testes/Grafo3/Nodes3.csv", "testes/Grafo3/Edges3.csv"},
		{"testes/Grafo4/Nodes4.csv", "testes/Grafo4/Edges4.csv"},
		{"testes/Grafo5/Nodes5.csv", "testes/Grafo5/Edges5.csv"},
		{"testes/Grafo6/Nodes6.csv", "testes/Grafo6/Edges6.csv"},
	}

	// roda todos os testes
	for _, g := range graphs {
		runTest(g[0], g[1])
	}
}
