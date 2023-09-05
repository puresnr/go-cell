package algorithm

import (
	"github.com/puresnr/go-cell/cast"
	"testing"
)

// 从测试结果来说, slice 传值还是传传指针在效率上几乎没有差别(绝对来说, 指针更快一些, 但是编码也麻烦一些),
// 不管存储的元素是不是指针(底层来说, 不引发重新分配的情况下, 这两者是没有区别的), 但是传值的话, 接收者修改长度时, 接收值和原始值可能就此分离
// 遍历时, 如果 slice 的元素是指针, 那么无论下标取值还是在遍历时同时接收值, 两者时间相差无几, 如果 slice 里的元素是值, 那么如果在遍历时同时接收值, 那么
// 耗时会增加到 12 倍之多, 如果不接收值, 仅通过下标取值, 那么耗时就变得和元素是指针一样了.
// 然后这三个同级别选手, 从实测的绝对速度来说, 存储值, 遍历时仅下标取值, 速度是最快的, 存储下标的两个操作就在伯仲之间了, 都有速度快的情况出现.
// 如果能避免赋值及重新分配内存的开销的话, 可以考虑优先选择存储值, 就是这个场景比较苛刻, 而且遍历时要不停的下标取值, 也挺麻烦的, 不如存储指针, 一概接收值那么方便, 适应性广

type BigStruct struct {
	Ints [512]int
	Strs [512]string
}

var bigStructs []BigStruct
var bigStructsP []*BigStruct

func init() {
	for i := 0; i != 512; i++ {
		bs := BigStruct{}
		for j := 0; j != 512; j++ {
			bs.Ints[j] = j
			bs.Strs[j] = cast.ToString(j)
		}
		bs.Strs[511] = cast.ToString(i)

		bigStructs = append(bigStructs, bs)

		bs1 := bs
		bigStructsP = append(bigStructsP, &bs1)
	}
}

func findbyvalue(sli []BigStruct) *BigStruct {
	for i := range sli {
		if sli[i].Strs[511] == "510" {
			return &sli[i]
		}
	}

	return nil
}

func findbypoint(sli *[]BigStruct) *BigStruct {
	for i := range *sli {
		if (*sli)[i].Strs[511] == "510" {
			return &(*sli)[i]
		}
	}

	return nil
}

func findevv(sli []BigStruct) *BigStruct {
	for i, v := range sli {
		if v.Strs[511] == "510" {
			return &sli[i]
		}
	}

	return nil
}

func findevi(sli []BigStruct) *BigStruct {
	for i := range sli {
		if sli[i].Strs[511] == "510" {
			return &sli[i]
		}
	}

	return nil
}

func findepv(sli []*BigStruct) **BigStruct {
	for i, v := range sli {
		if v.Strs[511] == "510" {
			return &sli[i]
		}
	}

	return nil
}

func findepi(sli []*BigStruct) **BigStruct {
	for i := range sli {
		if sli[i].Strs[511] == "510" {
			return &sli[i]
		}
	}

	return nil
}

func BenchmarkFindbyvalue(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findbyvalue(bigStructs)
	}
}

func BenchmarkFindbypoint(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findbypoint(&bigStructs)
	}
}

func BenchmarkFindEvv(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findevv(bigStructs)
	}
}

func BenchmarkEvi(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findevi(bigStructs)
	}
}

func BenchmarkFindEpv(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findepv(bigStructsP)
	}
}

func BenchmarkEpi(b *testing.B) {
	b.ResetTimer()

	for i := 0; i != b.N; i++ {
		findepi(bigStructsP)
	}
}
