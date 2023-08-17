package eaes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	InitModeCBC("2813281781704220", "0428012912989788", EncodeTypeBase64UrlPadding)
}

func TestEncryptModeCBC(t *testing.T) {
	out, _ := EncryptModeCBC("uuid=8252878d-2df5-48b7-89ab-1d9acf58d7a4&cldata=utm_source=google-play&utm_medium=organic")
	fmt.Println(out)
}

func TestDecipherModeCBC(t *testing.T) {
	out, _ := EncryptModeCBC("uuid=8252878d-2df5-48b7-89ab-1d9acf58d7a4&cldata=utm_source=google-play&utm_medium=organic")
	raw, _ := DecipherModeCBC(out)
	assert.Equal(t, true, raw == "uuid=8252878d-2df5-48b7-89ab-1d9acf58d7a4&cldata=utm_source=google-play&utm_medium=organic")
}
