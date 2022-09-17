//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"errors"
)

// Flags for the global position report.
type UTM_DATA_AVAIL_FLAGS uint32

const (
	// The field time contains valid data.
	UTM_DATA_AVAIL_FLAGS_TIME_VALID UTM_DATA_AVAIL_FLAGS = 1
	// The field uas_id contains valid data.
	UTM_DATA_AVAIL_FLAGS_UAS_ID_AVAILABLE UTM_DATA_AVAIL_FLAGS = 2
	// The fields lat, lon and h_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_POSITION_AVAILABLE UTM_DATA_AVAIL_FLAGS = 4
	// The fields alt and v_acc contain valid data.
	UTM_DATA_AVAIL_FLAGS_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 8
	// The field relative_alt contains valid data.
	UTM_DATA_AVAIL_FLAGS_RELATIVE_ALTITUDE_AVAILABLE UTM_DATA_AVAIL_FLAGS = 16
	// The fields vx and vy contain valid data.
	UTM_DATA_AVAIL_FLAGS_HORIZONTAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 32
	// The field vz contains valid data.
	UTM_DATA_AVAIL_FLAGS_VERTICAL_VELO_AVAILABLE UTM_DATA_AVAIL_FLAGS = 64
	// The fields next_lat, next_lon and next_alt contain valid data.
	UTM_DATA_AVAIL_FLAGS_NEXT_WAYPOINT_AVAILABLE UTM_DATA_AVAIL_FLAGS = 128
)

var labels_UTM_DATA_AVAIL_FLAGS = map[UTM_DATA_AVAIL_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UTM_DATA_AVAIL_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_UTM_DATA_AVAIL_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UTM_DATA_AVAIL_FLAGS = map[string]UTM_DATA_AVAIL_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UTM_DATA_AVAIL_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UTM_DATA_AVAIL_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UTM_DATA_AVAIL_FLAGS) String() string {
	if l, ok := labels_UTM_DATA_AVAIL_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
