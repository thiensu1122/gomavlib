//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"errors"
)

// Precision land modes (used in MAV_CMD_NAV_LAND).
type PRECISION_LAND_MODE uint32

const (
	// Normal (non-precision) landing.
	PRECISION_LAND_MODE_DISABLED PRECISION_LAND_MODE = 0
	// Use precision landing if beacon detected when land command accepted, otherwise land normally.
	PRECISION_LAND_MODE_OPPORTUNISTIC PRECISION_LAND_MODE = 1
	// Use precision landing, searching for beacon if not found when land command accepted (land normally if beacon cannot be found).
	PRECISION_LAND_MODE_REQUIRED PRECISION_LAND_MODE = 2
)

var labels_PRECISION_LAND_MODE = map[PRECISION_LAND_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PRECISION_LAND_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_PRECISION_LAND_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PRECISION_LAND_MODE = map[string]PRECISION_LAND_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PRECISION_LAND_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PRECISION_LAND_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PRECISION_LAND_MODE) String() string {
	if l, ok := labels_PRECISION_LAND_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
