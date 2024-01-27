package errors

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

// GQLErrorPresenter ...
func GQLErrorPresenter(ctx context.Context, err error) *gqlerror.Error {
	var (
		logger    = log.Ctx(ctx).With().Logger()
		unwrapErr error
		path      = graphql.GetPath(ctx)
	)

	gqlErr, ok := err.(*gqlerror.Error)
	if ok {
		unwrapErr = gqlErr.Unwrap()
	} else {
		if path != nil {
			gqlErr = gqlerror.WrapPath(path, err)
		} else {
			gqlErr = gqlerror.Errorf(err.Error())
		}
	}

	if unwrapErr == nil {
		unwrapErr = err
	}

	if gqlErr.Extensions == nil {
		gqlErr.Extensions = make(map[string]interface{})
	}

	causeErr := Cause(unwrapErr)
	_err, ok := causeErr.(*_error)
	if !ok || _err == nil {
		gqlErr.Extensions["raw_err_msg"] = unwrapErr.Error()
		if strings.Contains(unwrapErr.Error(), "input") {
			_err = ErrInvalidInput
		} else {
			_err = ErrInternalError
		}
	}

	if _err.Details != nil {
		for k, v := range _err.Details {
			gqlErr.Extensions[k] = v
		}
	}

	gqlErr.Extensions["msg"] = gqlErr.Message
	gqlErr.Message = _err.Code
	gqlErr.Extensions["err_msg"] = _err.Message
	if Is(gqlErr, ErrInternalError) {
		gqlErr.Extensions["err_msg"] = ErrInternalServerError.Message
	}

	logger = logger.With().
		Bool("access_log", true).
		Interface("gql_err", gqlErr.Extensions).Logger()

	switch {
	case _err.Status >= 400 && _err.Status < 500:
		logger.Warn().Msgf("error handle: %+v", unwrapErr)
	default:
		logger.Error().Msgf("error handle: %+v", unwrapErr)
	}

	return gqlErr
}

// ConvertMySQLError convert mysql error
func ConvertMySQLError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithStack(ErrResourceNotFound)
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if mysqlErr.Number == 1062 {
			// the duplicate key error.
			return errors.WithStack(ErrResourceAlreadyExists)
		}
	}

	if strings.Contains(err.Error(), context.Canceled.Error()) {
		return errors.Wrapf(ErrContextCancel, "user cancel the connection")
	}

	if errors.Is(err, context.Canceled) {
		return errors.Wrapf(ErrContextCancel, "user cancel the connection")
	}

	return errors.Wrapf(ErrInternalError, err.Error())
}

func ConvertAWSError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, context.Canceled) {
		return errors.Wrapf(ErrContextCancel, "user cancel the connection")
	}

	aerr, ok := err.(awserr.Error)
	if !ok {
		return errors.Wrapf(ErrInternalError, "can't convert aws error, err: %s", err.Error())
	}

	switch aerr.Code() {
	case s3.ErrCodeNoSuchKey:
		return errors.Wrapf(ErrResourceNotFound, "no such key, err: %s", err.Error())
	case request.WaiterResourceNotReadyErrorCode:
		return errors.Wrapf(ErrBadRequest, "waiter resource not ready, err: %s", err.Error())
	case request.CanceledErrorCode:
		return errors.Wrapf(ErrContextCancel, "user cancel the connection")
	default:
		return errors.Wrapf(ErrInternalError, "aws error, err: %s", err.Error())
	}
}

func ConvertRedisError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, redis.Nil) {
		return errors.Wrapf(ErrResourceNotFound, "redis key not found")
	}

	if errors.Is(err, context.Canceled) {
		return errors.Wrapf(ErrBadRequest, "user cancel the connection")
	}

	return errors.Wrapf(ErrInternalError, "redis err, %+v", err)
}

func ConvertMongoError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return errors.Wrapf(ErrResourceNotFound, "document not found")
	}
	if errors.Is(err, mongo.ErrClientDisconnected) {
		return errors.Wrapf(ErrContextCancel, "client disconnected")
	}
	if errors.Is(err, context.Canceled) {
		return errors.Wrapf(ErrContextCancel, "context cancel")
	}

	return errors.Wrapf(ErrInternalError, "mongo err: %+v", err)
}
