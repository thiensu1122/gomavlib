//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"errors"
)

type MAV_ODID_TIME_ACC uint32

const (
	// The timestamp accuracy is unknown.
	MAV_ODID_TIME_ACC_UNKNOWN MAV_ODID_TIME_ACC = 0
	// The timestamp accuracy is smaller than or equal to 0.1 second.
	MAV_ODID_TIME_ACC_0_1_SECOND MAV_ODID_TIME_ACC = 1
	// The timestamp accuracy is smaller than or equal to 0.2 second.
	MAV_ODID_TIME_ACC_0_2_SECOND MAV_ODID_TIME_ACC = 2
	// The timestamp accuracy is smaller than or equal to 0.3 second.
	MAV_ODID_TIME_ACC_0_3_SECOND MAV_ODID_TIME_ACC = 3
	// The timestamp accuracy is smaller than or equal to 0.4 second.
	MAV_ODID_TIME_ACC_0_4_SECOND MAV_ODID_TIME_ACC = 4
	// The timestamp accuracy is smaller than or equal to 0.5 second.
	MAV_ODID_TIME_ACC_0_5_SECOND MAV_ODID_TIME_ACC = 5
	// The timestamp accuracy is smaller than or equal to 0.6 second.
	MAV_ODID_TIME_ACC_0_6_SECOND MAV_ODID_TIME_ACC = 6
	// The timestamp accuracy is smaller than or equal to 0.7 second.
	MAV_ODID_TIME_ACC_0_7_SECOND MAV_ODID_TIME_ACC = 7
	// The timestamp accuracy is smaller than or equal to 0.8 second.
	MAV_ODID_TIME_ACC_0_8_SECOND MAV_ODID_TIME_ACC = 8
	// The timestamp accuracy is smaller than or equal to 0.9 second.
	MAV_ODID_TIME_ACC_0_9_SECOND MAV_ODID_TIME_ACC = 9
	// The timestamp accuracy is smaller than or equal to 1.0 second.
	MAV_ODID_TIME_ACC_1_0_SECOND MAV_ODID_TIME_ACC = 10
	// The timestamp accuracy is smaller than or equal to 1.1 second.
	MAV_ODID_TIME_ACC_1_1_SECOND MAV_ODID_TIME_ACC = 11
	// The timestamp accuracy is smaller than or equal to 1.2 second.
	MAV_ODID_TIME_ACC_1_2_SECOND MAV_ODID_TIME_ACC = 12
	// The timestamp accuracy is smaller than or equal to 1.3 second.
	MAV_ODID_TIME_ACC_1_3_SECOND MAV_ODID_TIME_ACC = 13
	// The timestamp accuracy is smaller than or equal to 1.4 second.
	MAV_ODID_TIME_ACC_1_4_SECOND MAV_ODID_TIME_ACC = 14
	// The timestamp accuracy is smaller than or equal to 1.5 second.
	MAV_ODID_TIME_ACC_1_5_SECOND MAV_ODID_TIME_ACC = 15
)

var labels_MAV_ODID_TIME_ACC = map[MAV_ODID_TIME_ACC]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_TIME_ACC) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_TIME_ACC[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_TIME_ACC = map[string]MAV_ODID_TIME_ACC{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_TIME_ACC) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_TIME_ACC[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_TIME_ACC) String() string {
	if l, ok := labels_MAV_ODID_TIME_ACC[e]; ok {
		return l
	}
	return "invalid value"
}
