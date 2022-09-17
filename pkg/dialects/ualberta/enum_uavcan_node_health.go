//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"errors"
)

// Generalized UAVCAN node health
type UAVCAN_NODE_HEALTH uint32

const (
	// The node is functioning properly.
	UAVCAN_NODE_HEALTH_OK UAVCAN_NODE_HEALTH = 0
	// A critical parameter went out of range or the node has encountered a minor failure.
	UAVCAN_NODE_HEALTH_WARNING UAVCAN_NODE_HEALTH = 1
	// The node has encountered a major failure.
	UAVCAN_NODE_HEALTH_ERROR UAVCAN_NODE_HEALTH = 2
	// The node has suffered a fatal malfunction.
	UAVCAN_NODE_HEALTH_CRITICAL UAVCAN_NODE_HEALTH = 3
)

var labels_UAVCAN_NODE_HEALTH = map[UAVCAN_NODE_HEALTH]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVCAN_NODE_HEALTH) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVCAN_NODE_HEALTH[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVCAN_NODE_HEALTH = map[string]UAVCAN_NODE_HEALTH{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVCAN_NODE_HEALTH) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVCAN_NODE_HEALTH[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVCAN_NODE_HEALTH) String() string {
	if l, ok := labels_UAVCAN_NODE_HEALTH[e]; ok {
		return l
	}
	return "invalid value"
}
