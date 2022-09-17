//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

type MAV_ODID_DESC_TYPE uint32

const (
	// Optional free-form text description of the purpose of the flight.
	MAV_ODID_DESC_TYPE_TEXT MAV_ODID_DESC_TYPE = 0
	// Optional additional clarification when status == MAV_ODID_STATUS_EMERGENCY.
	MAV_ODID_DESC_TYPE_EMERGENCY MAV_ODID_DESC_TYPE = 1
	// Optional additional clarification when status != MAV_ODID_STATUS_EMERGENCY.
	MAV_ODID_DESC_TYPE_EXTENDED_STATUS MAV_ODID_DESC_TYPE = 2
)

var labels_MAV_ODID_DESC_TYPE = map[MAV_ODID_DESC_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_DESC_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_DESC_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_DESC_TYPE = map[string]MAV_ODID_DESC_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_DESC_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_DESC_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_DESC_TYPE) String() string {
	if l, ok := labels_MAV_ODID_DESC_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
