//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// The ROI (region of interest) for the vehicle. This can be
// be used by the vehicle for camera/vehicle attitude alignment (see
// MAV_CMD_NAV_ROI).
type MAV_ROI uint32

const (
    // No region of interest.
    MAV_ROI_NONE MAV_ROI = 0
    // Point toward next waypoint, with optional pitch/roll/yaw offset.
    MAV_ROI_WPNEXT MAV_ROI = 1
    // Point toward given waypoint.
    MAV_ROI_WPINDEX MAV_ROI = 2
    // Point toward fixed location.
    MAV_ROI_LOCATION MAV_ROI = 3
    // Point toward of given id.
    MAV_ROI_TARGET MAV_ROI = 4
)

var labels_MAV_ROI = map[MAV_ROI]string{
    MAV_ROI_NONE: "MAV_ROI_NONE",
    MAV_ROI_WPNEXT: "MAV_ROI_WPNEXT",
    MAV_ROI_WPINDEX: "MAV_ROI_WPINDEX",
    MAV_ROI_LOCATION: "MAV_ROI_LOCATION",
    MAV_ROI_TARGET: "MAV_ROI_TARGET",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ROI) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_ROI {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ROI) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_ROI
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_ROI {
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
func (e MAV_ROI) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
