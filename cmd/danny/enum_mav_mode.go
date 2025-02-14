//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// These defines are predefined OR-combined mode flags. There is no need to use values from this enum, but it
// simplifies the use of the mode flags. Note that manual input is enabled in all modes as a safety override.
type MAV_MODE uint32

const (
    // System is not ready to fly, booting, calibrating, etc. No flag is set.
    MAV_MODE_PREFLIGHT MAV_MODE = 0
    // System is allowed to be active, under assisted RC control.
    MAV_MODE_STABILIZE_DISARMED MAV_MODE = 80
    // System is allowed to be active, under assisted RC control.
    MAV_MODE_STABILIZE_ARMED MAV_MODE = 208
    // System is allowed to be active, under manual (RC) control, no stabilization
    MAV_MODE_MANUAL_DISARMED MAV_MODE = 64
    // System is allowed to be active, under manual (RC) control, no stabilization
    MAV_MODE_MANUAL_ARMED MAV_MODE = 192
    // System is allowed to be active, under autonomous control, manual setpoint
    MAV_MODE_GUIDED_DISARMED MAV_MODE = 88
    // System is allowed to be active, under autonomous control, manual setpoint
    MAV_MODE_GUIDED_ARMED MAV_MODE = 216
    // System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
    MAV_MODE_AUTO_DISARMED MAV_MODE = 92
    // System is allowed to be active, under autonomous control and navigation (the trajectory is decided onboard and not pre-programmed by waypoints)
    MAV_MODE_AUTO_ARMED MAV_MODE = 220
    // UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
    MAV_MODE_TEST_DISARMED MAV_MODE = 66
    // UNDEFINED mode. This solely depends on the autopilot - use with caution, intended for developers only.
    MAV_MODE_TEST_ARMED MAV_MODE = 194
)

var labels_MAV_MODE = map[MAV_MODE]string{
    MAV_MODE_PREFLIGHT: "MAV_MODE_PREFLIGHT",
    MAV_MODE_STABILIZE_DISARMED: "MAV_MODE_STABILIZE_DISARMED",
    MAV_MODE_STABILIZE_ARMED: "MAV_MODE_STABILIZE_ARMED",
    MAV_MODE_MANUAL_DISARMED: "MAV_MODE_MANUAL_DISARMED",
    MAV_MODE_MANUAL_ARMED: "MAV_MODE_MANUAL_ARMED",
    MAV_MODE_GUIDED_DISARMED: "MAV_MODE_GUIDED_DISARMED",
    MAV_MODE_GUIDED_ARMED: "MAV_MODE_GUIDED_ARMED",
    MAV_MODE_AUTO_DISARMED: "MAV_MODE_AUTO_DISARMED",
    MAV_MODE_AUTO_ARMED: "MAV_MODE_AUTO_ARMED",
    MAV_MODE_TEST_DISARMED: "MAV_MODE_TEST_DISARMED",
    MAV_MODE_TEST_ARMED: "MAV_MODE_TEST_ARMED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MODE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_MODE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MODE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_MODE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_MODE {
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
func (e MAV_MODE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
