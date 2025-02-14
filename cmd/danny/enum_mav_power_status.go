//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Power supply status flags (bitmask)
type MAV_POWER_STATUS uint32

const (
    // main brick power supply valid
    MAV_POWER_STATUS_BRICK_VALID MAV_POWER_STATUS = 1
    // main servo power supply valid for FMU
    MAV_POWER_STATUS_SERVO_VALID MAV_POWER_STATUS = 2
    // USB power is connected
    MAV_POWER_STATUS_USB_CONNECTED MAV_POWER_STATUS = 4
    // peripheral supply is in over-current state
    MAV_POWER_STATUS_PERIPH_OVERCURRENT MAV_POWER_STATUS = 8
    // hi-power peripheral supply is in over-current state
    MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT MAV_POWER_STATUS = 16
    // Power status has changed since boot
    MAV_POWER_STATUS_CHANGED MAV_POWER_STATUS = 32
)

var labels_MAV_POWER_STATUS = map[MAV_POWER_STATUS]string{
    MAV_POWER_STATUS_BRICK_VALID: "MAV_POWER_STATUS_BRICK_VALID",
    MAV_POWER_STATUS_SERVO_VALID: "MAV_POWER_STATUS_SERVO_VALID",
    MAV_POWER_STATUS_USB_CONNECTED: "MAV_POWER_STATUS_USB_CONNECTED",
    MAV_POWER_STATUS_PERIPH_OVERCURRENT: "MAV_POWER_STATUS_PERIPH_OVERCURRENT",
    MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT: "MAV_POWER_STATUS_PERIPH_HIPOWER_OVERCURRENT",
    MAV_POWER_STATUS_CHANGED: "MAV_POWER_STATUS_CHANGED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_POWER_STATUS) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_POWER_STATUS {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_POWER_STATUS) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_POWER_STATUS
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_POWER_STATUS {
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
func (e MAV_POWER_STATUS) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
