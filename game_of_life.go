package main

import (
	"github.com/olekukonko/tablewriter"
	"math/rand"
	"os"
)

type DataSet [][]int

type GameOfLife struct {
	cols, rows int
	dataset    DataSet
}

func NewGameOfLife(cols, rows int, dataset DataSet) GameOfLife {
	return GameOfLife{
		cols:    cols,
		rows:    rows,
		dataset: dataset,
	}
}

func (g *GameOfLife) sumNeighbors(dataset DataSet, x, y int) int {
	sum := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			col := (x + i + g.cols) % g.cols
			row := (y + j + g.rows) % g.rows

			sum += dataset[col][row]
		}
	}

	sum -= dataset[x][y]

	return sum
}

func (g *GameOfLife) NextIteration() {
	var next DataSet
	next = CreateArray(g.cols, g.rows)

	for i := 0; i < g.cols; i++ {
		for j := 0; j < g.rows; j++ {
			state := g.dataset[i][j]
			neighbors := g.sumNeighbors(g.dataset, i, j)

			// Dies by underpopulation
			if state == 1 && neighbors < 2 {
				next[i][j] = 0
				// Alive until next generation
			} else if state == 1 && (neighbors == 2 || neighbors == 3) {
				next[i][j] = 1
				// Dies by overcrowding
			} else if state == 1 && neighbors > 3 {
				next[i][j] = 0
				// Alive by reproduction
			} else if state == 0 && neighbors == 3 {
				next[i][j] = 1
			} else {
				next[i][j] = state
			}
		}
	}

	g.dataset = next
}

func (g *GameOfLife) GetDataset() DataSet {
	return g.dataset
}

func CreateArray(cols, rows int) DataSet {
	var result DataSet
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			row[j] = rand.Intn(2)
		}
		result = append(result, row)
	}
	return result
}

func RenderDataSet(dataset DataSet) {
	table := tablewriter.NewWriter(os.Stdout)
	var data [][]string

	for _, row := range dataset {
		strRow := make([]string, len(row))
		for k, col := range row {
			if col == 1 {
				strRow[k] = "X"
			} else {
				strRow[k] = "-"
			}
		}
		data = append(data, strRow)
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
