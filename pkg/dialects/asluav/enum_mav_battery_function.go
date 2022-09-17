//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"errors"
)

// Enumeration of battery functions
type MAV_BATTERY_FUNCTION uint32

const (
	// Battery function is unknown
	MAV_BATTERY_FUNCTION_UNKNOWN MAV_BATTERY_FUNCTION = 0
	// Battery supports all flight systems
	MAV_BATTERY_FUNCTION_ALL MAV_BATTERY_FUNCTION = 1
	// Battery for the propulsion system
	MAV_BATTERY_FUNCTION_PROPULSION MAV_BATTERY_FUNCTION = 2
	// Avionics battery
	MAV_BATTERY_FUNCTION_AVIONICS MAV_BATTERY_FUNCTION = 3
	// Payload battery
	MAV_BATTERY_TYPE_PAYLOAD MAV_BATTERY_FUNCTION = 4
)

var labels_MAV_BATTERY_FUNCTION = map[MAV_BATTERY_FUNCTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_FUNCTION) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_BATTERY_FUNCTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_BATTERY_FUNCTION = map[string]MAV_BATTERY_FUNCTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_FUNCTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_BATTERY_FUNCTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_FUNCTION) String() string {
	if l, ok := labels_MAV_BATTERY_FUNCTION[e]; ok {
		return l
	}
	return "invalid value"
}
