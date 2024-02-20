package platform

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"boyi/pkg/delivery/graph/view"
	"context"
	"fmt"
)

// CreateMerchant is the resolver for the createMerchant field.
func (r *mutationResolver) CreateMerchant(ctx context.Context, in view.MerchantCreateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: CreateMerchant - createMerchant"))
}

// UpdateMerchant is the resolver for the updateMerchant field.
func (r *mutationResolver) UpdateMerchant(ctx context.Context, filter view.MerchantFilterInput, in view.MerchantUpdateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: UpdateMerchant - updateMerchant"))
}

// DeleteMerchant is the resolver for the deleteMerchant field.
func (r *mutationResolver) DeleteMerchant(ctx context.Context, filter view.MerchantFilterInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: DeleteMerchant - deleteMerchant"))
}

// CreateMerchantWithdrawMethod is the resolver for the createMerchantWithdrawMethod field.
func (r *mutationResolver) CreateMerchantWithdrawMethod(ctx context.Context, in view.MerchantWithdrawMethodCreateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: CreateMerchantWithdrawMethod - createMerchantWithdrawMethod"))
}

// UpdateMerchantWithdrawMethod is the resolver for the updateMerchantWithdrawMethod field.
func (r *mutationResolver) UpdateMerchantWithdrawMethod(ctx context.Context, filter view.MerchantWithdrawMethodFilterInput, in view.MerchantWithdrawMethodUpdateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: UpdateMerchantWithdrawMethod - updateMerchantWithdrawMethod"))
}

// DeleteMerchantWithdrawMethod is the resolver for the deleteMerchantWithdrawMethod field.
func (r *mutationResolver) DeleteMerchantWithdrawMethod(ctx context.Context, filter view.MerchantWithdrawMethodFilterInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: DeleteMerchantWithdrawMethod - deleteMerchantWithdrawMethod"))
}

// CreateMerchantDepositMethod is the resolver for the createMerchantDepositMethod field.
func (r *mutationResolver) CreateMerchantDepositMethod(ctx context.Context, in view.MerchantDepositMethodCreateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: CreateMerchantDepositMethod - createMerchantDepositMethod"))
}

// UpdateMerchantDepositMethod is the resolver for the updateMerchantDepositMethod field.
func (r *mutationResolver) UpdateMerchantDepositMethod(ctx context.Context, filter view.MerchantDepositMethodFilterInput, in view.MerchantDepositMethodUpdateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: UpdateMerchantDepositMethod - updateMerchantDepositMethod"))
}

// DeleteMerchantDepositMethod is the resolver for the deleteMerchantDepositMethod field.
func (r *mutationResolver) DeleteMerchantDepositMethod(ctx context.Context, filter view.MerchantDepositMethodFilterInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: DeleteMerchantDepositMethod - deleteMerchantDepositMethod"))
}

// CreateMerchantFeeMode is the resolver for the createMerchantFeeMode field.
func (r *mutationResolver) CreateMerchantFeeMode(ctx context.Context, in view.MerchantFeeModeCreateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: CreateMerchantFeeMode - createMerchantFeeMode"))
}

// UpdateMerchantFeeMode is the resolver for the updateMerchantFeeMode field.
func (r *mutationResolver) UpdateMerchantFeeMode(ctx context.Context, filter view.MerchantFeeModeFilterInput, in view.MerchantFeeModeUpdateInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: UpdateMerchantFeeMode - updateMerchantFeeMode"))
}

// DeleteMerchantFeeMode is the resolver for the deleteMerchantFeeMode field.
func (r *mutationResolver) DeleteMerchantFeeMode(ctx context.Context, filter view.MerchantFeeModeFilterInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: DeleteMerchantFeeMode - deleteMerchantFeeMode"))
}

// ApplyMerchantBalance is the resolver for the applyMerchantBalance field.
func (r *mutationResolver) ApplyMerchantBalance(ctx context.Context, in view.MerchantBalanceApplyInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: ApplyMerchantBalance - applyMerchantBalance"))
}

// AuditMerchantBalance is the resolver for the auditMerchantBalance field.
func (r *mutationResolver) AuditMerchantBalance(ctx context.Context, filter view.MerchantBalanceFilterInput, in view.MerchantBalanceAuditInput) (uint64, error) {
	panic(fmt.Errorf("not implemented: AuditMerchantBalance - auditMerchantBalance"))
}

// ListMerchant is the resolver for the listMerchant field.
func (r *queryResolver) ListMerchant(ctx context.Context, filter *view.MerchantFilterInput, pagination *view.PaginationInput) (*view.ListMerchantResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchant - listMerchant"))
}

// ListMerchantWithdrawMethod is the resolver for the listMerchantWithdrawMethod field.
func (r *queryResolver) ListMerchantWithdrawMethod(ctx context.Context, filter *view.MerchantWithdrawMethodFilterInput, pagination *view.PaginationInput) (*view.ListMerchantWithdrawMethodResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantWithdrawMethod - listMerchantWithdrawMethod"))
}

// ListMerchantDepositMethod is the resolver for the listMerchantDepositMethod field.
func (r *queryResolver) ListMerchantDepositMethod(ctx context.Context, filter *view.MerchantDepositMethodFilterInput, pagination *view.PaginationInput) (*view.ListMerchantDepositMethodResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantDepositMethod - listMerchantDepositMethod"))
}

// ListMerchantFeeMode is the resolver for the listMerchantFeeMode field.
func (r *queryResolver) ListMerchantFeeMode(ctx context.Context, filter *view.MerchantFeeModeFilterInput, pagination *view.PaginationInput) (*view.ListMerchantFeeModeResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantFeeMode - listMerchantFeeMode"))
}

// ListMerchantBalanceLog is the resolver for the listMerchantBalanceLog field.
func (r *queryResolver) ListMerchantBalanceLog(ctx context.Context, filter *view.MerchantBalanceLogFilterInput, pagination *view.PaginationInput) (*view.ListMerchantBalanceLogResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantBalanceLog - listMerchantBalanceLog"))
}

// ListMerchantWithdrawLog is the resolver for the listMerchantWithdrawLog field.
func (r *queryResolver) ListMerchantWithdrawLog(ctx context.Context, filter *view.MerchantWithdrawLogFilterInput, pagination *view.PaginationInput) (*view.ListMerchantWithdrawLogResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantWithdrawLog - listMerchantWithdrawLog"))
}

// ListMerchantDepositLog is the resolver for the listMerchantDepositLog field.
func (r *queryResolver) ListMerchantDepositLog(ctx context.Context, filter *view.MerchantDepositLogFilterInput, pagination *view.PaginationInput) (*view.ListMerchantDepositLogResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantDepositLog - listMerchantDepositLog"))
}

// ListMerchantLoginHistory is the resolver for the listMerchantLoginHistory field.
func (r *queryResolver) ListMerchantLoginHistory(ctx context.Context, filter *view.MerchantLoginHistoryFilterInput, pagination *view.PaginationInput) (*view.ListMerchantLoginHistoryResp, error) {
	panic(fmt.Errorf("not implemented: ListMerchantLoginHistory - listMerchantLoginHistory"))
}
