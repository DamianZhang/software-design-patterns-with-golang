package computationModel

import "sync"

var instances *Instances = newInstances()

func getInstances() *Instances {
	return instances
}

type Instances struct {
	mutex             sync.Mutex
	computationModels map[string]*ComputationModel
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
