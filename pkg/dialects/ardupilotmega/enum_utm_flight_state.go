//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// Airborne status of UAS.
type UTM_FLIGHT_STATE uint32

const (
	// The flight state can't be determined.
	UTM_FLIGHT_STATE_UNKNOWN UTM_FLIGHT_STATE = 1
	// UAS on ground.
	UTM_FLIGHT_STATE_GROUND UTM_FLIGHT_STATE = 2
	// UAS airborne.
	UTM_FLIGHT_STATE_AIRBORNE UTM_FLIGHT_STATE = 3
	// UAS is in an emergency flight state.
	UTM_FLIGHT_STATE_EMERGENCY UTM_FLIGHT_STATE = 16
	// UAS has no active controls.
	UTM_FLIGHT_STATE_NOCTRL UTM_FLIGHT_STATE = 32
)

var labels_UTM_FLIGHT_STATE = map[UTM_FLIGHT_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UTM_FLIGHT_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_UTM_FLIGHT_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UTM_FLIGHT_STATE = map[string]UTM_FLIGHT_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UTM_FLIGHT_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UTM_FLIGHT_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UTM_FLIGHT_STATE) String() string {
	if l, ok := labels_UTM_FLIGHT_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
