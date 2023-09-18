package computationModel

import (
	"fmt"
	"sync"
)

var (
	instances             *Instances = nil
	mutexForInstances     sync.Mutex
	onceMutexForInstances sync.Once
)

type Instances struct {
	mutex             sync.Mutex
	computationModels map[string]*ComputationModel
}

func getInstances() *Instances {
	mutexForInstances.Lock()
	defer mutexForInstances.Unlock()

	lazyInitializationOfInstances()

	return instances
}

func lazyInitializationOfInstances() {
	if instances == nil {
		instances = newInstances()
		fmt.Println("利用 ComputationModelCreator, 實現 Instances 的延遲載入")
	}
}

func newInstances() *Instances {
	return &Instances{
		computationModels: make(map[string]*ComputationModel),
	}
}

func (i *Instances) getComputationModel(modelName string) *ComputationModel {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	_, present := i.computationModels[modelName]
	if !present {
		i.computationModels[modelName] = newComputationModel(modelName)
	}

	return i.computationModels[modelName].clone()
}

// Code Review 後補充如何優化此處:
// 可以使用 double check instance 的技巧，解決不斷上鎖、解鎖時所產生的效能問題。
// 理想上的實作是如此沒錯，但因為 Golang 的語言特性，
// 執行 "go run --race main.go" 後會發現，依然會產生資源競爭。
func getV2Instances() *Instances {
	if instances == nil {
		mutexForInstances.Lock()
		defer mutexForInstances.Unlock()

		// 裡面還有一層 if instances == nil{} 的判斷
		// 所以此技巧才會被稱呼其為 double check instance
		lazyInitializationOfInstances()
	}

	return instances
}

// 使用 sync.Once 解決 getV2Instances() 的問題
// 可參考: https://blog.csdn.net/q5706503/article/details/105870179
func getV3Instances() *Instances {
	onceMutexForInstances.Do(func() {
		lazyInitializationOfInstances()
	})

	return instances
}
