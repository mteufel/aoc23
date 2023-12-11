package day8

type Graph struct {
	vertices map[Vertex][]Vertex
}

type Vertex struct {
	Label string
}

func (g *Graph) AddVertex(label string) {
	vertex := Vertex{Label: label}
	g.vertices[vertex] = make([]Vertex, 0)
}

func (g *Graph) AddEdge(label1 string, label2 string) {
	vertex1 := Vertex{Label: label1}
	vertex2 := Vertex{Label: label2}
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

func LoadGraph() Graph {

	graph := Graph{vertices: make(map[Vertex][]Vertex)}

	graph.AddVertex("AAA")
	graph.AddVertex("BBB")
	graph.AddVertex("CCC")
	graph.AddVertex("DDD")
	graph.AddVertex("EEE")
	graph.AddVertex("GGG")
	graph.AddVertex("ZZZ")

	graph.AddEdge("AAA", "BBB") // Left
	graph.AddEdge("AAA", "CCC") // Right

	graph.AddEdge("BBB", "DDD") // Left
	graph.AddEdge("BBB", "EEE") // Right

	graph.AddEdge("CCC", "ZZZ") // Left
	graph.AddEdge("CCC", "GGG") // Right

	graph.AddEdge("DDD", "DDD") // Left
	graph.AddEdge("DDD", "DDD") // Right

	graph.AddEdge("EEE", "EEE") // Left
	graph.AddEdge("EEE", "EEE") // Right

	graph.AddEdge("GGG", "GGG") // Left
	graph.AddEdge("GGG", "GGG") // Right

	graph.AddEdge("ZZZ", "ZZZ") // Right
	graph.AddEdge("ZZZ", "ZZZ") // Right

	return graph
}
