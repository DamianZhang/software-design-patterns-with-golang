package main

import "sp-2-adapter/advanced"

func main() {
	// client := basic.NewClient(basic.NewSuperRelationshipAnalyzerAdapter())
	// client.Start()

	client := advanced.NewClient(advanced.NewGraphAdapter())
	client.Start()
}
