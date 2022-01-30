//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package storm32

import (
	"errors"
)

// Flags for gimbal device operation. Used for setting and reporting, unless specified otherwise. Settings which are in violation of the capability flags are ignored by the gimbal device.
type MAV_STORM32_GIMBAL_DEVICE_FLAGS int

const (
	// Retracted safe position (no stabilization), takes presedence over NEUTRAL flag. If supported by the gimbal, the angles in the retracted position can be set in addition.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_RETRACT MAV_STORM32_GIMBAL_DEVICE_FLAGS = 1
	// Neutral position (horizontal, forward looking, with stabiliziation).
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_NEUTRAL MAV_STORM32_GIMBAL_DEVICE_FLAGS = 2
	// Lock roll angle to absolute angle relative to horizon (not relative to drone). This is generally the default.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_ROLL_LOCK MAV_STORM32_GIMBAL_DEVICE_FLAGS = 4
	// Lock pitch angle to absolute angle relative to horizon (not relative to drone). This is generally the default.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_PITCH_LOCK MAV_STORM32_GIMBAL_DEVICE_FLAGS = 8
	// Lock yaw angle to absolute angle relative to earth (not relative to drone). When the YAW_ABSOLUTE flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute), else it is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle).
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_YAW_LOCK MAV_STORM32_GIMBAL_DEVICE_FLAGS = 16
	// Gimbal device can accept absolute yaw angle input. This flag cannot be set, is only for reporting (attempts to set it are rejected by the gimbal device).
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_CAN_ACCEPT_YAW_ABSOLUTE MAV_STORM32_GIMBAL_DEVICE_FLAGS = 256
	// Yaw angle is absolute (is only accepted if CAN_ACCEPT_YAW_ABSOLUTE is set). If this flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute), else it is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle).
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_YAW_ABSOLUTE MAV_STORM32_GIMBAL_DEVICE_FLAGS = 512
	// RC control. The RC input signal fed to the gimbal device exclusively controls the gimbal's orientation. Overrides RC_MIXED flag if that is also set.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_RC_EXCLUSIVE MAV_STORM32_GIMBAL_DEVICE_FLAGS = 1024
	// RC control. The RC input signal fed to the gimbal device is mixed into the gimbal's orientation. Is overriden by RC_EXCLUSIVE flag if that is also set.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_RC_MIXED MAV_STORM32_GIMBAL_DEVICE_FLAGS = 2048
	// UINT16_MAX = ignore.
	MAV_STORM32_GIMBAL_DEVICE_FLAGS_NONE MAV_STORM32_GIMBAL_DEVICE_FLAGS = 65535
)

var labels_MAV_STORM32_GIMBAL_DEVICE_FLAGS = map[MAV_STORM32_GIMBAL_DEVICE_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_DEVICE_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_GIMBAL_DEVICE_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_GIMBAL_DEVICE_FLAGS = map[string]MAV_STORM32_GIMBAL_DEVICE_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_DEVICE_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_GIMBAL_DEVICE_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_DEVICE_FLAGS) String() string {
	if l, ok := labels_MAV_STORM32_GIMBAL_DEVICE_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
