package main

import "fmt"

type graph struct {
	vertices int
	adjMat   [][]int
	visited  []int
}

func constructor(vertices int) *graph {
	g := &graph{}
	g.vertices = vertices
	g.visited = make([]int, g.vertices)
	for i := 0; i < g.vertices; i++ {
		g.adjMat = append(g.adjMat, make([]int, g.vertices))
	}
	//fmt.Println(g.adjMat)
	return g
}

func (g *graph) insert_edge(u, v, x int) {
	if x == 0 {
		x = 1
	}
	g.adjMat[u][v] = x
}

func (g *graph) remove_edge(u, v int) {
	g.adjMat[u][v] = 0
}

func (g *graph) exist_edge(u, v int) bool {
	return g.adjMat[u][v] != 0
}

func (g *graph) vertix_count() int {
	return g.vertices
}

func (g *graph) edge_count() int {
	count := 0
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if g.adjMat[i][j] != 0 {
				count++
			}
		}
	}
	return count
}

func (g *graph) vertice() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func (g *graph) edges() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if g.adjMat[i][j] != 0 {
				fmt.Printf("%d--%d\n", i, j)
			}
		}
	}
}

func (g *graph) outdegree(v int) int {
	count := 0
	for i := 0; i < g.vertices; i++ {
		if g.adjMat[v][i] != 0 {
			count++
		}
	}
	return count
}

func (g *graph) indegree(v int) int {
	count := 0
	for i := 0; i < g.vertices; i++ {
		if g.adjMat[i][v] != 0 {
			count++
		}
	}
	return count
}

func (g *graph) displayadjMat() {
	fmt.Println(g.adjMat)
}

func (g *graph) BFS(s int) {
	i := s
	q := []int{}
	visited := make([]int, g.vertices)
	fmt.Print(i, " ")
	visited[i] = 1
	q = append(q, i)
	for len(q) != 0 {
		i = q[0]
		q = q[1:]
		for j := 0; j < g.vertices; j++ {
			if g.adjMat[i][j] != 0 && visited[j] == 0 {
				fmt.Print(j, " ")
				visited[j] = 1
				q = append(q, j)
			}
		}
	}
	fmt.Println()
}

func (g *graph) DFS(s int) {

	if g.visited[s] == 0 {
		fmt.Print(s, " ")
		g.visited[s] = 1
		for j := 0; j < g.vertices; j++ {
			if g.adjMat[s][j] != 0 && g.visited[j] == 0 {
				g.DFS(j)
			}
		}
	}
}

func main() {
	g := constructor(7)
	g.displayadjMat()
	fmt.Printf("%d\n", g.vertix_count())
	fmt.Printf("%d\n", g.edge_count())
	g.insert_edge(0, 1, 0)
	g.insert_edge(0, 5, 0)
	g.insert_edge(0, 6, 0)
	g.insert_edge(1, 0, 0)
	g.insert_edge(1, 6, 0)
	g.insert_edge(1, 5, 0)
	g.insert_edge(1, 2, 0)
	g.insert_edge(2, 6, 0)
	g.insert_edge(2, 3, 0)
	g.insert_edge(3, 4, 0)
	g.insert_edge(4, 5, 0)
	g.insert_edge(5, 2, 0)
	g.insert_edge(5, 3, 0)
	g.insert_edge(6, 3, 0)
	g.displayadjMat()
	fmt.Printf("%d\n", g.vertix_count())
	fmt.Printf("%d\n", g.edge_count())
	g.vertice()
	g.edges()
	fmt.Printf("Indegree of %d : %d\n", 2, g.indegree(2))
	fmt.Printf("%t\n", g.exist_edge(1, 0))
	g.remove_edge(1, 0)
	fmt.Printf("%t\n", g.exist_edge(1, 0))
	//fmt.Println()
	g.BFS(0)
	//fmt.Println()
	g.DFS(0)

}
