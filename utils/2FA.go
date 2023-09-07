package utils

import (
	"fmt"
	"time"

	"github.com/xlzd/gotp"
)

// https://medium.com/@kittipat_1413/implementing-two-factor-authentication-2fa-with-totp-in-golang-1a3fd0dd7662

func generateTOTPSecret() string {
	randomSecret := gotp.RandomSecret(16)
	return randomSecret
}

func generateTOTPWithSecret(randomSecret string, userID string, userMail string) string {

	totp := gotp.NewDefaultTOTP(randomSecret)
	fmt.Println("current one-time password is:", totp.Now())

	uri := totp.ProvisioningUri(userMail, userID)
	fmt.Println(uri)

	return uri
}

func testOTPVerify(randomSecret string) {
	totp := gotp.NewDefaultTOTP(randomSecret)
	otpValue := totp.Now()
	fmt.Println("current one-time password is:", otpValue)

	ok := totp.Verify(otpValue, time.Now().Unix())
	fmt.Println("verify OTP success:", ok)
}
