//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type PID_TUNING_AXIS int

const (
	PID_TUNING_ROLL    PID_TUNING_AXIS = 1
	PID_TUNING_PITCH   PID_TUNING_AXIS = 2
	PID_TUNING_YAW     PID_TUNING_AXIS = 3
	PID_TUNING_ACCZ    PID_TUNING_AXIS = 4
	PID_TUNING_STEER   PID_TUNING_AXIS = 5
	PID_TUNING_LANDING PID_TUNING_AXIS = 6
)

var labels_PID_TUNING_AXIS = map[PID_TUNING_AXIS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PID_TUNING_AXIS) MarshalText() ([]byte, error) {
	if l, ok := labels_PID_TUNING_AXIS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PID_TUNING_AXIS = map[string]PID_TUNING_AXIS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PID_TUNING_AXIS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PID_TUNING_AXIS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PID_TUNING_AXIS) String() string {
	if l, ok := labels_PID_TUNING_AXIS[e]; ok {
		return l
	}
	return "invalid value"
}
