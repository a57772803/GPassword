package util

import (
	"encoding/json"
	"log"
	"path/filepath"
	"testing"

	"github.com/shibukawa/configdir"
)

func TestConfigGet(t *testing.T) {
	var config configdir.Config
	configDirs := configdir.New("vendor-name", "application-name")
	configDirs.LocalPath, _ = filepath.Abs(".")
	folder := configDirs.QueryFolderContainsFile("setting.json")
	if folder != nil {
		data, _ := folder.ReadFile("setting.json")
		json.Unmarshal(data, &config)
	}
}
func TestConfigWrite(t *testing.T) {
	data := `{"aaa":"aaa"}`
	configDirs := configdir.New("vendorname", "applicationname")
	configDirs.LocalPath, _ = filepath.Abs(".")
	folders := configDirs.QueryFolders(configdir.Global)
	err := folders[0].WriteFile("setting.json", []byte(data))
	log.Print(err)
}
