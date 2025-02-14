//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Parachute actions. Trigger release and enable/disable auto-release.
type PARACHUTE_ACTION uint32

const (
    // Disable auto-release of parachute (i.e. release triggered by crash detectors).
    PARACHUTE_DISABLE PARACHUTE_ACTION = 0
    // Enable auto-release of parachute.
    PARACHUTE_ENABLE PARACHUTE_ACTION = 1
    // Release parachute and kill motors.
    PARACHUTE_RELEASE PARACHUTE_ACTION = 2
)

var labels_PARACHUTE_ACTION = map[PARACHUTE_ACTION]string{
    PARACHUTE_DISABLE: "PARACHUTE_DISABLE",
    PARACHUTE_ENABLE: "PARACHUTE_ENABLE",
    PARACHUTE_RELEASE: "PARACHUTE_RELEASE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PARACHUTE_ACTION) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_PARACHUTE_ACTION {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PARACHUTE_ACTION) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask PARACHUTE_ACTION
    for _, label := range labels {
        found := false
        for value, l := range labels_PARACHUTE_ACTION {
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
func (e PARACHUTE_ACTION) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
