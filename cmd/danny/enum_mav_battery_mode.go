//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Battery mode. Note, the normal operation mode (i.e. when flying) should be reported as MAV_BATTERY_MODE_UNKNOWN to allow message trimming in normal flight.
type MAV_BATTERY_MODE uint32

const (
    // Battery mode not supported/unknown battery mode/normal operation.
    MAV_BATTERY_MODE_UNKNOWN MAV_BATTERY_MODE = 0
    // Battery is auto discharging (towards storage level).
    MAV_BATTERY_MODE_AUTO_DISCHARGING MAV_BATTERY_MODE = 1
    // Battery in hot-swap mode (current limited to prevent spikes that might damage sensitive electrical circuits).
    MAV_BATTERY_MODE_HOT_SWAP MAV_BATTERY_MODE = 2
)

var labels_MAV_BATTERY_MODE = map[MAV_BATTERY_MODE]string{
    MAV_BATTERY_MODE_UNKNOWN: "MAV_BATTERY_MODE_UNKNOWN",
    MAV_BATTERY_MODE_AUTO_DISCHARGING: "MAV_BATTERY_MODE_AUTO_DISCHARGING",
    MAV_BATTERY_MODE_HOT_SWAP: "MAV_BATTERY_MODE_HOT_SWAP",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_MODE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_BATTERY_MODE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_MODE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_BATTERY_MODE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_BATTERY_MODE {
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
func (e MAV_BATTERY_MODE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
