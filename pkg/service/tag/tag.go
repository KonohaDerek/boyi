package tag

import (
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// GetTag 取得Tag的資訊
func (s *service) GetTag(ctx context.Context, opt *option.TagWhereOption) (dto.Tag, error) {
	var (
		tag dto.Tag
	)
	if err := s.repo.Get(ctx, nil, &tag, opt); err != nil {
		return tag, err
	}
	return tag, nil
}

// CreateTag 建立Tag
func (s *service) CreateTag(ctx context.Context, data *dto.Tag) error {
	return s.repo.Create(ctx, nil, data)
}

// ListTags 列出Tag
func (s *service) ListTags(ctx context.Context, opt *option.TagWhereOption) ([]dto.Tag, int64, error) {
	var (
		tags []dto.Tag
	)
	total, err := s.repo.List(ctx, nil, &tags, opt)
	if err != nil {
		return nil, 0, err
	}
	return tags, total, nil
}

// UpdateTag 更新Tag
func (s *service) UpdateTag(ctx context.Context, opt *option.TagWhereOption, col *option.TagUpdateColumn) error {
	err := s.repo.Update(ctx, nil, opt, col)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTag 刪除Tag
func (s *service) DeleteTag(ctx context.Context, opt *option.TagWhereOption) error {
	return s.repo.Delete(ctx, nil, &dto.Tag{}, opt)
}
