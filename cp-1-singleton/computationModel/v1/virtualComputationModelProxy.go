package computationModel

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type VirtualComputationModelProxy struct {
	modelName string

	// rwMutex          sync.RWMutex
	mutex            sync.Mutex
	computationModel *ComputationModel
}

func newVirtualComputationModelProxy(modelName string) *VirtualComputationModelProxy {
	return &VirtualComputationModelProxy{
		modelName:        modelName,
		computationModel: nil,
	}
}

func (m *VirtualComputationModelProxy) LinearTransformation(v *Vector) (r *Vector, err error) {
	// 支援非同步的方法一: 直接上全鎖
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// // 支援非同步的方法二步驟一: 執行 "寫" 操作前使用讀寫鎖上全鎖，寫完後解鎖
	// m.rwMutex.Lock()
	if m.computationModel == nil {
		// // 使用方法二時，鎖不能再判斷式後才上，因為可能會有
		// // 一個 goroutine 在執行 if m.computationModel == nil
		// // 而另一個 goroutine 同時在執行 m.lazyInitialization() 裡的
		// // m.computationModel = newComputationModel(m.modelName, NewMatrix(matrix))
		// // 進而造成 race condition
		// m.rwMutex.Lock()
		if err := m.lazyInitialization(); err != nil {
			// m.rwMutex.Unlock()
			return nil, fmt.Errorf("linear transformation failed to lazy initialization: %s", err)
		}
		// m.rwMutex.Unlock()
	}
	// m.rwMutex.Unlock()

	// // 支援非同步的方法二步驟二: 執行 "讀" 操作前使用讀寫鎖上讀鎖，讀完後解鎖
	// m.rwMutex.RLock()
	// defer m.rwMutex.RUnlock()

	return m.computationModel.LinearTransformation(v)
}

func (m *VirtualComputationModelProxy) lazyInitialization() error {
	var (
		fileName  = fmt.Sprintf("./%s.mat", m.modelName)
		file, err = os.Open(fileName)
	)
	if err != nil {
		return fmt.Errorf("lazy initialization failed to open file: %s", err)
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
			return fmt.Errorf("lazy initialization failed to parse float64 slice: %s", err)
		}

		matrix = append(matrix, rowFloat64Slice)
	}

	m.computationModel = newComputationModel(m.modelName, NewMatrix(matrix))
	return nil
}
