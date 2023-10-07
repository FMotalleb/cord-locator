package iterable

type Mapper[P any, R any] func(from P) R

func Map[P, R any](from []P, mapper Mapper[P, R]) []R {
	result := make([]R, len(from))
	for i := range from {
		result[i] = mapper(from[i])
	}
	return result
}
