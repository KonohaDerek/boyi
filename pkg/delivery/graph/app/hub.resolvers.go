package app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"boyi/pkg/delivery/graph/view"
	"context"
	"fmt"
)

// ReceiveMessage is the resolver for the receiveMessage field.
func (r *subscriptionResolver) ReceiveMessage(ctx context.Context, userAuth view.UserAuth) (<-chan string, error) {
	panic(fmt.Errorf("not implemented: ReceiveMessage - receiveMessage"))
}
