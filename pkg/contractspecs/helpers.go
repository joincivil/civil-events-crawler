package contractspecs

import (
	"fmt"
	"strings"
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
	return !DisableCrawl[key] && EnableListener[key]
}

// FlagKey generates the key used for DisableCrawl and EnableListeners maps
func FlagKey(contractName string, eventType string) string {
	return fmt.Sprintf("%v:%v", strings.ToLower(contractName), strings.ToLower(eventType))
}
