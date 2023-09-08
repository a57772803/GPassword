package encrypt_test

import (
	"GPassword/utils/encrypt"
	"encoding/base64"
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
	e := encrypt.NewEncrypt()
	assert.Equal(t, e.Md5(oSt), "202cb962ac59075b964b07152d234b701", "they are no the same")

}
func TestAesEncrypt(t *testing.T) {
	key := "202cb962ac59075b964b07152d234b70"
	data := "123"
	aesResult, err := encrypt.NewEncrypt().AesEncrypt([]byte(key), []byte(data))
	log.Print(err)
	log.Print(aesResult)
	log.Println(base64.StdEncoding.EncodeToString(aesResult))
}

func TestAesDecrypt(t *testing.T) {
	key := "202cb962ac59075b964b07152d234b70"
	data := "BpuJ5MN0bY6gUsOkOonP/hv7TvJ+poK6/AJmAv67Ag=="

	decodeData, _ := base64.StdEncoding.DecodeString(data)

	aesResult, err := encrypt.NewEncrypt().AesDecrypt([]byte(key), decodeData)
	log.Print(err)
	log.Print(string(aesResult))
	// log.Println(&base64.Encoding{}.Strict().(aesResult))
}
