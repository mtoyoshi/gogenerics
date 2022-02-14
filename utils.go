package main

func mapp[T, U any](target []T, f func(T) U) []U {
	var ret = make([]U, 0, cap(target))
	for _, v := range target {
		ret = append(ret, f(v))
	}
	return ret
}

func flatMap[T, U any](target []T, f func(T) []U) []U {
	var ret = make([]U, 0)
	for _, v := range target {
		for _, vv := range f(v) {
			ret = append(ret, vv)
		}
	}
	return ret
}

func foldLeft[T, U any](target []T, z U, f func(U, T) U) U {
	if len(target) == 0 {
		return z
	}
	return foldLeft(target[1:], f(z, target[0]), f)
}

type Addable interface {
	int | float64
}

func sum[A Addable](target []A) A {
	return foldLeft(target, 0, func(n1, n2 A) A {
		return n1 + n2
	})
}

func filter[T any](target []T, f func(T) bool) []T {
	var ret = make([]T, 0, cap(target))
	for _, v := range target {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func contains[T comparable](target []T, elem T) bool {
	for _, v := range target {
		if elem == v {
			return true
		}
	}
	return false
}

func count[T any](target []T, f func(T) bool) int {
	n := 0
	for _, v := range target {
		if f(v) {
			n += 1
		}
	}
	return n
}

func span[T any](target []T, f func(T) bool) (trueSlice, falseSlice []T) {
	for _, v := range target {
		if f(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return
}
