package contractspecs

import (
	"fmt"
	"strings"
)

const (
	// EnableAllListenersKey is the key that enables all listeners in the
	// EnableListener map
	EnableAllListenersKey = "all"
)

// IsEventDisabled is a convenience func that returns true if the given
// contract/event type is disabled
func IsEventDisabled(contractName string, eventType string) bool {
	return DisableCrawl[FlagKey(contractName, eventType)]
}

// IsListenerEnabledForEvent is a convenience func that returns true if the given
// contract/event type is enabled for websocket/eth-subscribe
func IsListenerEnabledForEvent(contractName string, eventType string) bool {
	key := FlagKey(contractName, eventType)

	// Check for "enable all" key in the map
	enableAll, ok := EnableListener[EnableAllListenersKey]
	if ok && enableAll {
		return !DisableCrawl[key]
	}

	return !DisableCrawl[key] && EnableListener[key]
}

// FlagKey generates the key used for DisableCrawl and EnableListeners maps
func FlagKey(contractName string, eventType string) string {
	return fmt.Sprintf("%v:%v", strings.ToLower(contractName), strings.ToLower(eventType))
}
