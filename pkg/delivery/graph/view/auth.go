package view

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/vo"
	"strings"

	"context"

	"boyi/pkg/Infra/ctxutil"
)

func (in *RegisterReqInput) ConvertToVo(ctx context.Context) vo.RegisterReq {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)

	return vo.RegisterReq{
		Username:    in.Username,
		Password:    in.Password,
		AccountType: types.AccountType__Member,
	}
}

func (in *LoginReqInput) ConvertToVo(ctx context.Context) vo.LoginReq {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)

	return vo.LoginReq{
		Username: in.Username,
		Password: in.Password,
		Origin:   ctxutil.GetOriginFromContext(ctx),
	}
}

func (c *Claims) FromClaims(in claims.Claims) *Claims {
	if c == nil {
		c = &Claims{}
	}
	c.ID = in.Id
	c.AccountType = AccountTypeFromDTO[types.AccountType(in.AccountType)]
	c.Username = in.Username

	tmp := make(dto.Authority)
	for menuKeys := range in.Competences {
		tmp[dto.ManagerMenuKey(menuKeys)] = struct{}{}
	}
	menus := tmp.GetMenus()

	c.Menu = make([]*Menu, len(menus))
	for i := range menus {
		c.Menu[i] = c.Menu[i].FromDTO(menus[i])
	}

	return c
}

func (c *Claims) FromUser(in dto.User, fileURI string) *Claims {
	if c == nil {
		c = &Claims{}
	}

	c.AliasName = in.AliasName
	if in.AvatarKey != "" {
		c.AvatarURL = in.AvatarKey.ToURL(fileURI)
	}

	return c
}

func (res *CreateCommonUserResp) FromUser(in dto.User, fileURI string) *CreateCommonUserResp {
	if res == nil {
		res = &CreateCommonUserResp{}
	}
	res.User = res.User.FromDTO(&in, fileURI)
	return res
}

func (res *LoginResp) FromClaims(claim claims.Claims) *LoginResp {
	if res == nil {
		res = &LoginResp{}
	}
	res.Token = claim.Token
	return res
}
