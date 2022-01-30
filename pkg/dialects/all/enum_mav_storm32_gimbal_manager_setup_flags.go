//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

// Flags for gimbal manager set up. Used for setting and reporting, unless specified otherwise.
type MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS int

const (
	// Enable gimbal manager. This flag is only for setting, is not reported.
	MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS_ENABLE MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS = 16384
	// Disable gimbal manager. This flag is only for setting, is not reported.
	MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS_DISABLE MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS = 32768
)

var labels_MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS = map[MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS = map[string]MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS) String() string {
	if l, ok := labels_MAV_STORM32_GIMBAL_MANAGER_SETUP_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
