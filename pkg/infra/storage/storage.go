package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"boyi/pkg/infra/errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	// restful file url path
	AppFileRestfulURI = "/app/apis/v1/files"
	// restful file url path
	PlatformFileRestfulURI = "/b/apis/v1/files"
)

var S3URL string

type Config struct {
	PublicURL string `mapstructure:"public_url"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Region    string `mapstructure:"region"`
	Bucket    string `mapstructure:"bucket"`
}

type S3 struct {
	session *s3.S3
	bucket  *string
}

type StorageS3 interface {
	ListFiles(ctx context.Context) ([]*s3.Object, error)
	UploadFile(ctx context.Context, key, body string) error
	UploadFileByReader(ctx context.Context, key string, body io.Reader, expire *time.Time) error
	UploadFileByBuffer(ctx context.Context, key string, body *bytes.Buffer, expire time.Duration) error
	GetFile(ctx context.Context, key string) (string, error)
	DeleteFile(ctx context.Context, key string) error
	CreatePreSignedUploadURL(ctx context.Context, fileInfo FileInfo, expire time.Duration) (url string, err error)
}

func New(cfg *Config) (StorageS3, error) {
	S3URL = cfg.PublicURL
	if S3URL == "" {
		S3URL = fmt.Sprintf("https://%s.s3.%s.amazonaws.com", cfg.Bucket, cfg.Region)
	}

	s3Session := s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
	})))

	return &S3{
		session: s3Session,
		bucket:  aws.String(cfg.Bucket),
	}, nil
}

func (s *S3) ListFiles(ctx context.Context) ([]*s3.Object, error) {
	output, err := s.session.ListObjectsV2WithContext(ctx, &s3.ListObjectsV2Input{
		Bucket: s.bucket,
	})
	if err != nil {
		return nil, errors.ConvertAWSError(err)
	}

	return output.Contents, nil
}

func (s *S3) UploadFile(ctx context.Context, key, body string) error {
	_, err := s.session.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(key),
		Body:   strings.NewReader(body),
		ACL:    aws.String(s3.BucketCannedACLPrivate),
	})

	if err != nil {
		return errors.ConvertAWSError(err)
	}
	return nil
}

func (s *S3) UploadFileByBuffer(ctx context.Context, key string, body *bytes.Buffer, expire time.Duration) error {
	_, err := s.session.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:  s.bucket,
		Key:     aws.String(key),
		Body:    bytes.NewReader(body.Bytes()),
		ACL:     aws.String(s3.BucketCannedACLPrivate),
		Expires: aws.Time(time.Now().Add(expire)),
	})

	if err != nil {
		return errors.ConvertAWSError(err)
	}
	return nil
}

func (s *S3) UploadFileByReader(ctx context.Context, key string, body io.Reader, expire *time.Time) error {
	tmp, err := ioutil.ReadAll(body)
	if err != nil {
		return errors.Wrapf(errors.ErrInternalError, "fail to read to byte")
	}
	r := bytes.NewReader(tmp)
	_, err = s.session.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:  s.bucket,
		Key:     aws.String(key),
		Body:    r,
		ACL:     aws.String(s3.BucketCannedACLPrivate),
		Expires: expire,
	})

	if err != nil {
		return errors.ConvertAWSError(err)
	}
	return nil
}

func (s *S3) GetFile(ctx context.Context, key string) (string, error) {
	out, err := s.session.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: s.bucket,
		Key:    &key,
	})
	if err != nil {
		return "", errors.ConvertAWSError(err)
	}
	defer out.Body.Close()

	body, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return "", errors.Wrapf(errors.ErrInternalError, "fail to read body, err: %s", err.Error())
	}

	return string(body), nil
}

func (s *S3) DeleteFile(ctx context.Context, key string) error {
	_, err := s.session.DeleteObjectsWithContext(ctx, &s3.DeleteObjectsInput{
		Bucket: s.bucket,
		Delete: &s3.Delete{
			Objects: []*s3.ObjectIdentifier{
				{
					Key: &key,
				},
			},
		},
	})
	if err != nil {
		return errors.ConvertAWSError(err)
	}
	return nil
}

type FileInfo struct {
	Key           string
	ContentMD5    string
	ContentLength int64
}

func (s *S3) CreatePreSignedUploadURL(ctx context.Context, file FileInfo, expire time.Duration) (url string, err error) {
	// TODO 先拔掉
	objInput := s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    &file.Key,
		// ContentMD5:    &file.ContentMD5,
		// ContentLength: &file.ContentLength,
	}
	if expire != 0 {
		t := time.Now().Add(expire)
		objInput.Expires = &t
	}
	req, _ := s.session.PutObjectRequest(&objInput)

	// Create the pre-signed url with an expiry
	url, err = req.Presign(15 * time.Minute)
	if err != nil {
		return url, errors.ConvertAWSError(err)
	}

	return url, nil
}
