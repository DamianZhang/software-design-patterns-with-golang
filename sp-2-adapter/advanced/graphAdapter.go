package advanced

import (
	"fmt"
	"strings"

	"github.com/dominikbraun/graph"
)

type GraphAdapter struct {
	graph graph.Graph[string, string]
}

func NewGraphAdapter() *GraphAdapter {
	return &GraphAdapter{
		graph: graph.New(graph.StringHash),
	}
}

func (g *GraphAdapter) Parse(script string) RelationshipGraph {
	members := strings.Split(script, "\n")

	for _, member := range members {
		var (
			infoOfMember           = strings.Split(member, ": ")
			nameOfMember           = infoOfMember[0]
			namesOfFriendsOfMember = strings.Split(infoOfMember[1], " ")
		)

		g.graph.AddVertex(nameOfMember)

		for _, nameOfFriendOfMember := range namesOfFriendsOfMember {
			g.graph.AddVertex(nameOfFriendOfMember)
			g.graph.AddEdge(nameOfMember, nameOfFriendOfMember)
		}
	}

	return g
}

func (g *GraphAdapter) HasConnection(name1, name2 string) bool {
	path, err := graph.ShortestPath(g.graph, name1, name2)
	if err == graph.ErrTargetNotReachable {
		return false
	}

	fmt.Printf("path: %v\n", path)
	return true
}
