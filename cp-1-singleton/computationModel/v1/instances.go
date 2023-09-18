package computationModel

import "sync"

// 需求:
// 1. 盡可能地節省記憶體資源的使用。
// 2. 盡量讓套件中每一個物件延遲載入。
// 3. 支援非同步。

// 解釋 Instances, ComputationModel, VirtualComputationModelProxy:
// 1. 為了避免多位 Clients 透過某種無節制的操作來過度消耗記憶體資源，
//	例如: 無節制的產生新的 Instances, ComputationModel, VirtualComputationModelProxy，
//	所以 Instances 使用單體模式，
//	而 ComputationModel 則透過關閉建構子並將建構子交給 VirtualComputationModelProxy 管理，
//	然後 VirtualComputationModelProxy 也關閉建構子並將建構子交給 Instances 管理的方式，
//	讓 ComputationModel, VirtualComputationModelProxy 達到單體模式的效果，
//	進而避免過度消耗記憶體資源的情況發生。
// 2. Instances 沒有支援延遲載入，ComputationModel 支援延遲載入。
//	原因: 單體模式 Instances 物件，本身不包含大資料，所以此處實現延遲載入的效益不大。
//	且想要延遲載入的同時並支援非同步，目前只知道 Lock 的解法，
//	為了拿單體物件必須不斷上鎖、解鎖，造成效能大幅下降，
//	權衡利弊後，不支援延遲載入是更好的選擇。
//	而 ComputationModel 物件包含大資料(Matrix)，所以此處實現延遲載入的效益較大。
//	因此透過 VirtualComputationModelProxy 支援延遲載入。
// 3. 支援非同步。
//	原因: 在使用單體模式的情況下，無論如何都只有一個實體，
//	實體下的各種操作沒有資源競爭的問題，代表著該實體滿足支援非同步的條件。

var instances *Instances = newInstances()

func getInstances() *Instances {
	return instances
}

type Instances struct {
	mutex                          sync.Mutex
	virtualComputationModelProxies map[string]*VirtualComputationModelProxy
}

func newInstances() *Instances {
	return &Instances{
		virtualComputationModelProxies: make(map[string]*VirtualComputationModelProxy),
	}
}

func (i *Instances) getVirtualComputationModelProxy(modelName string) *VirtualComputationModelProxy {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	_, present := i.virtualComputationModelProxies[modelName]
	if !present {
		i.virtualComputationModelProxies[modelName] = newVirtualComputationModelProxy(modelName)
	}

	return i.virtualComputationModelProxies[modelName]
}
