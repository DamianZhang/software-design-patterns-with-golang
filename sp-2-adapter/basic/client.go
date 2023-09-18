package basic

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
	c.analyzer.Parse(script)

	mutualFriendsOfAAndB := c.analyzer.GetMutualFriends("A", "B")
	fmt.Printf("mutualFriendsOfAAndB: %v\n", mutualFriendsOfAAndB)
}
