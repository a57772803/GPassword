package util

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var user User

func TestUserLogin(t *testing.T) {
	accountName := "testAccount"
	user.UserLogin(accountName)
}
func TestRegist(t *testing.T) {

	UserData := `{"userAccount":"Account",
"userPassword":"Password-hashed",
"userMail":"userEmail",
"AuthSecret":"hashed by Account,to verify OTP",
"ActiveAuth":"check Auth is Active"
}`
	user.UserRegist(UserData)
	db, err := GetDB("")
	if err != nil {
		log.Fatal(err)

	}
	UserAccount, _ := db.Get([]byte("UserAccount"), nil)
	UserPassword, _ := db.Get([]byte("UserPassword"), nil)
	UserMail, _ := db.Get([]byte("UserMail"), nil)
	AuthSecret, _ := db.Get([]byte("AuthSecret"), nil)

	assert.NotEmpty(t, UserAccount)
	assert.NotEmpty(t, UserPassword)
	assert.NotEmpty(t, UserMail)
	assert.NotEmpty(t, AuthSecret)

}
