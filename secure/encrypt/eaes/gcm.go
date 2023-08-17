package eaes

import (
	"encoding/hex"
	"github.com/puresnr/go-cell/cerror"
	"github.com/puresnr/go-cell/secure"
	"github.com/unknwon/com"
)

var aesGcmKey string

// InitAesGcm : 初始化aesgcm加密模块
// param :
//
//	key: 密钥
//
// return :
//
//	(EcodeSecureSuccess, EcodeSecureKeyLen64)
func InitAesGcm(key string) secure.EcodeSecure {
	if len(key) != 64 {
		return secure.EcodeSecureKeyLen64
	}

	aesGcmKey = key

	return secure.EcodeSecureSuccess
}

func AesGcmEncrypt(data []byte) ([]byte, error) {
	dk, err := hex.DecodeString(aesGcmKey)
	if err != nil {
		return nil, cerror.Wrap(err)
	}

	return com.AESGCMEncrypt(dk, data)
}

func AesGcmDecipher(data []byte) ([]byte, error) {
	dk, err := hex.DecodeString(aesGcmKey)
	if err != nil {
		return nil, cerror.Wrap(err)
	}

	return com.AESGCMDecrypt(dk, data)
}
