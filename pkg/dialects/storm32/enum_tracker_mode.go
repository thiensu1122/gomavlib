//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"errors"
)

// A mapping of antenna tracker flight modes for custom_mode field of heartbeat.
type TRACKER_MODE uint32

const (
	TRACKER_MODE_MANUAL       TRACKER_MODE = 0
	TRACKER_MODE_STOP         TRACKER_MODE = 1
	TRACKER_MODE_SCAN         TRACKER_MODE = 2
	TRACKER_MODE_SERVO_TEST   TRACKER_MODE = 3
	TRACKER_MODE_AUTO         TRACKER_MODE = 10
	TRACKER_MODE_INITIALIZING TRACKER_MODE = 16
)

var labels_TRACKER_MODE = map[TRACKER_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e TRACKER_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_TRACKER_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_TRACKER_MODE = map[string]TRACKER_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *TRACKER_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_TRACKER_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e TRACKER_MODE) String() string {
	if l, ok := labels_TRACKER_MODE[e]; ok {
		return l
	}
	return "invalid value"
}
