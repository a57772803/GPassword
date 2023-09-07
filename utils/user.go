package utils

import (
	utils "GPassword/utils/encrypt"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

type User struct {
	accountName string
}
type UserInfo struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
	UserMail     string `json:"userMail"`
	AuthSecret   string `json:"authSecret"`
}

func GetUser() *User {
	return &User{}
}
func (user *User) UserLogin(accountName string) {
	AppPath := fmt.Sprint(GetAppDataPath(), "/", accountName)
	log.Print("hello ,", accountName)
	log.Println(AppPath)
	db, _ := GetDB(AppPath)
	res, _ := db.Get([]byte(accountName), nil)
	if res == nil {
		log.Print("用戶不存在")
		return
	}

	user.accountName = accountName

}

// check if A user is Registed
func (user *User) IsRegisted(AccountName string) bool {

	path := fmt.Sprint(GetAppDataPath() + "/" + AccountName)

	db, _ := GetDB(path)

	res, err := db.Get([]byte(AccountName), nil)
	log.Print(err)

	db.Close()

	if res == nil {
		return false
	} else {
		return true
	}
}

func (user *User) UserRegist(UserData string) {
	var userInfo UserInfo
	err := json.Unmarshal([]byte(UserData), &userInfo)

	if err != nil {
		log.Fatal(err)

	}
	path := fmt.Sprint(GetAppDataPath() + "/" + user.accountName)

	db, _ := GetDB(path)
	log.Print(err)
	if !user.IsRegisted(userInfo.UserAccount) {
		//TODO raise UserAlreadyExist
	} else {
		db.Put([]byte("UserAccount"), []byte(userInfo.UserAccount), nil)
		db.Put([]byte("UserPassword"), []byte(userInfo.UserPassword), nil)
		db.Put([]byte("UserMail"), []byte(userInfo.UserMail), nil)
		db.Put([]byte("AuthSecret"), []byte(userInfo.AuthSecret), nil)
	}

	fmt.Println(err)
	fmt.Println(userInfo)
	db.Close()
}

// 生成2FA金鑰並加密存入DB
func (user *User) BindAuthSecret(AccountName string) error {

	path := filepath.Join(GetAppDataPath(), AccountName) //new folder by different account

	db, err := GetDB(path)

	RandomSecret := generateTOTPSecret()

	//TODO HASH "RandomSecret" BY AES
	var encrypt *utils.Encrypt
	Md5AccountName := encrypt.Md5(AccountName) //key=> AES(key,Md5(accountName))
	AuthSecret, err := encrypt.AesEncrypt([]byte(Md5AccountName), []byte(RandomSecret))

	base64Str := base64.StdEncoding.EncodeToString(AuthSecret)
	log.Print(err)

	err2 := db.Put([]byte("AuthSecret"), []byte(base64Str), nil)

	log.Print(err2)

	err3 := db.Put([]byte("ActiveAuth"), []byte("true"), nil)
	log.Print(err3)
	db.Close()

	return nil
}
