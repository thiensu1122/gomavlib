//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Bitmap to indicate which dimensions should be ignored by the vehicle: a value of 0b00000000 indicates that none of the setpoint dimensions should be ignored.
type ATTITUDE_TARGET_TYPEMASK uint32

const (
	// Ignore body roll rate
	ATTITUDE_TARGET_TYPEMASK_BODY_ROLL_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 1
	// Ignore body pitch rate
	ATTITUDE_TARGET_TYPEMASK_BODY_PITCH_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 2
	// Ignore body yaw rate
	ATTITUDE_TARGET_TYPEMASK_BODY_YAW_RATE_IGNORE ATTITUDE_TARGET_TYPEMASK = 4
	// Use 3D body thrust setpoint instead of throttle
	ATTITUDE_TARGET_TYPEMASK_THRUST_BODY_SET ATTITUDE_TARGET_TYPEMASK = 32
	// Ignore throttle
	ATTITUDE_TARGET_TYPEMASK_THROTTLE_IGNORE ATTITUDE_TARGET_TYPEMASK = 64
	// Ignore attitude
	ATTITUDE_TARGET_TYPEMASK_ATTITUDE_IGNORE ATTITUDE_TARGET_TYPEMASK = 128
)

var labels_ATTITUDE_TARGET_TYPEMASK = map[ATTITUDE_TARGET_TYPEMASK]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ATTITUDE_TARGET_TYPEMASK) MarshalText() ([]byte, error) {
	if l, ok := labels_ATTITUDE_TARGET_TYPEMASK[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ATTITUDE_TARGET_TYPEMASK = map[string]ATTITUDE_TARGET_TYPEMASK{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ATTITUDE_TARGET_TYPEMASK) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ATTITUDE_TARGET_TYPEMASK[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ATTITUDE_TARGET_TYPEMASK) String() string {
	if l, ok := labels_ATTITUDE_TARGET_TYPEMASK[e]; ok {
		return l
	}
	return "invalid value"
}
