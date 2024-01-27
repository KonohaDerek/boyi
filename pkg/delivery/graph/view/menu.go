package view

import (
	"boyi/pkg/model/dto"

	"github.com/rs/zerolog/log"
)

const (
	MaxMenuDeep = 5
)

func (in *MenuInput) ConvertToDTO(deep int) (result dto.Menu) {
	if in == nil {
		return
	}

	if deep == MaxMenuDeep {
		log.Warn().Msgf("menu tree is to deep")
		return
	}

	result.Key = dto.ManagerMenuKey(in.Key)

	result.Next = make([]*dto.Menu, len(in.Next))
	for i := range result.Next {
		tmp := in.Next[i].ConvertToDTO(deep + 1)
		result.Next[i] = &tmp
	}

	return
}

func (result *Menu) FromDTO(in *dto.Menu) *Menu {
	if result == nil {
		result = &Menu{}
	}

	result.Name = in.Name
	result.Key = in.Key.String()
	result.SuperKey = in.SuperKey.String()

	result.Next = make([]*Menu, len(in.Next))
	for i := range result.Next {
		result.Next[i] = &Menu{}
		result.Next[i] = result.Next[i].FromDTO(in.Next[i])
	}

	return result
}
