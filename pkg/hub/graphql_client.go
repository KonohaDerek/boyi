package hub

import (
	"boyi/pkg/model/enums/types"
	"context"

	"boyi/pkg/Infra/errors"
	"boyi/pkg/Infra/helper"

	"github.com/rs/zerolog/log"
)

type Protocol string

const (
	GRPC      Protocol = "grpc"
	GraphQL   Protocol = "graphql"
	WebSocket Protocol = "websocket"
)

type GraphQLClient struct {
	*UserInfo
	send chan string

	readyFlag chan struct{}
	closeFlag chan struct{}
	// readCloseFlag  chan struct{}
	writeCloseFlag chan struct{}
	OnWrite        func(msg string) error
}

type UserInfo struct {
	UserID      uint64
	Token       string
	DeviceUID   string
	AccountType types.AccountType
	ConnCtx     context.Context
	CloseFunc   func()
}

func newGraphQLClient(option clientOptions) Client {
	c := &GraphQLClient{
		UserInfo:  option.userInfo,
		OnWrite:   option.onWrite,
		send:      make(chan string, 256),
		readyFlag: make(chan struct{}, 1),
		closeFlag: make(chan struct{}, 1),
		// readCloseFlag:  make(chan struct{}, 1),
		writeCloseFlag: make(chan struct{}, 1),
	}
	return c
}
func (c *GraphQLClient) GetUserID() uint64 {
	return c.UserID
}

func (c *GraphQLClient) GetProtocol() Protocol {
	return GraphQL
}

func (c *GraphQLClient) GetAccountType() types.AccountType {
	return c.AccountType
}

func (c *GraphQLClient) Close() {
	// c.readCloseFlag <- struct{}{}
	c.writeCloseFlag <- struct{}{}
	c.closeFlag <- struct{}{}
}

func (c *GraphQLClient) IsClose() <-chan struct{} {
	return c.closeFlag
}

func (c *GraphQLClient) ReadPump(ctx context.Context) {
	// defer helper.Recover(ctx)
	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		log.Ctx(ctx).Debug().Msg("close read connection")
	// 		return
	// 	case <-c.readCloseFlag:
	// 		log.Ctx(ctx).Debug().Msg("close read connection")
	// 	}
	// }
}

func (c *GraphQLClient) SendMsg(msg string) {
	c.send <- msg
}

func (c *GraphQLClient) WritePump(ctx context.Context) {
	defer helper.Recover(ctx)
	defer c.UserInfo.CloseFunc()
	msg := "Hi"

	err := c.OnWrite(msg)
	if err != nil {
		errors.LogError(ctx, err, "fail to write msg to client, err: %+v", err)
	}

	for {
		select {
		case msg := <-c.send:
			err := c.OnWrite(msg)
			if err != nil {
				errors.LogError(ctx, err, "fail to write msg to client, err: %+v", err)
			}
		case <-c.writeCloseFlag:
			log.Ctx(ctx).Debug().Msg("close write connection")
			close(c.send)
			return
		}
	}
}

func (c *GraphQLClient) GetToken() string {
	return c.Token
}
func (c *GraphQLClient) GetDeviceUID() string {
	return c.DeviceUID
}

func (c *GraphQLClient) IsReady() {
	c.readyFlag <- struct{}{}
}
