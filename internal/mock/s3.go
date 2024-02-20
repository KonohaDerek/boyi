package mock

import (
	"testing"

	"boyi/pkg/infra/storage"

	gomock "github.com/golang/mock/gomock"
)

// NewStorageSvc ...
func NewStorageSvc(t *testing.T) storage.StorageS3 {
	m := gomock.NewController(t)
	mock := NewMockStorageS3(m)

	mock.EXPECT().UploadFile(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UploadFileByReader(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().DeleteFile(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
	mock.EXPECT().UploadFileByBuffer(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(nil)

	return mock
}
