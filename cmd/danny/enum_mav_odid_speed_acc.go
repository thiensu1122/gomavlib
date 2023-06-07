//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
type MAV_ODID_SPEED_ACC uint32

const (
    // The speed accuracy is unknown.
    MAV_ODID_SPEED_ACC_UNKNOWN MAV_ODID_SPEED_ACC = 0
    // The speed accuracy is smaller than 10 meters per second.
    MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 1
    // The speed accuracy is smaller than 3 meters per second.
    MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 2
    // The speed accuracy is smaller than 1 meters per second.
    MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 3
    // The speed accuracy is smaller than 0.3 meters per second.
    MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND MAV_ODID_SPEED_ACC = 4
)

var labels_MAV_ODID_SPEED_ACC = map[MAV_ODID_SPEED_ACC]string{
    MAV_ODID_SPEED_ACC_UNKNOWN: "MAV_ODID_SPEED_ACC_UNKNOWN",
    MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND: "MAV_ODID_SPEED_ACC_10_METERS_PER_SECOND",
    MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND: "MAV_ODID_SPEED_ACC_3_METERS_PER_SECOND",
    MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND: "MAV_ODID_SPEED_ACC_1_METERS_PER_SECOND",
    MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND: "MAV_ODID_SPEED_ACC_0_3_METERS_PER_SECOND",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ODID_SPEED_ACC) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_ODID_SPEED_ACC {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ODID_SPEED_ACC) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_ODID_SPEED_ACC
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_ODID_SPEED_ACC {
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
func (e MAV_ODID_SPEED_ACC) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
