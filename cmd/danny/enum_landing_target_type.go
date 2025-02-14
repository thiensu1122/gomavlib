//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Type of landing target
type LANDING_TARGET_TYPE uint32

const (
    // Landing target signaled by light beacon (ex: IR-LOCK)
    LANDING_TARGET_TYPE_LIGHT_BEACON LANDING_TARGET_TYPE = 0
    // Landing target signaled by radio beacon (ex: ILS, NDB)
    LANDING_TARGET_TYPE_RADIO_BEACON LANDING_TARGET_TYPE = 1
    // Landing target represented by a fiducial marker (ex: ARTag)
    LANDING_TARGET_TYPE_VISION_FIDUCIAL LANDING_TARGET_TYPE = 2
    // Landing target represented by a pre-defined visual shape/feature (ex: X-marker, H-marker, square)
    LANDING_TARGET_TYPE_VISION_OTHER LANDING_TARGET_TYPE = 3
)

var labels_LANDING_TARGET_TYPE = map[LANDING_TARGET_TYPE]string{
    LANDING_TARGET_TYPE_LIGHT_BEACON: "LANDING_TARGET_TYPE_LIGHT_BEACON",
    LANDING_TARGET_TYPE_RADIO_BEACON: "LANDING_TARGET_TYPE_RADIO_BEACON",
    LANDING_TARGET_TYPE_VISION_FIDUCIAL: "LANDING_TARGET_TYPE_VISION_FIDUCIAL",
    LANDING_TARGET_TYPE_VISION_OTHER: "LANDING_TARGET_TYPE_VISION_OTHER",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e LANDING_TARGET_TYPE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_LANDING_TARGET_TYPE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *LANDING_TARGET_TYPE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask LANDING_TARGET_TYPE
    for _, label := range labels {
        found := false
        for value, l := range labels_LANDING_TARGET_TYPE {
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
func (e LANDING_TARGET_TYPE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
