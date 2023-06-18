package containers

type ContainerSize uint64

type Container[T any] interface {
	Contains(T, func(*T, *T) bool) bool
	Size() ContainerSize
}

type Iterator[T any] interface {
	Next() Iterator[T]
	GetVal() *T
	Begin() Iterator[T]
	End() Iterator[T]
}

func Sort[T any](v Iterator[T], comp func(*T, *T) bool) {

}
