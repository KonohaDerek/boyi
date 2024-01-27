package vo

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
)

type RegisterReq struct {
	Username    string
	Password    string
	AccountType types.AccountType
}

type LoginReq struct {
	Username string
	Password string
	Origin   string
}

type CertificationResp struct {
	Token string
	User  *dto.User
}
