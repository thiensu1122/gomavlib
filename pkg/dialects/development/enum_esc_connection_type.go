//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package development

import (
	"errors"
)

// Indicates the ESC connection type.
type ESC_CONNECTION_TYPE int

const (
	// Traditional PPM ESC.
	ESC_CONNECTION_TYPE_PPM ESC_CONNECTION_TYPE = 0
	// Serial Bus connected ESC.
	ESC_CONNECTION_TYPE_SERIAL ESC_CONNECTION_TYPE = 1
	// One Shot PPM ESC.
	ESC_CONNECTION_TYPE_ONESHOT ESC_CONNECTION_TYPE = 2
	// I2C ESC.
	ESC_CONNECTION_TYPE_I2C ESC_CONNECTION_TYPE = 3
	// CAN-Bus ESC.
	ESC_CONNECTION_TYPE_CAN ESC_CONNECTION_TYPE = 4
	// DShot ESC.
	ESC_CONNECTION_TYPE_DSHOT ESC_CONNECTION_TYPE = 5
)

var labels_ESC_CONNECTION_TYPE = map[ESC_CONNECTION_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ESC_CONNECTION_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_ESC_CONNECTION_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ESC_CONNECTION_TYPE = map[string]ESC_CONNECTION_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ESC_CONNECTION_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ESC_CONNECTION_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ESC_CONNECTION_TYPE) String() string {
	if l, ok := labels_ESC_CONNECTION_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
