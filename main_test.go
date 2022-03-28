package main

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	iterations := 4
	cols := 5
	rows := 5

	dataset := DataSet{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	steps := []DataSet{
		{
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 1, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{0, 0, 0, 0, 1},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0},
		},
	}

	gameOfLife := NewGameOfLife(cols, rows, dataset)

	RenderDataSet(gameOfLife.GetDataset())
	require.Equal(t, dataset, gameOfLife.GetDataset())

	for i := 0; i < iterations; i++ {
		RenderDataSet(gameOfLife.GetDataset())
		gameOfLife.NextIteration()
		require.Equal(t, steps[i], gameOfLife.GetDataset())
	}
}
