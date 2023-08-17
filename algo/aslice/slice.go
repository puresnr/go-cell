package aslice

const (
	InvalidIdx = -1
)

func Erase[T comparable](sli *[]T, t T) {
	if sli == nil {
		return
	}

	if idx := Find(*sli, t); idx != InvalidIdx {
		*sli = append((*sli)[:idx], (*sli)[idx+1:]...)
	}
}

func EraseIf[T any](sli *[]T, compare func(v T) bool) {
	if sli == nil {
		return
	}

	if idx := FindIf(*sli, compare); idx != InvalidIdx {
		*sli = append((*sli)[:idx], (*sli)[idx+1:]...)
	}
}

func Find[T comparable](sli []T, t T) int {
	for i, v := range sli {
		if v == t {
			return i
		}
	}

	return InvalidIdx
}

func FindIf[T any](sli []T, compare func(v T) bool) int {
	for i, v := range sli {
		if compare(v) {
			return i
		}
	}

	return InvalidIdx
}

func Exist[T comparable](sli []T, t T) bool {
	return Find(sli, t) != InvalidIdx
}

func ExistIf[T any](sli []T, compare func(v T) bool) bool {
	return FindIf(sli, compare) != InvalidIdx
}

// Reverse reverses the order of the elements in sli
func Reverse[T any](sli []T) {
	for i, j := 0, len(sli)-1; i != len(sli)/2; func() { i++; j-- }() {
		sli[i], sli[j] = sli[j], sli[i]
	}
}

// ReverseCopy copies the elements from sli to another slice in reverse order
// return: new slice in reverse order
func ReverseCopy[T any](sli []T) []T {
	rsli := make([]T, len(sli))
	for i := 0; i != len(sli); i++ {
		rsli[len(sli)-1-i] = sli[i]
	}
	return rsli
}
