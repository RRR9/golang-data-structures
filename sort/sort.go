package sort

import "golang.org/x/exp/constraints"

func Less[Ord constraints.Ordered](v1 Ord, v2 Ord) bool {
	return v1 >= v2
}

func Greater[Ord constraints.Ordered](v1 Ord, v2 Ord) bool {
	return v1 <= v2
}

func Sort[Ord constraints.Ordered](v []Ord, comp func(Ord, Ord) bool) {

}

func SortByKey[T any, Ord constraints.Ordered](arr []T, v []Ord, comp func(Ord, Ord) bool) {

}
