package computationModel

type Model interface {
	LinearTransformation(v *Vector) (r *Vector, err error)
}
