//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Micro air vehicle / autopilot classes. This identifies the individual model.
type MAV_COMMAND_TYPE uint32

const (
    // Generic autopilot, full support for everything
    MAV_AUTOPILOT_GENERIC MAV_COMMAND_TYPE = 0
)

var labels_MAV_COMMAND_TYPE = map[MAV_COMMAND_TYPE]string{
    MAV_AUTOPILOT_GENERIC: "MAV_AUTOPILOT_GENERIC",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COMMAND_TYPE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_COMMAND_TYPE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COMMAND_TYPE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_COMMAND_TYPE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_COMMAND_TYPE {
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
func (e MAV_COMMAND_TYPE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
