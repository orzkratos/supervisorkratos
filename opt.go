package supervisorkratos

type Opt[T any] struct {
	Value T
	isSet bool
}

func NewOpt[T any](v T) *Opt[T] {
	return &Opt[T]{Value: v, isSet: false}
}

func (sv *Opt[T]) Get() T {
	return sv.Value
}

func (sv *Opt[T]) Set(v T) {
	sv.Value = v
	sv.isSet = true
}

func (sv *Opt[T]) IsSet() bool {
	return sv.isSet
}
