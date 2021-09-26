// Package bitset
// 本文件定义一个 bitset 的特化版本，仅支持最多8个 bit 位
package bitset

type bitset8 struct {
	// bitset的长度
	size uint8
	// 存储bit
	bits byte
}

func New8(size uint8) (*bitset8, error) {
	if size > 8 {
		return nil, ErrSize
	}

	return &bitset8{size: size}, nil
}

func (b *bitset8) Size() uint8 { return b.size }

func (b *bitset8) Set(pos uint8) error {
	if pos >= b.size {
		return ErrOverFlow
	}

	b.bits |= maskSet8[pos]

	return nil
}

func (b *bitset8) Test(pos uint8) (bool, error) {
	if pos >= b.size {
		return false, ErrOverFlow
	}

	return b.bits&maskSet8[pos] != 0, nil
}
