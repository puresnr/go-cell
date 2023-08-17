// Package cast 提供类型相互转换的功能
package cast

import (
	"github.com/puresnr/go-cell/generic"
	"strconv"
)

/*
变量转字符串:
	性能测试:
		1. FormatInt 直接指定类型肯定比泛型快, 但是时间相差无几, 不足1纳秒, 但是泛型写着更简单
		2. FormatInt 比 Sprint 之类的快 4 倍
		3. Sprint 和 Sprintf 不相伯仲, 不能说哪个一定更快, 然后 Sprintf 中 %d 和 %v 也是, 不能说哪个更快. 最后, 一般来说泛型更慢, 而且这种情况下用 interface 就行
	总结:
		1. 整数类用 FormatInt 泛型
		2. 其它的用 Sprint
*/

func ToString[T generic.Signed](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func ToString_u[T generic.Unsigned](i T) string {
	return strconv.FormatUint(uint64(i), 10)
}

func ToStringBase[T1 generic.Signed, T2 generic.Integer](i T1, base T2) string {
	return strconv.FormatInt(int64(i), int(base))
}

func ToStringBase_u[T1 generic.Unsigned, T2 generic.Integer](i T1, base T2) string {
	return strconv.FormatUint(uint64(i), int(base))
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

// StoiSlice_64u : 把一个string的slice转换成uint64的slice, 转换中不会报错，不能转换的元素会被置为0
func StoiSlice_64u(strs []string) []uint64 {
	ints := make([]uint64, len(strs))
	for idx, s := range strs {
		ints[idx] = Stoi_64u(s)
	}

	return ints
}

func ByteToIntSlice_32u(bs []byte) []uint32 {
	is := make([]uint32, len(bs))
	for i, v := range bs {
		is[i] = uint32(v)
	}

	return is
}

func IntToByteSlice_32u(is []uint32) []byte {
	bs := make([]byte, len(is))
	for i, v := range is {
		bs[i] = byte(v)
	}

	return bs
}
