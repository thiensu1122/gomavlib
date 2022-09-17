//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"errors"
)

type MAV_ODID_HEIGHT_REF uint32

const (
	// The height field is relative to the take-off location.
	MAV_ODID_HEIGHT_REF_OVER_TAKEOFF MAV_ODID_HEIGHT_REF = 0
	// The height field is relative to ground.
	MAV_ODID_HEIGHT_REF_OVER_GROUND MAV_ODID_HEIGHT_REF = 1
)

var labels_MAV_ODID_HEIGHT_REF = map[MAV_ODID_HEIGHT_REF]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_HEIGHT_REF) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_HEIGHT_REF[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_HEIGHT_REF = map[string]MAV_ODID_HEIGHT_REF{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_HEIGHT_REF) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_HEIGHT_REF[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_HEIGHT_REF) String() string {
	if l, ok := labels_MAV_ODID_HEIGHT_REF[e]; ok {
		return l
	}
	return "invalid value"
}
