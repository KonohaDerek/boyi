package vo

import "boyi/pkg/model/dto"

type FileInfo struct {
	Name string
	Key  dto.FileKey
	URL  string
	Size int64
	MD5  string
}
