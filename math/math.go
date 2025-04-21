package math

import "golang.org/x/exp/constraints"

/*type SumConstraint interface {
	int | float64 | string
}*/

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}