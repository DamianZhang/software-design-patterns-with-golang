package computationModel

// v3 特徵:
// 1. 相對 v1 有著更多的實體化。
//	ComputationModelCreator: 只有 1 個實體。
//	Instances: 只有 1 個實體。
//	ComputationModel: 呼叫幾次 getInstance() 就有幾個實體。
//	Matrix: 只有 len(modelNames) 個實體。
// 2. 支援 Instances, Matrix 的延遲載入。
// 3. 支援非同步。
// 4. 相對 v3，v4 有著更全面的延遲載入，
//	但是 v4 效能更差，因為 getInstances() 需要一直上鎖、解鎖。

var computationModelCreator *ComputationModelCreator = newComputationModelCreator()

func GetComputationModelCreator() Models {
	return computationModelCreator
}

type ComputationModelCreator struct{}

func newComputationModelCreator() *ComputationModelCreator {
	return &ComputationModelCreator{}
}

func (c *ComputationModelCreator) CreateModel(modelName string) Model {
	return getInstances().getComputationModel(modelName)
}
