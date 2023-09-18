package computationModel

import (
	"fmt"
	"sync"
)

type ComputationModel struct {
	modelName string

	mutex  sync.Mutex
	matrix *Matrix
}

func newComputationModel(modelName string) *ComputationModel {
	return &ComputationModel{
		modelName: modelName,
		matrix:    nil,
	}
}

// f(v) = v x m = r
func (m *ComputationModel) LinearTransformation(v *Vector) (r *Vector, err error) {
	if m.hasEmptyMatrix() {
		return nil, fmt.Errorf("ComputationModel.matrix should NOT be nil")
	}

	result := make([]float64, v.Len())
	for row, valueOfRow := range *(m.matrix) {
		for col := range valueOfRow {
			result[row] += (v.GetValue(col) * m.matrix.getValue(col, row))
		}
	}

	return NewVector(result), nil
}

func (m *ComputationModel) hasEmptyMatrix() bool {
	return m.matrix == nil
}

func (m *ComputationModel) setMatrix(matrix *Matrix) {
	m.matrix = matrix
}
