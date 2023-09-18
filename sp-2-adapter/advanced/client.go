package advanced

import (
	"fmt"
	"os"
)

type Client struct {
	analyzer RelationshipAnalyzer
}

func NewClient(analyzer RelationshipAnalyzer) *Client {
	return &Client{
		analyzer: analyzer,
	}
}

func (c *Client) Start() {
	fmt.Println("relationship analyzer is starting...")

	bytesOfScript, err := os.ReadFile("./script.txt")
	if err != nil {
		panic(err)
	}

	script := string(bytesOfScript)
	graph := c.analyzer.Parse(script)

	AAndBHasConnection := graph.HasConnection("A", "B")
	fmt.Printf("AAndBHasConnection: %v\n", AAndBHasConnection)
}
