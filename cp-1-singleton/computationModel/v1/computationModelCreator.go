package computationModel

// v1 特徵:
// 1. 最少實體化。
//	ComputationModelCreator: 只有 1 個實體。
//	Instances: 只有 1 個實體。
//	VirtualComputationModelProxy: 只有 len(modelNames) 個實體。
//	ComputationModel: 只有 len(modelNames) 個實體。
//	Matrix: 只有 len(modelNames) 個實體。
// 2. 支援 ComputationModel 的延遲載入。
// 3. 支援非同步。

// 需求:
// 1. 盡可能地節省記憶體資源的使用。
// 2. 盡量讓套件中每一個物件延遲載入。
// 3. 支援非同步。

// 解釋 ComputationModelCreator:
// 1. 為了避免多位 Clients 透過某種無節制的操作來過度消耗記憶體資源，
//	例如: 無節制的產生新的 ComputationModelCreator，
//	所以使用單體模式，來避免此情況發生。
// 2. 沒有支援延遲載入。
//	原因: 單體模式物件，本身不包含大資料，所以此處實現延遲載入的效益不大。
//	且想要延遲載入的同時並支援非同步，目前只知道 Lock 的解法，
//	為了拿單體物件必須不斷上鎖、解鎖，造成效能大幅下降，
//	權衡利弊後，不支援延遲載入是更好的選擇。
// 3. 支援非同步。
//	原因: 在使用單體模式的情況下，無論如何都只有一個 ComputationModelCreator 實體，
//	ComputationModelCreator 下的各種操作沒有資源競爭的問題，
//	代表著 ComputationModelCreator 滿足支援非同步的條件。

var computationModelCreator *ComputationModelCreator = newComputationModelCreator()

func GetComputationModelCreator() Models {
	return computationModelCreator
}

type ComputationModelCreator struct{}

func newComputationModelCreator() *ComputationModelCreator {
	return &ComputationModelCreator{}
}

func (c *ComputationModelCreator) CreateModel(modelName string) Model {
	return getInstances().getVirtualComputationModelProxy(modelName)
}
