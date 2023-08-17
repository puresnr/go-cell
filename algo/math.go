package algo

import "golang.org/x/exp/constraints"

// Min 返回两个值中较小的一个，相同时，返回第一个
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v2
	}

	return v1
}

// Max 返回两个值中较大的一个，相同时，返回第一个
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v2
	}

	return v1
}

// MultiplyS_32u 安全的乘法, 如果没有溢出, 则返回(相乘后的值, true), 溢出则返回(溢出后的值, false)
func MultiplyS_32u(v1, v2 uint32) (uint32, bool) {
	if v2 == 0 {
		return 0, true
	}

	v := v1 * v2
	return v, v/v2 == v1
}

// Pick 根据 compare 的结果返回元素, 如果是 true 返回前者, 否则返回后者
// todo: Pick 无法实现短路特性, 传参总是会被计算
func Pick[T any](t1 T, t2 T, compare func() bool) T {
	if compare() {
		return t1
	}
	return t2
}
