package util

import (
	util "GPassword/util/encrypt"
	"encoding/json"
	"fmt"
	"log"
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

func (user *User) BindAuthSecret(AccountName string) error {

	path := fmt.Sprint(GetAppDataPath() + "/" + AccountName)

	db, _ := GetDB(path)

	RandomSecret := generateTOTPSecret()

	//TODO HASH "RandomSecret" BY AES
	Md5AccountName := (&util.Encrypt{}).Md5(AccountName) //key=> AES(key,Md5(accountName))
	AuthSecret, err := util.AesEncrypt([]byte(Md5AccountName), []byte(RandomSecret))
	log.Print(err)

	err2 := db.Put([]byte("AuthSecret"), AuthSecret, nil)

	log.Print(err2)

	db.Close()

	return nil
}
