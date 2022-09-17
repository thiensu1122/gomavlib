//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"errors"
)

type MAV_ODID_HOR_ACC uint32

const (
	// The horizontal accuracy is unknown.
	MAV_ODID_HOR_ACC_UNKNOWN MAV_ODID_HOR_ACC = 0
	// The horizontal accuracy is smaller than 10 Nautical Miles. 18.52 km.
	MAV_ODID_HOR_ACC_10NM MAV_ODID_HOR_ACC = 1
	// The horizontal accuracy is smaller than 4 Nautical Miles. 7.408 km.
	MAV_ODID_HOR_ACC_4NM MAV_ODID_HOR_ACC = 2
	// The horizontal accuracy is smaller than 2 Nautical Miles. 3.704 km.
	MAV_ODID_HOR_ACC_2NM MAV_ODID_HOR_ACC = 3
	// The horizontal accuracy is smaller than 1 Nautical Miles. 1.852 km.
	MAV_ODID_HOR_ACC_1NM MAV_ODID_HOR_ACC = 4
	// The horizontal accuracy is smaller than 0.5 Nautical Miles. 926 m.
	MAV_ODID_HOR_ACC_0_5NM MAV_ODID_HOR_ACC = 5
	// The horizontal accuracy is smaller than 0.3 Nautical Miles. 555.6 m.
	MAV_ODID_HOR_ACC_0_3NM MAV_ODID_HOR_ACC = 6
	// The horizontal accuracy is smaller than 0.1 Nautical Miles. 185.2 m.
	MAV_ODID_HOR_ACC_0_1NM MAV_ODID_HOR_ACC = 7
	// The horizontal accuracy is smaller than 0.05 Nautical Miles. 92.6 m.
	MAV_ODID_HOR_ACC_0_05NM MAV_ODID_HOR_ACC = 8
	// The horizontal accuracy is smaller than 30 meter.
	MAV_ODID_HOR_ACC_30_METER MAV_ODID_HOR_ACC = 9
	// The horizontal accuracy is smaller than 10 meter.
	MAV_ODID_HOR_ACC_10_METER MAV_ODID_HOR_ACC = 10
	// The horizontal accuracy is smaller than 3 meter.
	MAV_ODID_HOR_ACC_3_METER MAV_ODID_HOR_ACC = 11
	// The horizontal accuracy is smaller than 1 meter.
	MAV_ODID_HOR_ACC_1_METER MAV_ODID_HOR_ACC = 12
)

var labels_MAV_ODID_HOR_ACC = map[MAV_ODID_HOR_ACC]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_HOR_ACC) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ODID_HOR_ACC[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ODID_HOR_ACC = map[string]MAV_ODID_HOR_ACC{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_HOR_ACC) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ODID_HOR_ACC[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ODID_HOR_ACC) String() string {
	if l, ok := labels_MAV_ODID_HOR_ACC[e]; ok {
		return l
	}
	return "invalid value"
}
