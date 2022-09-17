//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"errors"
)

// Direction of VTOL transition
type VTOL_TRANSITION_HEADING uint32

const (
	// Respect the heading configuration of the vehicle.
	VTOL_TRANSITION_HEADING_VEHICLE_DEFAULT VTOL_TRANSITION_HEADING = 0
	// Use the heading pointing towards the next waypoint.
	VTOL_TRANSITION_HEADING_NEXT_WAYPOINT VTOL_TRANSITION_HEADING = 1
	// Use the heading on takeoff (while sitting on the ground).
	VTOL_TRANSITION_HEADING_TAKEOFF VTOL_TRANSITION_HEADING = 2
	// Use the specified heading in parameter 4.
	VTOL_TRANSITION_HEADING_SPECIFIED VTOL_TRANSITION_HEADING = 3
	// Use the current heading when reaching takeoff altitude (potentially facing the wind when weather-vaning is active).
	VTOL_TRANSITION_HEADING_ANY VTOL_TRANSITION_HEADING = 4
)

var labels_VTOL_TRANSITION_HEADING = map[VTOL_TRANSITION_HEADING]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e VTOL_TRANSITION_HEADING) MarshalText() ([]byte, error) {
	if l, ok := labels_VTOL_TRANSITION_HEADING[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_VTOL_TRANSITION_HEADING = map[string]VTOL_TRANSITION_HEADING{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *VTOL_TRANSITION_HEADING) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_VTOL_TRANSITION_HEADING[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e VTOL_TRANSITION_HEADING) String() string {
	if l, ok := labels_VTOL_TRANSITION_HEADING[e]; ok {
		return l
	}
	return "invalid value"
}
