package build_order

import "graph"

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
	Name string
}

type ProjectGraph struct {
	graph.Graph[string, Project]
}

func (pg *ProjectGraph) GetOrCreateNode(name string) *graph.Node[Project] {
	if _, ok := pg.Links[name]; !ok {
		p := &graph.Node[Project]{
			Data: Project{
				Name: name,
			},
		}
		pg.Nodes = append(pg.Nodes, p)
		pg.Links[name] = p
	}
	return pg.Links[name]
}

func (pg *ProjectGraph) AddEdge(startName string, endName string) {
	start := pg.GetOrCreateNode(startName)
	end := pg.GetOrCreateNode(endName)
	start.AddNeighbor(end)
}

func BuildOrder(projects []string, dependencies [][]string) []Project {
	// TODO(nick):
	return nil
}
