//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// Type of mission items being requested/sent in mission protocol.
type MAV_MISSION_TYPE int

const (
	// Items are mission commands for main mission.
	MAV_MISSION_TYPE_MISSION MAV_MISSION_TYPE = 0
	// Specifies GeoFence area(s). Items are MAV_CMD_NAV_FENCE_ GeoFence items.
	MAV_MISSION_TYPE_FENCE MAV_MISSION_TYPE = 1
	// Specifies the rally points for the vehicle. Rally points are alternative RTL points. Items are MAV_CMD_NAV_RALLY_POINT rally point items.
	MAV_MISSION_TYPE_RALLY MAV_MISSION_TYPE = 2
	// Only used in MISSION_CLEAR_ALL to clear all mission types.
	MAV_MISSION_TYPE_ALL MAV_MISSION_TYPE = 255
)

var labels_MAV_MISSION_TYPE = map[MAV_MISSION_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MISSION_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_MISSION_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_MISSION_TYPE = map[string]MAV_MISSION_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MISSION_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_MISSION_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_MISSION_TYPE) String() string {
	if l, ok := labels_MAV_MISSION_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
