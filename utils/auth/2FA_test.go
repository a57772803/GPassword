package auth_test

import (
	"GPassword/utils/auth"
	"fmt"
	"testing"

	"github.com/xlzd/gotp"
)

func TestGotp(t *testing.T) {
	randomSecret := gotp.RandomSecret(16)
	fmt.Println("Random secret:", randomSecret)
	auth.GenerateTOTPWithSecret(randomSecret, "", "")
	auth.TestOTPVerify(randomSecret)
}
