//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"errors"
)

type MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL uint32

const (
	// Switch Low.
	MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL_LOW MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL = 0
	// Switch Middle.
	MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL_MIDDLE MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL = 1
	// Switch High.
	MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL_HIGH MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL = 2
)

var labels_MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL = map[MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL = map[string]MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL) String() string {
	if l, ok := labels_MAV_CMD_DO_AUX_FUNCTION_SWITCH_LEVEL[e]; ok {
		return l
	}
	return "invalid value"
}
