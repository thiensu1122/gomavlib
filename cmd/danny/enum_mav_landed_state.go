//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Enumeration of landed detector states
type MAV_LANDED_STATE uint32

const (
    // MAV landed state is unknown
    MAV_LANDED_STATE_UNDEFINED MAV_LANDED_STATE = 0
    // MAV is landed (on ground)
    MAV_LANDED_STATE_ON_GROUND MAV_LANDED_STATE = 1
    // MAV is in air
    MAV_LANDED_STATE_IN_AIR MAV_LANDED_STATE = 2
    // MAV currently taking off
    MAV_LANDED_STATE_TAKEOFF MAV_LANDED_STATE = 3
    // MAV currently landing
    MAV_LANDED_STATE_LANDING MAV_LANDED_STATE = 4
)

var labels_MAV_LANDED_STATE = map[MAV_LANDED_STATE]string{
    MAV_LANDED_STATE_UNDEFINED: "MAV_LANDED_STATE_UNDEFINED",
    MAV_LANDED_STATE_ON_GROUND: "MAV_LANDED_STATE_ON_GROUND",
    MAV_LANDED_STATE_IN_AIR: "MAV_LANDED_STATE_IN_AIR",
    MAV_LANDED_STATE_TAKEOFF: "MAV_LANDED_STATE_TAKEOFF",
    MAV_LANDED_STATE_LANDING: "MAV_LANDED_STATE_LANDING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_LANDED_STATE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_LANDED_STATE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_LANDED_STATE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_LANDED_STATE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_LANDED_STATE {
            if l == label {
                mask |= value
                found = true
                break
            }
        }
        if !found {
            return fmt.Errorf("invalid label '%s'", label)
        }
    }
    *e = mask
    return nil
}

// String implements the fmt.Stringer interface.
func (e MAV_LANDED_STATE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
