package hub

import (
	"boyi/pkg/model/enums/types"
	"context"
	"sync"

	"boyi/pkg/infra/helper"

	"github.com/rs/zerolog/log"

	"go.uber.org/fx"
)

type Client interface {
	GetUserID() uint64
	GetProtocol() Protocol
	GetToken() string
	GetDeviceUID() string
	GetAccountType() types.AccountType

	Close()
	ReadPump(ctx context.Context)
	WritePump(ctx context.Context)

	SendMsg(str string)
	IsClose() <-chan struct{}
	IsReady() // 已經加入到 Hub
}

var Module = fx.Options(
	fx.Provide(New),
	fx.Invoke(Run),
)

type Hub struct {
	clients     map[Client]struct{}
	userClients map[uint64]map[Client]struct{}
	clientsLock sync.RWMutex

	broadcast chan string

	register chan Client

	unRegister chan Client
}

func New() *Hub {
	h := &Hub{
		broadcast:   make(chan string),
		register:    make(chan Client),
		unRegister:  make(chan Client),
		clients:     make(map[Client]struct{}),
		clientsLock: sync.RWMutex{},
		userClients: make(map[uint64]map[Client]struct{}),
	}
	return h
}

func Run(lc fx.Lifecycle, h *Hub) {
	closeFlag := make(chan struct{}, 1)
	checkCloseFlag := make(chan struct{}, 1)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				defer helper.Recover(ctx)
				defer func() {
					checkCloseFlag <- struct{}{}
				}()
				for {
					select {
					case client := <-h.register:
						h.AddClient(client)
					case client := <-h.unRegister:
						h.RemoveClient(client)
					case message := <-h.broadcast:
						h.Broadcast(message)
					case <-closeFlag:
						log.Ctx(ctx).Info().Msgf("Shutdown connection hub")
						return
					}
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			h.CloseAllConnection()
			log.Ctx(ctx).Info().Msgf("close hub all connection")
			return nil
		},
	})
}
func (h *Hub) Broadcast(msg string) {
	h.clientsLock.RLock()
	defer h.clientsLock.RUnlock()

	for c := range h.clients {
		c.SendMsg(msg)
	}
}

func (h *Hub) AddClient(c Client) {
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()

	h.clients[c] = struct{}{}

	if _, ok := h.userClients[c.GetUserID()]; !ok {
		h.userClients[c.GetUserID()] = make(map[Client]struct{})
	}
	h.userClients[c.GetUserID()][c] = struct{}{}

	c.IsReady()
}

func (h *Hub) RemoveClient(c Client) {
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()
	defer c.Close()

	delete(h.clients, c)
	delete(h.userClients[c.GetUserID()], c)

	if len(h.userClients[c.GetUserID()]) == 0 {
		delete(h.userClients, c.GetUserID())
	}
}

func (h *Hub) AddClientToHub(ctx context.Context, c Client) {
	h.register <- c

	go func() {
		defer helper.Recover(ctx)
		<-ctx.Done()
		h.unRegister <- c
		log.Ctx(ctx).Debug().Msgf("client hub is close")
	}()

	go c.WritePump(ctx)
}

func (h *Hub) CloseAllConnection() {
	for c := range h.clients {
		h.RemoveClient(c)
	}
}

type UserHubInfo struct {
	UserID  uint64
	Devices []UserInfo
}

func (h *Hub) GetAllClients() []UserHubInfo {
	h.clientsLock.RLock()
	defer h.clientsLock.RUnlock()

	out := make([]UserHubInfo, 0, len(h.userClients))
	for userID := range h.userClients {
		var tmp UserHubInfo
		tmp.UserID = userID
		tmp.Devices = make([]UserInfo, 0, len(h.userClients))
		for client := range h.userClients[userID] {
			tmp.Devices = append(tmp.Devices, UserInfo{
				Token:     client.GetToken(),
				DeviceUID: client.GetDeviceUID(),
			})
		}
		out = append(out, tmp)
	}

	return out
}

func (h *Hub) GetHubClientByAccountType(accountType types.AccountType, msg string) []Client {
	result := []Client{}
	for _, clients := range h.userClients {
		for client := range clients {
			if client.GetAccountType() == accountType {
				result = append(result, client)
			}
		}
	}
	return result
}
