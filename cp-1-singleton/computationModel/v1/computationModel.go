package computationModel

type ComputationModel struct {
	modelName string
	matrix    *Matrix
}

func newComputationModel(modelName string, matrix *Matrix) *ComputationModel {
	return &ComputationModel{
		modelName: modelName,
		matrix:    matrix,
	}
}

// f(v) = v x m = r
func (m *ComputationModel) LinearTransformation(v *Vector) (r *Vector, err error) {
	result := make([]float64, v.Len())
	for row, valueOfRow := range *(m.matrix) {
		for col := range valueOfRow {
			result[row] += (v.GetValue(col) * m.matrix.getValue(col, row))
		}
	}

	return NewVector(result), nil
}
