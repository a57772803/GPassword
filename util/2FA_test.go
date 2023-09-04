package util

import (
	"fmt"
	"testing"

	"github.com/xlzd/gotp"
)

func TestGotp(t *testing.T) {
	randomSecret := gotp.RandomSecret(16)
	fmt.Println("Random secret:", randomSecret)
	generateTOTPWithSecret(randomSecret, "", "")
	testOTPVerify(randomSecret)
}
