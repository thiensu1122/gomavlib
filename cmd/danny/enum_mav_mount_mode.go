//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Enumeration of possible mount operation modes. This message is used by obsolete/deprecated gimbal messages.
type MAV_MOUNT_MODE uint32

const (
    // Load and keep safe position (Roll,Pitch,Yaw) from permant memory and stop stabilization
    MAV_MOUNT_MODE_RETRACT MAV_MOUNT_MODE = 0
    // Load and keep neutral position (Roll,Pitch,Yaw) from permanent memory.
    MAV_MOUNT_MODE_NEUTRAL MAV_MOUNT_MODE = 1
    // Load neutral position and start MAVLink Roll,Pitch,Yaw control with stabilization
    MAV_MOUNT_MODE_MAVLINK_TARGETING MAV_MOUNT_MODE = 2
    // Load neutral position and start RC Roll,Pitch,Yaw control with stabilization
    MAV_MOUNT_MODE_RC_TARGETING MAV_MOUNT_MODE = 3
    // Load neutral position and start to point to Lat,Lon,Alt
    MAV_MOUNT_MODE_GPS_POINT MAV_MOUNT_MODE = 4
    // Gimbal tracks system with specified system ID
    MAV_MOUNT_MODE_SYSID_TARGET MAV_MOUNT_MODE = 5
    // Gimbal tracks home position
    MAV_MOUNT_MODE_HOME_LOCATION MAV_MOUNT_MODE = 6
)

var labels_MAV_MOUNT_MODE = map[MAV_MOUNT_MODE]string{
    MAV_MOUNT_MODE_RETRACT: "MAV_MOUNT_MODE_RETRACT",
    MAV_MOUNT_MODE_NEUTRAL: "MAV_MOUNT_MODE_NEUTRAL",
    MAV_MOUNT_MODE_MAVLINK_TARGETING: "MAV_MOUNT_MODE_MAVLINK_TARGETING",
    MAV_MOUNT_MODE_RC_TARGETING: "MAV_MOUNT_MODE_RC_TARGETING",
    MAV_MOUNT_MODE_GPS_POINT: "MAV_MOUNT_MODE_GPS_POINT",
    MAV_MOUNT_MODE_SYSID_TARGET: "MAV_MOUNT_MODE_SYSID_TARGET",
    MAV_MOUNT_MODE_HOME_LOCATION: "MAV_MOUNT_MODE_HOME_LOCATION",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_MOUNT_MODE) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_MOUNT_MODE {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_MOUNT_MODE) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_MOUNT_MODE
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_MOUNT_MODE {
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
func (e MAV_MOUNT_MODE) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
