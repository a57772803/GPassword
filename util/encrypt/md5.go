package util

import (
	"crypto/md5"
	"fmt"
	"io"
)

func (e *Encrypt) Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str2 := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(md5Str2)
	return md5Str2
}
