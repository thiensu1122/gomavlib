//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

type MAV_STATE int

const (
	// Uninitialized system, state is unknown.
	MAV_STATE_UNINIT MAV_STATE = 0
	// System is booting up.
	MAV_STATE_BOOT MAV_STATE = 1
	// System is calibrating and not flight-ready.
	MAV_STATE_CALIBRATING MAV_STATE = 2
	// System is grounded and on standby. It can be launched any time.
	MAV_STATE_STANDBY MAV_STATE = 3
	// System is active and might be already airborne. Motors are engaged.
	MAV_STATE_ACTIVE MAV_STATE = 4
	// System is in a non-normal flight mode. It can however still navigate.
	MAV_STATE_CRITICAL MAV_STATE = 5
	// System is in a non-normal flight mode. It lost control over parts or over the whole airframe. It is in mayday and going down.
	MAV_STATE_EMERGENCY MAV_STATE = 6
	// System just initialized its power-down sequence, will shut down now.
	MAV_STATE_POWEROFF MAV_STATE = 7
	// System is terminating itself.
	MAV_STATE_FLIGHT_TERMINATION MAV_STATE = 8
)

var labels_MAV_STATE = map[MAV_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STATE = map[string]MAV_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STATE) String() string {
	if l, ok := labels_MAV_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
