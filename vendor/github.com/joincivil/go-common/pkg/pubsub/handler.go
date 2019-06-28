package pubsub

// EventHandler is an interface to a governance event handler.
type EventHandler interface {
	// Handle runs the logic to handle the event as appropriate for the event
	// Returns a bool whether the event was handled and an error if occurred
	Handle(event []byte) (bool, error)
	// Name returns a readable name for this particular event handler
	Name() string
}
