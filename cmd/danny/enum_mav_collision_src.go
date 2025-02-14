//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Source of information about this collision.
type MAV_COLLISION_SRC uint32

const (
    // ID field references ADSB_VEHICLE packets
    MAV_COLLISION_SRC_ADSB MAV_COLLISION_SRC = 0
    // ID field references MAVLink SRC ID
    MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT MAV_COLLISION_SRC = 1
)

var labels_MAV_COLLISION_SRC = map[MAV_COLLISION_SRC]string{
    MAV_COLLISION_SRC_ADSB: "MAV_COLLISION_SRC_ADSB",
    MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT: "MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COLLISION_SRC) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_COLLISION_SRC {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COLLISION_SRC) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_COLLISION_SRC
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_COLLISION_SRC {
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
func (e MAV_COLLISION_SRC) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
