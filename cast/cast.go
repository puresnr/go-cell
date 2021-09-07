// Package cast 提供类型相互转换的功能
package cast

import "strconv"

// ToString : equivalent to FormatInt(int64(i), 10).
func ToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

// ToString_64 : equivalent to FormatInt(i, 10).
func ToString_64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToString_64u : equivalent to FormatUint(i, 10).
func ToString_64u(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// ToString_32u : equivalent to FormatUint(uint64(i), 10).
func ToString_32u(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Stoi_64 : 把一个string转换为int64, 不会报错，不能转换的值会设为0
func Stoi_64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// Stoi_64u : 把一个string转换为uint64, 不会报错，不能转换的值会设为0
func Stoi_64u(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

// Stoi : 把一个string转换为int, 不会报错，不能转换的值会设为0
func Stoi(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

// Stoi_32 : 把一个string转换为int32, 不会报错，不能转换的值会设为0
func Stoi_32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int32(i)
}

// Stoi_32u : 把一个string转换为uint32, 不会报错，不能转换的值会设为0
func Stoi_32u(s string) uint32 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return uint32(i)
}

// StoiSlice_64 : 把一个string的slice转换成int64的slice, 转换中不会报错，不能转换的元素会被置为0
func StoiSlice_64(strs []string) []int64 {
	ints := make([]int64, len(strs))
	for idx, s := range strs {
		ints[idx] = Stoi_64(s)
	}

	return ints
}

// StoiSlice_32u : 把一个string的slice转换成uint32的slice, 转换中不会报错，不能转换的元素会被置为0
func StoiSlice_32u(strs []string) []uint32 {
	ints := make([]uint32, len(strs))
	for idx, s := range strs {
		ints[idx] = Stoi_32u(s)
	}

	return ints
}
