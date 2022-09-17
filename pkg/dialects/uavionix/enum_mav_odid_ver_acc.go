//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"errors"
)

type MAV_ODID_VER_ACC uint32

const (
	// The vertical accuracy is unknown.
	MAV_ODID_VER_ACC_UNKNOWN MAV_ODID_VER_ACC = 0
	// The vertical accuracy is smaller than 150 meter.
	MAV_ODID_VER_ACC_150_METER MAV_ODID_VER_ACC = 1
	// The vertical accuracy is smaller than 45 meter.
	MAV_ODID_VER_ACC_45_METER MAV_ODID_VER_ACC = 2
	// The vertical accuracy is smaller than 25 meter.
	MAV_ODID_VER_ACC_25_METER MAV_ODID_VER_ACC = 3
	// The vertical accuracy is smaller than 10 meter.
	MAV_ODID_VER_ACC_10_METER MAV_ODID_VER_ACC = 4
	// The vertical accuracy is smaller than 3 meter.
	MAV_ODID_VER_ACC_3_METER MAV_ODID_VER_ACC = 5
	// The vertical accuracy is smaller than 1 meter.
	MAV_ODID_VER_ACC_1_METER MAV_ODID_VER_ACC = 6
)

var labels_MAV_ODID_VER_ACC = map[MAV_ODID_VER_ACC]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_VER_ACC) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_VER_ACC[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_VER_ACC = map[string]MAV_ODID_VER_ACC{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_VER_ACC) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_VER_ACC[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_VER_ACC) String() string {
	if l, ok := labels_MAV_ODID_VER_ACC[e]; ok {
		return l
	}
	return "invalid value"
}
