//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package pythonarraytest

import (
	"errors"
)

// Possible actions an aircraft can take to avoid a collision.
type MAV_COLLISION_ACTION uint32

const (
	// Ignore any potential collisions
	MAV_COLLISION_ACTION_NONE MAV_COLLISION_ACTION = 0
	// Report potential collision
	MAV_COLLISION_ACTION_REPORT MAV_COLLISION_ACTION = 1
	// Ascend or Descend to avoid threat
	MAV_COLLISION_ACTION_ASCEND_OR_DESCEND MAV_COLLISION_ACTION = 2
	// Move horizontally to avoid threat
	MAV_COLLISION_ACTION_MOVE_HORIZONTALLY MAV_COLLISION_ACTION = 3
	// Aircraft to move perpendicular to the collision's velocity vector
	MAV_COLLISION_ACTION_MOVE_PERPENDICULAR MAV_COLLISION_ACTION = 4
	// Aircraft to fly directly back to its launch point
	MAV_COLLISION_ACTION_RTL MAV_COLLISION_ACTION = 5
	// Aircraft to stop in place
	MAV_COLLISION_ACTION_HOVER MAV_COLLISION_ACTION = 6
)

var labels_MAV_COLLISION_ACTION = map[MAV_COLLISION_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COLLISION_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_COLLISION_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_COLLISION_ACTION = map[string]MAV_COLLISION_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COLLISION_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_COLLISION_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_COLLISION_ACTION) String() string {
	if l, ok := labels_MAV_COLLISION_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
