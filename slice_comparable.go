package stream

import "golang.org/x/exp/slices"

// SliceComparableStream Generics constraints based on comparable
type SliceComparableStream[E comparable] struct {
	SliceStream[E]
}

// NewSliceByComparable  new stream instance, generics constraints based on comparable
func NewSliceByComparable[E comparable](source []E) SliceComparableStream[E] {
	return SliceComparableStream[E]{SliceStream: NewSlice(source)}
}

// Distinct Returns a stream consisting of the distinct elements of this stream.
// Remove duplicate according to map comparable.
func (stream SliceComparableStream[E]) Distinct() SliceComparableStream[E] {
	stream.evaluation()
	if stream.source == nil && len(stream.source) < 2 {
		return stream
	}

	newSlice := make([]E, 0)
	distinct := map[E]struct{}{}
	for _, v := range stream.source {
		if _, ok := distinct[v]; ok {
			continue
		}
		distinct[v] = struct{}{}
		newSlice = append(newSlice, v)
	}
	stream.source = newSlice
	return stream
}

// Equal Returns whether the source in the stream is equal to the destination source.
// Equal according to the slices.Equal.
func (stream SliceComparableStream[E]) Equal(dest []E) bool {
	stream.evaluation()
	return slices.Equal(stream.source, dest)
}

// Find Returns the index of the first element in the stream that matches the target element.
// If not found then -1 is returned.
func (stream SliceComparableStream[E]) Find(dest E) int {
	stream.evaluation()
	for i, v := range stream.source {
		if v == dest {
			return i
		}
	}
	return -1
}

// Parallel See: SliceStream.Parallel
func (stream SliceComparableStream[E]) Parallel(goroutines int) SliceComparableStream[E] {
	stream.SliceStream = stream.SliceStream.Parallel(goroutines)
	return stream
}

// ForEach See: SliceStream.ForEach
func (stream SliceComparableStream[E]) ForEach(action func(int, E)) SliceComparableStream[E] {
	stream.SliceStream.ForEach(action)
	return stream
}

// Filter See: SliceStream.Filter
func (stream SliceComparableStream[E]) Filter(predicate func(E) bool) SliceComparableStream[E] {
	stream.SliceStream = stream.SliceStream.Filter(predicate)
	return stream
}

// Limit See: SliceStream.Limit
func (stream SliceComparableStream[E]) Limit(maxSize int) SliceComparableStream[E] {
	stream.SliceStream = stream.SliceStream.Limit(maxSize)
	return stream
}

// Map See: SliceStream.Map
func (stream SliceComparableStream[E]) Map(mapper func(E) E) SliceComparableStream[E] {
	stream.SliceStream = stream.SliceStream.Map(mapper)
	return stream
}

// SortFunc See: SliceStream.SortFunc
func (stream SliceComparableStream[E]) SortFunc(less func(a, b E) int) SliceComparableStream[E] {
	stream.SliceStream = stream.SliceStream.SortFunc(less)
	return stream
}
