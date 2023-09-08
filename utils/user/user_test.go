package user_test

import (
	"GPassword/utils/appPath"
	"GPassword/utils/db"
	"GPassword/utils/encrypt"
	"GPassword/utils/user"
	"encoding/base64"
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	accountName := "testAccount"
	user.GetUser().UserLogin(accountName)
}
func TestRegist(t *testing.T) {

	UserData := `{"userAccount":"Account",
"userPassword":"Password-hashed",
"userMail":"userEmail",
"AuthSecret":"hashed by Account,to verify OTP",
"ActiveAuth":"check Auth is Active"
}`
	user.GetUser().UserRegist(UserData)
	db, err := db.GetDB(appPath.GetAppDataPath())
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

func TestBindAuthSecret(t *testing.T) {
	AccountName := "testUser"
	err := user.GetUser().BindAuthSecret(AccountName)

	assert.Nil(t, err)
}
func TestFetchAuthSecret(t *testing.T) {

	AccountName := "testUser"

	path := filepath.Join(appPath.GetAppDataPath(), AccountName) //new folder by different account

	Md5AccountName := encrypt.NewEncrypt().Md5(AccountName)
	db, err := db.GetDB(path)

	AuthSecret, err := db.Get([]byte("AuthSecret"), nil)
	log.Print(err)

	ByteEncryptSecret, _ := base64.StdEncoding.DecodeString(string(AuthSecret))

	OrigialSecret, _ := encrypt.NewEncrypt().AesDecrypt([]byte(Md5AccountName), ByteEncryptSecret)

	log.Println("original AuthKey is:", string(OrigialSecret))

	db.Close()
	assert.Nil(t, err)
}

func TestBool(t *testing.T) {
}
