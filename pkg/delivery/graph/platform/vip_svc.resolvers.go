package platform

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"boyi/pkg/delivery/graph/view"
	"context"
	"fmt"
)

// CreateVipLevel is the resolver for the createVipLevel field.
func (r *mutationResolver) CreateVipLevel(ctx context.Context, in view.VipLevelCreateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: CreateVipLevel - createVipLevel"))
}

// UpdateVipLevel is the resolver for the updateVipLevel field.
func (r *mutationResolver) UpdateVipLevel(ctx context.Context, filter view.VipLevelFilterInput, in view.VipLevelUpdateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: UpdateVipLevel - updateVipLevel"))
}

// DeleteVipLevel is the resolver for the deleteVipLevel field.
func (r *mutationResolver) DeleteVipLevel(ctx context.Context, filter view.VipLevelFilterInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: DeleteVipLevel - deleteVipLevel"))
}

// ClaimVipReward is the resolver for the claimVipReward field.
func (r *mutationResolver) ClaimVipReward(ctx context.Context, in view.ClaimVipRewardInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: ClaimVipReward - claimVipReward"))
}

// ListVipUpgradeLog is the resolver for the listVipUpgradeLog field.
func (r *queryResolver) ListVipUpgradeLog(ctx context.Context, filter view.VipUpgradeLogFilterInput, pagination *view.PaginationInput) (*view.ListVipUpgradeLogResp, error) {
	panic(fmt.Errorf("not implemented: ListVipUpgradeLog - listVipUpgradeLog"))
}

// ListVipClaimLog is the resolver for the listVipClaimLog field.
func (r *queryResolver) ListVipClaimLog(ctx context.Context, filter view.VipClaimLogFilterInput, pagination *view.PaginationInput) (*view.ListVipClaimLogResp, error) {
	panic(fmt.Errorf("not implemented: ListVipClaimLog - listVipClaimLog"))
}

// ListVipLevel is the resolver for the listVipLevel field.
func (r *queryResolver) ListVipLevel(ctx context.Context, filter view.VipLevelFilterInput, pagination *view.PaginationInput) (*view.ListVipLevelResp, error) {
	panic(fmt.Errorf("not implemented: ListVipLevel - listVipLevel"))
}
