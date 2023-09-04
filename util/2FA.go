package util

import (
	"fmt"
	"time"

	"github.com/skip2/go-qrcode"
	"github.com/xlzd/gotp"
)

func generateTOTPWithSecret(randomSecret string, userID string, userMail string) {
	totp := gotp.NewDefaultTOTP(randomSecret)
	fmt.Println("current one-time password is:", totp.Now())

	uri := totp.ProvisioningUri(userMail, userID)
	fmt.Println(uri)

	qrcode.WriteFile(uri, qrcode.Medium, 256, "qr.png")
}

func testOTPVerify(randomSecret string) {
	totp := gotp.NewDefaultTOTP(randomSecret)
	otpValue := totp.Now()
	fmt.Println("current one-time password is:", otpValue)

	ok := totp.Verify(otpValue, time.Now().Unix())
	fmt.Println("verify OTP success:", ok)
}
