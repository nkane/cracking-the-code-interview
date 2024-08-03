package build_order

import (
	"errors"
)

/*
	4.7: Build Order: You are given a list of projects and a list of dependencies (which is a list of pairs of
	projects, where the second project is dependent on the first project). All of the project's dependencies
	must be built before the project is. Find a build order that will allow the project to be built. If there
	is no valid build order, return an error.

	Example:
	Input:
		projects: a, b, c, d, e, f
		dependencies: (a, d), (f, b), (b, d), (f, a), (d, c)
	Output: f, e, a, b, d, c
*/

type Project struct {
	Children     []*Project
	Map          map[string]*Project
	Name         string
	Dependencies int
}

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

func AddNonDependent(order []*Project, projects []*Project, offset int) int {
	for _, project := range projects {
		if project.Dependencies == 0 {
			order[offset] = project
			offset++
		}
	}
	return offset
}

// OrderProjects returns a list of the projects in the correct build order
func OrderProjects(projects []*Project) []*Project {
	order := make([]*Project, len(projects))
	endOfList := AddNonDependent(order, projects, 0)
	toBeProcessed := 0
	for toBeProcessed < endOfList {
		current := order[toBeProcessed]
		if current == nil {
			return nil
		}
		children := current.Children
		for _, child := range children {
			child.Dependencies--
		}
		endOfList = AddNonDependent(order, children, endOfList)
		toBeProcessed++
	}
	if toBeProcessed < len(order) {
		return nil
	}
	return order
}

// BuildGraph builds the graph, adding the edge (a, b) if b is dependent on a
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

// FindBuildOrder finds a correct build order
func FindBuildOrder(projects []string, dependencies [][]string) ([]string, error) {
	graph := BuildGraph(projects, dependencies)
	projectOrder := OrderProjects(graph.Nodes)
	if projectOrder == nil {
		return nil, errors.New("no valid build order found")
	}

	result := make([]string, len(projectOrder))
	for i, project := range projectOrder {
		result[i] = project.Name
	}
	return result, nil
}
