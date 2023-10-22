package main

import (
	"math/rand"
)

type Field struct {
	cells      [][]bool
	generation int
}

func (field *Field) setRandomFieldState() {
	for i := 0; i < FIELD_LENGTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			field.cells[i][j] = rand.Intn(3) == 0
		}
	}
}

func (field *Field) calculateNextState() {
	newField := createNewField()

	for i := 0; i < FIELD_LENGTH; i++ {
		for j := 0; j < FIELD_HEIGHT; j++ {
			count := field.countNeighbors(i, j)
			if field.cells[i][j] && (count == 2 || count == 3) {
				newField.cells[i][j] = true
			} else if !field.cells[i][j] && count == 3 {
				newField.cells[i][j] = true
			} else {
				newField.cells[i][j] = false
			}
		}
	}

	field.cells = newField.cells
	field.generation++
}

func (field *Field) countNeighbors(i, j int) int {
	aliveNeighbors := 0
	if field.isAlive(i-1, j-1) {
		aliveNeighbors++
	}
	if field.isAlive(i, j-1) {
		aliveNeighbors++
	}
	if field.isAlive(i+1, j-1) {
		aliveNeighbors++
	}

	if field.isAlive(i-1, j) {
		aliveNeighbors++
	}
	if field.isAlive(i+1, j) {
		aliveNeighbors++
	}

	if field.isAlive(i-1, j+1) {
		aliveNeighbors++
	}
	if field.isAlive(i, j+1) {
		aliveNeighbors++
	}
	if field.isAlive(i+1, j+1) {
		aliveNeighbors++
	}

	return aliveNeighbors
}

func (field *Field) isAlive(x, y int) bool {
	if x < 0 || x >= FIELD_LENGTH || y < 0 || y >= FIELD_HEIGHT {
		return false
	}
	return field.cells[x][y]
}

func createNewField() *Field {
	field := &Field{}

	cells := make([][]bool, FIELD_LENGTH)
	for i := 0; i < FIELD_LENGTH; i++ {
		cells[i] = make([]bool, FIELD_HEIGHT)
	}

	field.cells = cells
	field.generation = 0

	return field
}
