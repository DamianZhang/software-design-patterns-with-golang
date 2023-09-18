package computationModel

import "sync"

type Matrix struct {
	rwMutex sync.RWMutex
	matrix  [][]float64
}

func NewMatrix(matrix [][]float64) *Matrix {
	return &Matrix{
		matrix: matrix,
	}
}

func (m *Matrix) Matrix() [][]float64 {
	// m.rwMutex.RLock()
	// defer m.rwMutex.RUnlock()

	return m.matrix
}

func (m *Matrix) getValue(row, col int) float64 {
	// m.rwMutex.RLock()
	// defer m.rwMutex.RUnlock()

	return m.matrix[row][col]
}

func (m *Matrix) isEmpty() bool {
	// m.rwMutex.RLock()
	// defer m.rwMutex.RUnlock()

	return m.matrix == nil
}

func (m *Matrix) setMatrix(matrix [][]float64) {
	// m.rwMutex.Lock()
	// defer m.rwMutex.Unlock()

	m.matrix = matrix
}
