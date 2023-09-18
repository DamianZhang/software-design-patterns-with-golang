package computationModel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ComputationModel struct {
	modelName string
	matrix    *Matrix
}

func newComputationModel(modelName string) *ComputationModel {
	return &ComputationModel{
		modelName: modelName,
		matrix:    NewMatrix(nil),
	}
}

func (m *ComputationModel) clone() *ComputationModel {
	return &ComputationModel{
		modelName: m.modelName,
		matrix:    m.matrix,
	}
}

// f(v) = v x m = r
func (m *ComputationModel) LinearTransformation(v *Vector) (r *Vector, err error) {
	m.matrix.rwMutex.Lock()
	defer m.matrix.rwMutex.Unlock()

	if err := m.lazyInitializationOfMatrix(); err != nil {
		return nil, fmt.Errorf("linear transformation failed to lazy initialization of matrix: %s", err)
	}

	result := make([]float64, v.Len())
	for row, valueOfRow := range m.matrix.Matrix() {
		for col := range valueOfRow {
			result[row] += (v.GetValue(col) * m.matrix.getValue(col, row))
		}
	}

	return NewVector(result), nil
}

func (m *ComputationModel) lazyInitializationOfMatrix() error {
	if !m.matrix.isEmpty() {
		return nil
	}

	var (
		fileName  = fmt.Sprintf("./%s.mat", m.modelName)
		file, err = os.Open(fileName)
	)
	if err != nil {
		return fmt.Errorf("lazy initialization of matrix failed to open file: %s", err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		matrix  = make([][]float64, 0)
	)
	for scanner.Scan() {
		var (
			rowStringSlice       = strings.Split(scanner.Text(), " ")
			rowFloat64Slice, err = ParseFloat64Slice(rowStringSlice)
		)
		if err != nil {
			return fmt.Errorf("lazy initialization of matrix failed to parse float64 slice: %s", err)
		}

		matrix = append(matrix, rowFloat64Slice)
	}

	m.matrix.setMatrix(matrix)
	fmt.Println("利用 ComputationModel, 實現 Matrix 的延遲載入")
	return nil
}
