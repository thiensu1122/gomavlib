//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type MAV_ODID_SPEED_ACC int

const (
	// The speed accuracy is unknown.
	MAV_ODID_SPEED_ACC_UNKNOWN MAV_ODID_SPEED_ACC = 0
	// The speed accuracy is smaller than 10 meters per second.
	MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 1
	// The speed accuracy is smaller than 3 meters per second.
	MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 2
	// The speed accuracy is smaller than 1 meters per second.
	MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 3
	// The speed accuracy is smaller than 0.3 meters per second.
	MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 4
)

var labels_MAV_ODID_SPEED_ACC = map[MAV_ODID_SPEED_ACC]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_SPEED_ACC) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_SPEED_ACC[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_SPEED_ACC = map[string]MAV_ODID_SPEED_ACC{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_SPEED_ACC) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_SPEED_ACC[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_SPEED_ACC) String() string {
	if l, ok := labels_MAV_ODID_SPEED_ACC[e]; ok {
		return l
	}
	return "invalid value"
}
