package computationModel

import "sync"

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

	return i.virtualComputationModelProxies[modelName].clone()
}
