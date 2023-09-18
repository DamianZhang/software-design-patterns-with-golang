package computationModel

type Matrix [][]float64

func NewMatrix(matrix [][]float64) *Matrix {
	m := Matrix(matrix)
	return &m
}

func (m *Matrix) getValue(row, col int) float64 {
	matrix := *(m)
	return matrix[row][col]
}
