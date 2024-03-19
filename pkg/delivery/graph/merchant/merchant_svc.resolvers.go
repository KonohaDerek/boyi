package merchant

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"boyi/internal/claims"
	"boyi/pkg/delivery/graph/view"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
)

// CreateMerchantUser is the resolver for the CreateMerchantUser field.
func (r *mutationResolver) CreateMerchantUser(ctx context.Context, in view.MerchantUserCreateInput) (*view.MerchantUser, error) {
	var (
		resp *view.MerchantUser
	)
	claims, err := claims.VerifyRole(ctx, dto.API_MerchantUser_Create.String())
	if err != nil {
		return nil, err
	}

	user := in.ConvertToDTO()
	user.CreateUserID = claims.Id
	if err := r.merchantSvc.CreateUser(ctx, &user); err != nil {
		return nil, err
	}

	return resp.FromDTO(user), nil
}

// UpdateMerchantUser is the resolver for the UpdateMerchantUser field.
func (r *mutationResolver) UpdateMerchantUser(ctx context.Context, filter *view.MerchantUserFilterInput, in view.MerchantUserUpdateInput) (*view.MerchantUser, error) {
	var (
		opt  option.MerchantUserWhereOption
		cols option.MerchantUserUpdateColumn
		resp view.MerchantUser
	)
	claims, err := claims.VerifyRole(ctx, dto.API_MerchantUser_Update.String())
	if err != nil {
		return nil, err
	}

	opt = filter.ConvertToOption()
	cols = in.ConvertToOption(&claims)

	result, err := r.merchantSvc.UpdateUser(ctx, &opt, &cols)
	if err != nil {
		return nil, err
	}
	return resp.FromDTO(result), nil
}

// DeleteMerchantUser is the resolver for the DeleteMerchantUser field.
func (r *mutationResolver) DeleteMerchantUser(ctx context.Context, filter *view.MerchantUserFilterInput) (uint64, error) {
	_, err := claims.VerifyRole(ctx, dto.API_MerchantUser_Delete.String())
	if err != nil {
		return 0, err
	}

	var opt = filter.ConvertToOption()
	if err := r.merchantSvc.DeleteUser(ctx, &opt); err != nil {
		return 0, err
	}

	return 1, nil
}

// ListMerchantUser is the resolver for the ListMerchantUser field.
func (r *queryResolver) ListMerchantUser(ctx context.Context, filter *view.MerchantUserFilterInput, pagination *view.PaginationInput) (*view.ListMerchantUserResp, error) {
	var (
		opt  option.MerchantUserWhereOption
		resp view.ListMerchantUserResp
	)

	_, err := claims.VerifyRole(ctx, dto.API_MerchantUser_Get.String())
	if err != nil {
		return nil, err
	}

	opt = filter.ConvertToOption()
	if pagination != nil {
		opt.Pagination = pagination.ConvertToPagination()
	}

	result, total, err := r.merchantSvc.ListUsers(ctx, &opt)
	if err != nil {
		return nil, err
	}

	resp.Meta = &view.Meta{
		Total: uint64(total),
	}
	resp.List = make([]*view.MerchantUser, 0, len(result))
	viewUsers := make([]view.MerchantUser, len(result))
	for i := range result {
		resp.List = append(resp.List, viewUsers[i].FromDTO(result[i]))
	}

	return &resp, nil
}
