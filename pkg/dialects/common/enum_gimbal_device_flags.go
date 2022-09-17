//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package common

import (
	"errors"
)

// Flags for gimbal device (lower level) operation.
type GIMBAL_DEVICE_FLAGS uint32

const (
	// Set to retracted safe position (no stabilization), takes presedence over all other flags.
	GIMBAL_DEVICE_FLAGS_RETRACT GIMBAL_DEVICE_FLAGS = 1
	// Set to neutral/default position, taking precedence over all other flags except RETRACT. Neutral is commonly forward-facing and horizontal (pitch=yaw=0) but may be any orientation.
	GIMBAL_DEVICE_FLAGS_NEUTRAL GIMBAL_DEVICE_FLAGS = 2
	// Lock roll angle to absolute angle relative to horizon (not relative to drone). This is generally the default with a stabilizing gimbal.
	GIMBAL_DEVICE_FLAGS_ROLL_LOCK GIMBAL_DEVICE_FLAGS = 4
	// Lock pitch angle to absolute angle relative to horizon (not relative to drone). This is generally the default.
	GIMBAL_DEVICE_FLAGS_PITCH_LOCK GIMBAL_DEVICE_FLAGS = 8
	// Lock yaw angle to absolute angle relative to North (not relative to drone). If this flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute). If this flag is not set, the quaternion frame is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle).
	GIMBAL_DEVICE_FLAGS_YAW_LOCK GIMBAL_DEVICE_FLAGS = 16
)

var labels_GIMBAL_DEVICE_FLAGS = map[GIMBAL_DEVICE_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_DEVICE_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_DEVICE_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_DEVICE_FLAGS = map[string]GIMBAL_DEVICE_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_DEVICE_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_DEVICE_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_DEVICE_FLAGS) String() string {
	if l, ok := labels_GIMBAL_DEVICE_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
