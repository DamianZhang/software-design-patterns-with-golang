package computationModel

type Vector []float64

func NewVector(vector []float64) *Vector {
	v := Vector(vector)
	return &v
}

func (v *Vector) Len() int {
	vector := *(v)
	return len(vector)
}

func (v *Vector) GetValue(index int) float64 {
	vector := *(v)
	return vector[index]
}
