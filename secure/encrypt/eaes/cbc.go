package eaes

import (
	"encoding/base64"
	"github.com/puresnr/go-cell/cerror"
	"github.com/wumansgy/goEncrypt"
)

type EncodeType byte

const (
	// 表示加密后的输出或者待解密的输入采用何种编码方式
	EncodeTypeRaw EncodeType = iota // 无编码
	// 下面几种类型是 base64 编码
	EncodeTypeBase64StdPadding
	EncodeTypeBase64StdNoPadding
	EncodeTypeBase64UrlPadding
	EncodeTypeBase64UrlNoPadding
)

var encodings = map[EncodeType]*base64.Encoding{
	EncodeTypeBase64StdPadding:   base64.StdEncoding,
	EncodeTypeBase64StdNoPadding: base64.RawStdEncoding,
	EncodeTypeBase64UrlPadding:   base64.URLEncoding,
	EncodeTypeBase64UrlNoPadding: base64.RawURLEncoding,
}

type ModeCBC struct {
	key, iv  []byte
	encoding *base64.Encoding
}

// NewModeCBC 创建一个新的 ModeCBC 实例
// param :
//
//	key: 密钥
//	iv: 初始化向量
//
// encodeType: 编码方式, 表示加密后的输出或者待解密的输入采用何种编码方式, 不传将使用 EncodeTypeBase64StdPadding 格式
// return : 返回一个新的实例, 保证不为空
func NewModeCBC(key, iv string, encodeType ...EncodeType) *ModeCBC {
	m := &ModeCBC{key: []byte(key), iv: []byte(iv)}
	if len(encodeType) != 0 {
		m.encoding = encodings[encodeType[0]]
	} else {
		m.encoding = encodings[EncodeTypeBase64StdPadding]
	}

	return m
}

// Encrypt 对 raw 进行加密
func (e *ModeCBC) Encrypt(raw string) (string, error) {
	data, err := goEncrypt.AesCbcEncrypt([]byte(raw), e.key, e.iv)
	if err != nil {
		return "", cerror.Wrap(err)
	}

	if e.encoding != nil {
		return e.encoding.EncodeToString(data), nil
	}

	return string(data), nil
}

// Decipher 对 raw 进行解密
func (e *ModeCBC) Decipher(raw string) (string, error) {
	var data []byte
	if e.encoding != nil {
		var err error
		if data, err = e.encoding.DecodeString(raw); err != nil {
			return "", cerror.Wrap(err)
		}
	} else {
		data = []byte(raw)
	}

	ret, err := goEncrypt.AesCbcDecrypt(data, e.key, e.iv)
	return string(ret), cerror.Wrap(err)
}

var modeCBC *ModeCBC

// InitModeCBC : 初始化 modeCBC
// param :
//
//	key: 密钥
//	iv: 初始化向量
//
// encodeType: 编码方式, 表示加密后的输出或者待解密的输入采用何种编码方式, 不传将使用 EncodeTypeBase64StdPadding 格式
func InitModeCBC(key, iv string, encodeType ...EncodeType) {
	if len(encodeType) != 0 {
		modeCBC = NewModeCBC(key, iv, encodeType[0])
	} else {
		modeCBC = NewModeCBC(key, iv, EncodeTypeBase64StdPadding)
	}
}

// EncryptModeCBC 对 raw 进行加密
func EncryptModeCBC(raw string) (string, error) {
	return modeCBC.Encrypt(raw)
}

// DecipherModeCBC 对 raw 进行解密
func DecipherModeCBC(raw string) (string, error) {
	return modeCBC.Decipher(raw)
}
