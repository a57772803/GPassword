package util

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
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
func (*User) UserLogin(accountName string) {
	fmt.Print("hello", accountName)
}

func (*User) UserRegist(UserData string) {
	var userInfo UserInfo
	err := json.Unmarshal([]byte(UserData), &userInfo)

	db, err := GetDB("")
	if err != nil {
		log.Fatal(err)

	}
	db.Put([]byte("UserAccount"), []byte(userInfo.UserAccount), nil)
	db.Put([]byte("UserPassword"), []byte(userInfo.UserPassword), nil)
	db.Put([]byte("UserMail"), []byte(userInfo.UserMail), nil)
	db.Put([]byte("AuthSecret"), []byte(userInfo.UserAccount), nil)

	fmt.Println(err)
	fmt.Println(userInfo.AuthSecret, "AuthSecret")
	fmt.Println(userInfo)
}
