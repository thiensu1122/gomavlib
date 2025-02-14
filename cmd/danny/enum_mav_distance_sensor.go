//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package danny

import (
    "strings"
    "fmt"
)
// Enumeration of distance sensor types
type MAV_DISTANCE_SENSOR uint32

const (
    // Laser rangefinder, e.g. LightWare SF02/F or PulsedLight units
    MAV_DISTANCE_SENSOR_LASER MAV_DISTANCE_SENSOR = 0
    // Ultrasound rangefinder, e.g. MaxBotix units
    MAV_DISTANCE_SENSOR_ULTRASOUND MAV_DISTANCE_SENSOR = 1
    // Infrared rangefinder, e.g. Sharp units
    MAV_DISTANCE_SENSOR_INFRARED MAV_DISTANCE_SENSOR = 2
    // Radar type, e.g. uLanding units
    MAV_DISTANCE_SENSOR_RADAR MAV_DISTANCE_SENSOR = 3
    // Broken or unknown type, e.g. analog units
    MAV_DISTANCE_SENSOR_UNKNOWN MAV_DISTANCE_SENSOR = 4
)

var labels_MAV_DISTANCE_SENSOR = map[MAV_DISTANCE_SENSOR]string{
    MAV_DISTANCE_SENSOR_LASER: "MAV_DISTANCE_SENSOR_LASER",
    MAV_DISTANCE_SENSOR_ULTRASOUND: "MAV_DISTANCE_SENSOR_ULTRASOUND",
    MAV_DISTANCE_SENSOR_INFRARED: "MAV_DISTANCE_SENSOR_INFRARED",
    MAV_DISTANCE_SENSOR_RADAR: "MAV_DISTANCE_SENSOR_RADAR",
    MAV_DISTANCE_SENSOR_UNKNOWN: "MAV_DISTANCE_SENSOR_UNKNOWN",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_DISTANCE_SENSOR) MarshalText() ([]byte, error) {
    var names []string
    for mask, label := range labels_MAV_DISTANCE_SENSOR {
        if e&mask == mask {
            names = append(names, label)
        }
    }
    return []byte(strings.Join(names, " | ")), nil
}


// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_DISTANCE_SENSOR) UnmarshalText(text []byte) error {
    labels := strings.Split(string(text), " | ")
    var mask MAV_DISTANCE_SENSOR
    for _, label := range labels {
        found := false
        for value, l := range labels_MAV_DISTANCE_SENSOR {
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
func (e MAV_DISTANCE_SENSOR) String() string {
    val, _ := e.MarshalText()
    return string(val)
}
