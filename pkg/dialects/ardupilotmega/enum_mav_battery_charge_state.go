//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// Enumeration for battery charge states.
type MAV_BATTERY_CHARGE_STATE uint32

const (
	// Low battery state is not provided
	MAV_BATTERY_CHARGE_STATE_UNDEFINED MAV_BATTERY_CHARGE_STATE = 0
	// Battery is not in low state. Normal operation.
	MAV_BATTERY_CHARGE_STATE_OK MAV_BATTERY_CHARGE_STATE = 1
	// Battery state is low, warn and monitor close.
	MAV_BATTERY_CHARGE_STATE_LOW MAV_BATTERY_CHARGE_STATE = 2
	// Battery state is critical, return or abort immediately.
	MAV_BATTERY_CHARGE_STATE_CRITICAL MAV_BATTERY_CHARGE_STATE = 3
	// Battery state is too low for ordinary abort sequence. Perform fastest possible emergency stop to prevent damage.
	MAV_BATTERY_CHARGE_STATE_EMERGENCY MAV_BATTERY_CHARGE_STATE = 4
	// Battery failed, damage unavoidable. Possible causes (faults) are listed in MAV_BATTERY_FAULT.
	MAV_BATTERY_CHARGE_STATE_FAILED MAV_BATTERY_CHARGE_STATE = 5
	// Battery is diagnosed to be defective or an error occurred, usage is discouraged / prohibited. Possible causes (faults) are listed in MAV_BATTERY_FAULT.
	MAV_BATTERY_CHARGE_STATE_UNHEALTHY MAV_BATTERY_CHARGE_STATE = 6
	// Battery is charging.
	MAV_BATTERY_CHARGE_STATE_CHARGING MAV_BATTERY_CHARGE_STATE = 7
)

var labels_MAV_BATTERY_CHARGE_STATE = map[MAV_BATTERY_CHARGE_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_CHARGE_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_BATTERY_CHARGE_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_BATTERY_CHARGE_STATE = map[string]MAV_BATTERY_CHARGE_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_CHARGE_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_BATTERY_CHARGE_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_CHARGE_STATE) String() string {
	if l, ok := labels_MAV_BATTERY_CHARGE_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
