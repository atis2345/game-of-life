package main

import "flag"

func main() {
	size := flag.Int("size", 5, "size of dataset")
	iterations := flag.Int("iterations", 5, "count of iterations")
	flag.Parse()

	dataset := CreateArray(*size, *size)

	gameOfLife := GameOfLife{
		cols:    *size,
		rows:    *size,
		dataset: dataset,
	}

	for i := 0; i < *iterations; i++ {
		RenderDataSet(gameOfLife.GetDataset())
		gameOfLife.NextIteration()
	}
}
