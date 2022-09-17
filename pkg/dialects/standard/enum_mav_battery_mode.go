//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package standard

import (
	"errors"
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

var labels_MAV_BATTERY_MODE = map[MAV_BATTERY_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_BATTERY_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_BATTERY_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_BATTERY_MODE = map[string]MAV_BATTERY_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_BATTERY_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_BATTERY_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_BATTERY_MODE) String() string {
	if l, ok := labels_MAV_BATTERY_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
