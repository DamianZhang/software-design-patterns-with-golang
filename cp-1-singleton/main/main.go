package main

import (
	"cp-1-singleton/computationModel/v4"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("model computer is starting...")

	var (
		vector, err  = loadVector()
		numOfClients = 10
		wg           sync.WaitGroup
	)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(numOfClients)
	for i := 0; i < numOfClients; i++ {
		go func(v *computationModel.Vector) {
			defer wg.Done()

			var (
				computationModelCreator = computationModel.GetComputationModelCreator()
				model                   = computationModelCreator.CreateModel("Reflection")
				// model  = computationModelCreator.CreateModel("Scaling")
				// model  = computationModelCreator.CreateModel("Shrinking")
				r, err = model.LinearTransformation(v)
			)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("r: %p\n", r)
			fmt.Printf("r.Len(): %v\n", r.Len())
			fmt.Printf("r.GetValue(999): %v\n", r.GetValue(999))
			fmt.Println("done")
		}(vector)
	}
	wg.Wait()
}

func loadVector() (*computationModel.Vector, error) {
	data, err := os.ReadFile("./Data.vec")
	if err != nil {
		return nil, fmt.Errorf("load vector failed to read file: %s", err)
	}

	rowStringSlice := strings.Split(string(data), " ")
	rowFloat64Slice, err := computationModel.ParseFloat64Slice(rowStringSlice)
	if err != nil {
		return nil, fmt.Errorf("load vector failed to parse float64 slice: %s", err)
	}

	return computationModel.NewVector(rowFloat64Slice), nil
}
