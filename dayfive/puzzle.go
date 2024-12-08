package dayfive

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	nodes    map[int][]int
	vertices []int
}

func (g *Graph) addEdge(u, v int) {
	g.nodes[u] = append(g.nodes[u], v)
}

func (g *Graph) topologicalSortUtil(v int, visited map[int]bool, stack *[]int, allowed map[int]bool) {
	visited[v] = true

	for _, u := range g.nodes[v] {
		if allowed[u] && !visited[u] {
			g.topologicalSortUtil(u, visited, stack, allowed)
		}
	}

	*stack = append([]int{v}, *stack...)
}

func (g *Graph) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)

	allowed := make(map[int]bool)
	for _, v := range g.vertices {
		allowed[v] = true
	}

	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack, allowed)
		}
	}

	return stack
}

func (g *Graph) middleVertice() int {
	return g.vertices[len(g.vertices)/2]
}

func Puzzle(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	parsingEdges := true
	g := &Graph{
		nodes:    make(map[int][]int),
		vertices: []int{},
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingEdges = false
			continue
		}

		if parsingEdges {
			parts := strings.Split(scanner.Text(), "|")
			g.addEdge(stringToInt(parts[0]), stringToInt(parts[1]))
		}

		if !parsingEdges {
			g.vertices = []int{}
			parsedVertices := strings.Split(line, ",")
			for _, n := range parsedVertices {
				g.vertices = append(g.vertices, stringToInt(n))
			}

			sorted := g.topologicalSort()
			if slicesEqual(g.vertices, sorted) {
				sum += g.middleVertice()
			}
		}
	}

	return sum, nil
}

func stringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
