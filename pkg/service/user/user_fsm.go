package user

import (
	"boyi/pkg/iface"
	"boyi/pkg/infra/errors"
	"boyi/pkg/infra/fsm"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/enums/types"
	"context"
)

type StateOption func(*state)

func Handlers(h ...StateHandler) StateOption {
	return func(s *state) { s.handlers = append(s.handlers, h...) }
}

type StateHandler func(context.Context, types.UserStatus, dto.User) error

// state represents order state manager, handles all states changes
type state struct {
	User     *dto.User
	svc      iface.IUserService
	Err      error
	handlers []StateHandler
}

func NewState(
	user *dto.User,
	svc iface.IUserService,
	options ...StateOption) *state {
	s := &state{
		User: user,
		// handlers: []StateHandler{svc.SetState},
	}
	for _, option := range options {
		option(s)
	}
	return s
}

var ErrPaymentNotFound = errors.New("payment not found")

func (s *state) Handle(ctx context.Context, st fsm.StateType) error {
	u := *s.User // make copy
	for _, h := range s.handlers {
		if err := h(ctx, u.Status, u); err != nil {
			return err
		}
	}
	return nil
}

type userStateMachine struct {
	*state
	*fsm.StateMachine
}

// NewUserFSM return User State Machine
func NewUserFSM(
	user *dto.User,
	svc iface.IUserService,
	options ...StateOption) *userStateMachine {
	st := NewState(user, svc, options...)
	states := fsm.States{
		fsm.Default: fsm.State{
			Events: fsm.Events{
				VerifyUser: UserUnVerified,
				ActiveUser: UserActived,
			},
		},
		UserUnVerified: fsm.State{
			Action: fsm.ActionFunc(st.onUnVerified),
			Events: fsm.Events{},
		},
		UserActived: fsm.State{
			Action: fsm.ActionFunc(st.onActived),
			Events: fsm.Events{
				LockUser:    UserLocked,
				DisableUser: UserDisabled,
				DeleteUser:  UserDeleted,
			},
		},
	}
	// FIXME: check status if exists
	current, exists := status[user.Status]
	if !exists {
		current = fsm.Default
	}
	sm := &fsm.StateMachine{
		Current:  current,
		States:   states,
		Handlers: []fsm.EventHandler{st},
	}
	return &userStateMachine{
		state:        st,
		StateMachine: sm,
	}
}

var status = map[types.UserStatus]fsm.StateType{
	types.UserStatus__UNKNOWN:    fsm.Default,
	types.UserStatus__UnVerified: UserUnVerified,
	types.UserStatus__Actived:    UserActived,
	types.UserStatus__Locked:     UserLocked,
	types.UserStatus__Disabled:   UserDisabled,
	types.UserStatus__Deleted:    UserDeleted,
}

// StateType
const (
	UserUnVerified fsm.StateType = "UserUnVerified"
	UserActived    fsm.StateType = "UserActived"
	UserLocked     fsm.StateType = "UserLocked"
	UserDisabled   fsm.StateType = "UserDisabled"
	UserDeleted    fsm.StateType = "UserDeleted"
)

// EventType
const (
	VerifyUser  fsm.EventType = "VerifyUser"
	ActiveUser  fsm.EventType = "ActiveUser"
	LockUser    fsm.EventType = "LockUser"
	DisableUser fsm.EventType = "DisableUser"
	DeleteUser  fsm.EventType = "DeleteUser"
)

func (s *state) onUnVerified(ctx context.Context, payload fsm.EventPayload) (fsm.EventType, fsm.EventPayload, error) {
	return fsm.NoOp, nil, nil
}

func (s *state) onActived(ctx context.Context, payload fsm.EventPayload) (fsm.EventType, fsm.EventPayload, error) {
	return fsm.NoOp, nil, nil
}
