package fsm

import (
	"context"
	"fmt"
	"sync"
)

const (
	// Default represents the default state of the system.
	Default StateType = ""

	// NoOp represents a no-op event.
	NoOp EventType = "NoOp"
)

// StateType
type StateType string

// EventType
type EventType string

// EventPayload
type EventPayload interface{}

// Action represents the action to be executed in a given state
type Action interface {
	Execute(ctx context.Context, payload EventPayload) (EventType, EventPayload, error)
}

// ActionFunc is a function wrapper for Action
type ActionFunc func(ctx context.Context, payload EventPayload) (EventType, EventPayload, error)

func (f ActionFunc) Execute(ctx context.Context, payload EventPayload) (EventType, EventPayload, error) {
	return f(ctx, payload)
}

// Events represents a mapping of events and states.
type Events map[EventType]StateType

// State binds a state with an action and a set of events it can handle.
type State struct {
	Action Action
	Events Events
}

type EventHandler interface {
	Handle(context.Context, StateType) error
}

type EventHandlerFunc func(context.Context, StateType) error

func (f EventHandlerFunc) Handle(ctx context.Context, st StateType) error {
	return f(ctx, st)
}

// States represents a mapping of states
type States map[StateType]State

type StateMachine struct {
	Previous StateType
	Current  StateType
	States   States
	mu       sync.Mutex
	Handlers []EventHandler
}

// getNextState returns the next state for the event given to the machine's current state.
func (s *StateMachine) getNextState(event EventType) (StateType, error) {
	if st, ok := s.States[s.Current]; ok {
		if st.Events != nil {
			if next, ok := st.Events[event]; ok {
				return next, nil
			}
		}
	}
	return Default, fmt.Errorf("%s: %s", "操作被拒", event)
}

// SendEvent send an event to the state machine.
func (s *StateMachine) SendEvent(ctx context.Context, event EventType, payload EventPayload) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for {
		nextState, err := s.getNextState(event)
		if err != nil {
			return fmt.Errorf("%s: %s", "操作被拒", event)
		}

		state, ok := s.States[nextState]
		if !ok {
			return fmt.Errorf("%s: %s", "state not found", event)
		}

		// transition over to the next state.
		// OnLeave, From -> To
		s.Previous = s.Current

		// OnEnter
		s.Current = nextState
		if state.Action != nil {
			nextEvent, evtPayload, err := state.Action.Execute(ctx, payload)
			if err != nil {
				return err
			}
			event = nextEvent
			payload = evtPayload
		}
		// handle state changes
		for _, h := range s.Handlers {
			if err = h.Handle(ctx, s.Current); err != nil {
				return err
			}
		}
		if event == NoOp {
			return nil
		}
	}

}
