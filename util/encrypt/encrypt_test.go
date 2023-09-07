package util

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	oSt := "123"
	// if Md5(oSt) == "202cb962ac59075b964b07152d234b70" {
	// 	fmt.Print(oSt)
	// } else {
	// 	t.Error("wrong result")

	// }
	e := new(Encrypt)
	assert.Equal(t, e.Md5(oSt), "202cb962ac59075b964b07152d234b701", "they are no the same")

}
func TestAesEncrypt(t *testing.T) {
	key := "444444"
	data := "123"

	aesResult, err := AesEncrypt([]byte(key), []byte(data))
	log.Print(err)
	log.Print(string(aesResult))
}
