//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"errors"
)

// Gripper actions.
type GRIPPER_ACTIONS uint32

const (
	// Gripper release cargo.
	GRIPPER_ACTION_RELEASE GRIPPER_ACTIONS = 0
	// Gripper grab onto cargo.
	GRIPPER_ACTION_GRAB GRIPPER_ACTIONS = 1
)

var labels_GRIPPER_ACTIONS = map[GRIPPER_ACTIONS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GRIPPER_ACTIONS) MarshalText() ([]byte, error) {
	if l, ok := labels_GRIPPER_ACTIONS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GRIPPER_ACTIONS = map[string]GRIPPER_ACTIONS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GRIPPER_ACTIONS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GRIPPER_ACTIONS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GRIPPER_ACTIONS) String() string {
	if l, ok := labels_GRIPPER_ACTIONS[e]; ok {
		return l
	}
	return "invalid value"
}
