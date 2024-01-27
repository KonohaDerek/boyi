package dto

import (
	"fmt"
	"strings"

	"github.com/rs/xid"
)

type FileKey string

func (FileKey) GenerateFileKey(fileName string) FileKey {
	key := xid.New().String()
	tmp := strings.Split(fileName, ".")
	if len(tmp) > 1 {
		key += "." + tmp[len(tmp)-1]
	}

	return FileKey(key)
}

func (key FileKey) String() string {
	return string(key)
}

func (key FileKey) ToURL(fileURI string) string {
	if key == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", fileURI, key)
}

func (key FileKey) Set(fileKey string) FileKey {
	return FileKey(fileKey)
}
