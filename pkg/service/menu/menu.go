package menu

import (
	"boyi/internal/claims"
	"boyi/pkg/model/dto"
)

func (s *service) GetParsedMenuTree(c claims.Claims) []*dto.Menu {
	return dto.GetParsedMenu(c)
}
