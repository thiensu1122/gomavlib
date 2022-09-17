//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// SERIAL_CONTROL device types
type SERIAL_CONTROL_DEV uint32

const (
	// First telemetry port
	SERIAL_CONTROL_DEV_TELEM1 SERIAL_CONTROL_DEV = 0
	// Second telemetry port
	SERIAL_CONTROL_DEV_TELEM2 SERIAL_CONTROL_DEV = 1
	// First GPS port
	SERIAL_CONTROL_DEV_GPS1 SERIAL_CONTROL_DEV = 2
	// Second GPS port
	SERIAL_CONTROL_DEV_GPS2 SERIAL_CONTROL_DEV = 3
	// system shell
	SERIAL_CONTROL_DEV_SHELL SERIAL_CONTROL_DEV = 10
	// SERIAL0
	SERIAL_CONTROL_SERIAL0 SERIAL_CONTROL_DEV = 100
	// SERIAL1
	SERIAL_CONTROL_SERIAL1 SERIAL_CONTROL_DEV = 101
	// SERIAL2
	SERIAL_CONTROL_SERIAL2 SERIAL_CONTROL_DEV = 102
	// SERIAL3
	SERIAL_CONTROL_SERIAL3 SERIAL_CONTROL_DEV = 103
	// SERIAL4
	SERIAL_CONTROL_SERIAL4 SERIAL_CONTROL_DEV = 104
	// SERIAL5
	SERIAL_CONTROL_SERIAL5 SERIAL_CONTROL_DEV = 105
	// SERIAL6
	SERIAL_CONTROL_SERIAL6 SERIAL_CONTROL_DEV = 106
	// SERIAL7
	SERIAL_CONTROL_SERIAL7 SERIAL_CONTROL_DEV = 107
	// SERIAL8
	SERIAL_CONTROL_SERIAL8 SERIAL_CONTROL_DEV = 108
	// SERIAL9
	SERIAL_CONTROL_SERIAL9 SERIAL_CONTROL_DEV = 109
)

var labels_SERIAL_CONTROL_DEV = map[SERIAL_CONTROL_DEV]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e SERIAL_CONTROL_DEV) MarshalText() ([]byte, error) {
	if l, ok := labels_SERIAL_CONTROL_DEV[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_SERIAL_CONTROL_DEV = map[string]SERIAL_CONTROL_DEV{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *SERIAL_CONTROL_DEV) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_SERIAL_CONTROL_DEV[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e SERIAL_CONTROL_DEV) String() string {
	if l, ok := labels_SERIAL_CONTROL_DEV[e]; ok {
		return l
	}
	return "invalid value"
}
