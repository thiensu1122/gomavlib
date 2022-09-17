//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"errors"
)

// Actions following geofence breach.
type FENCE_ACTION uint32

const (
	// Disable fenced mode. If used in a plan this would mean the next fence is disabled.
	FENCE_ACTION_NONE FENCE_ACTION = 0
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED FENCE_ACTION = 1
	// Report fence breach, but don't take action
	FENCE_ACTION_REPORT FENCE_ACTION = 2
	// Fly to geofence MAV_CMD_NAV_FENCE_RETURN_POINT with manual throttle control in GUIDED mode. Note: This action is only supported by ArduPlane, and may not be supported in all versions.
	FENCE_ACTION_GUIDED_THR_PASS FENCE_ACTION = 3
	// Return/RTL mode.
	FENCE_ACTION_RTL FENCE_ACTION = 4
	// Hold at current location.
	FENCE_ACTION_HOLD FENCE_ACTION = 5
	// Termination failsafe. Motors are shut down (some flight stacks may trigger other failsafe actions).
	FENCE_ACTION_TERMINATE FENCE_ACTION = 6
	// Land at current location.
	FENCE_ACTION_LAND FENCE_ACTION = 7
)

var labels_FENCE_ACTION = map[FENCE_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e FENCE_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_FENCE_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_FENCE_ACTION = map[string]FENCE_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *FENCE_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_FENCE_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e FENCE_ACTION) String() string {
	if l, ok := labels_FENCE_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
