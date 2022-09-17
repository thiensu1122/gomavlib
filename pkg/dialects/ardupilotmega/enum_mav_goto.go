//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"errors"
)

// Actions that may be specified in MAV_CMD_OVERRIDE_GOTO to override mission execution.
type MAV_GOTO uint32

const (
	// Hold at the current position.
	MAV_GOTO_DO_HOLD MAV_GOTO = 0
	// Continue with the next item in mission execution.
	MAV_GOTO_DO_CONTINUE MAV_GOTO = 1
	// Hold at the current position of the system
	MAV_GOTO_HOLD_AT_CURRENT_POSITION MAV_GOTO = 2
	// Hold at the position specified in the parameters of the DO_HOLD action
	MAV_GOTO_HOLD_AT_SPECIFIED_POSITION MAV_GOTO = 3
)

var labels_MAV_GOTO = map[MAV_GOTO]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_GOTO) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_GOTO[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_GOTO = map[string]MAV_GOTO{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_GOTO) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_GOTO[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_GOTO) String() string {
	if l, ok := labels_MAV_GOTO[e]; ok {
		return l
	}
	return "invalid value"
}
