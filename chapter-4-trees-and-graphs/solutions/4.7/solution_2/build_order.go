package build_order

import (
	"errors"
)

type Project struct {
	Name         string
	Dependencies int
	Children     []*Project
	Map          map[string]*Project
	State        int
}

const (
	Unvisited = iota
	Visiting
	Visited
)

func (p *Project) AddNeighbor(node *Project) {
	if _, ok := p.Map[node.Name]; !ok {
		p.Children = append(p.Children, node)
		p.Map[node.Name] = node
		node.Dependencies++
	}
}

type Graph struct {
	Nodes []*Project
	Map   map[string]*Project
}

func (g *Graph) GetOrCreateNode(name string) *Project {
	if _, ok := g.Map[name]; !ok {
		node := &Project{
			Name: name,
			Map:  make(map[string]*Project),
		}
		g.Nodes = append(g.Nodes, node)
		g.Map[name] = node
	}
	return g.Map[name]
}

func (g *Graph) AddEdge(startName string, endName string) {
	start := g.GetOrCreateNode(startName)
	end := g.GetOrCreateNode(endName)
	start.AddNeighbor(end)
}

func BuildGraph(projects []string, dependencies [][]string) *Graph {
	graph := &Graph{
		Map: make(map[string]*Project),
	}
	for _, project := range projects {
		graph.GetOrCreateNode(project)
	}
	for _, dep := range dependencies {
		first := dep[0]
		second := dep[1]
		graph.AddEdge(first, second)
	}
	return graph
}

func DFS(project *Project, stack *[]string) error {
	if project.State == Visiting {
		return errors.New("no valid build order found: cycle detected")
	}
	if project.State == Unvisited {
		project.State = Visiting
		for _, child := range project.Children {
			if err := DFS(child, stack); err != nil {
				return err
			}
		}
		project.State = Visited
		*stack = append(*stack, project.Name)
	}
	return nil
}

func OrderProjectsDFS(projects []*Project) ([]string, error) {
	stack := []string{}
	for _, project := range projects {
		if project.State == Unvisited {
			if err := DFS(project, &stack); err != nil {
				return nil, err
			}
		}
	}
	Reverse(stack)
	return stack, nil
}

func Reverse(slice []string) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

func FindBuildOrder(projects []string, dependencies [][]string) ([]string, error) {
	graph := BuildGraph(projects, dependencies)
	return OrderProjectsDFS(graph.Nodes)
}
