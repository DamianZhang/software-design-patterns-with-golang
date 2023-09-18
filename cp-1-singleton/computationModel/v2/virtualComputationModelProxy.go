package computationModel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VirtualComputationModelProxy struct {
	modelName        string
	computationModel *ComputationModel
}

func newVirtualComputationModelProxy(modelName string) *VirtualComputationModelProxy {
	return &VirtualComputationModelProxy{
		modelName:        modelName,
		computationModel: newComputationModel(modelName),
	}
}

func (m *VirtualComputationModelProxy) clone() *VirtualComputationModelProxy {
	return &VirtualComputationModelProxy{
		modelName:        m.modelName,
		computationModel: m.computationModel,
	}
}

func (m *VirtualComputationModelProxy) LinearTransformation(v *Vector) (r *Vector, err error) {
	m.computationModel.mutex.Lock()
	defer m.computationModel.mutex.Unlock()

	if err := m.lazyInitializationOfMatrix(); err != nil {
		return nil, fmt.Errorf("linear transformation failed to lazy initialization of matrix: %s", err)
	}

	return m.computationModel.LinearTransformation(v)
}

func (m *VirtualComputationModelProxy) lazyInitializationOfMatrix() error {
	if !m.computationModel.hasEmptyMatrix() {
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

	m.computationModel.setMatrix(NewMatrix(matrix))
	return nil
}
