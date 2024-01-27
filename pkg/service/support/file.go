package support

import (
	"boyi/pkg/model/vo"
	"context"
	"time"

	"boyi/pkg/Infra/errors"
	"boyi/pkg/Infra/storage"
)

// CreateUploadURL 預先產生上傳 URL, 直接透過 S3 上傳
func (s *service) CreateUploadURL(ctx context.Context, in []vo.FileInfo, expire time.Duration) ([]vo.FileInfo, error) {
	var err error

	for i := range in {
		if in[i].Size == 0 {
			return nil, errors.NewWithMessagef(errors.ErrInvalidInput, "index [%d], size 0", i+1)
		} else if in[i].MD5 == "" {
			return nil, errors.NewWithMessagef(errors.ErrInvalidInput, "index [%d], MD5 empty", i+1)
		}
	}

	for i := range in {
		in[i].Key = in[i].Key.GenerateFileKey(in[i].Name)
		in[i].URL, err = s.s3Svc.CreatePreSignedUploadURL(ctx, storage.FileInfo{
			Key:           in[i].Key.String(),
			ContentMD5:    in[i].MD5,
			ContentLength: in[i].Size,
		}, expire)
		if err != nil {
			return nil, err
		}
	}

	return in, nil
}
