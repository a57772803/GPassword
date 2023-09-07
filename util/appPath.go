package util

import (
	"path/filepath"

	"github.com/shibukawa/configdir"
)

func GetAppDataPath() string {
	configDirs := configdir.New("vendorname", "applicationname")
	configDirs.LocalPath, _ = filepath.Abs(".")
	folders := configDirs.QueryFolders(configdir.Global)

	return folders[0].Path
}
